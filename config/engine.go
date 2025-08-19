package config

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	gincors "github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	global "github.com/rs/zerolog/log"
	lm "github.com/ulule/limiter/v3"
	lmgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	lsredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

var (
	reqIDPrefix string
	reqIDCount  uint64

	corsMaxAge         = 12 * time.Hour
	corsAllowedOrigins = []string{
		"http://localhost:3000",
		"http://localhost:8000",
	}
	corsAllowedHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"Cache-Control",
		"Origin",
		"Cookie",
		"X-Requested-With",
		"X-Request-ID",
	}
	corsAllowedMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}
)

func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		if _, err = rand.Read(buf[:]); err != nil {
			continue
		}
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	reqIDPrefix = fmt.Sprintf("%s-%s", hostname, b64[0:10])
}

func logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		path := ctx.Request.URL.Path
		if raw := ctx.Request.URL.RawQuery; raw != "" {
			path += "?" + raw
		}
		statusCode := ctx.Writer.Status()
		logLevel := determineLogLevel(statusCode)
		userID := 0
		if val, exists := ctx.Get("user_id"); exists {
			if f64, ok := val.(float64); ok {
				userID = int(f64)
			}
		}
		roleName := "-"
		if val, exists := ctx.Get("role_name"); exists {
			if rn, ok := val.(string); ok {
				roleName = rn
			}
		}
		method := ctx.Request.Method
		requestID := requestid.Get(ctx)
		userAgent := ctx.Request.UserAgent()
		clientIP := ctx.ClientIP()
		protocol := ctx.Request.Proto
		bodySize := ctx.Writer.Size()
		latency := time.Since(start).String()
		errorMsg := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		logMsg := fmt.Sprintf("HTTPAccessLog | %7s | %3d | %s | rid: %s, uid: %d ",
			method, statusCode, path, requestID, userID)

		// print log
		go global.Logger.WithLevel(logLevel).
			Str("client_ip", clientIP).
			Str("request_id", requestID).
			Str("protocol", protocol).
			Str("agent", userAgent).
			Str("method", method).
			Int("status_code", statusCode).
			Int("body_size", bodySize).
			Str("path", path).
			Str("latency", latency).
			Str("error", errorMsg).
			Msg(logMsg)

		// store log
		go func() {
			if userID == 0 || roleName == "-" {
				ctx.Next()
				return
			}

			stmt := "INSERT INTO activity_logs (user_id, role, description, created_at) values ($1, $2, $3, EXTRACT(EPOCH FROM NOW())::BIGINT)"
			if _, err := PgxPool.Exec(ctx, stmt, userID, roleName, logMsg); err != nil {
				log.Println("failed to store activity log", err.Error())
				ctx.Next()
				return
			}

			ctx.Next()
		}()
	}
}

func determineLogLevel(status int) zerolog.Level {
	switch {
	case status >= http.StatusInternalServerError:
		return zerolog.ErrorLevel
	case status >= http.StatusBadRequest:
		return zerolog.WarnLevel
	default:
		return zerolog.InfoLevel
	}
}

func cors() gin.HandlerFunc {
	return gincors.New(gincors.Config{
		MaxAge:           corsMaxAge,
		AllowOrigins:     corsAllowedOrigins,
		AllowMethods:     corsAllowedMethods,
		AllowHeaders:     corsAllowedHeaders,
		AllowCredentials: true,
	})
}

func limiter(
	serverLimiter,
	serverName string,
) gin.HandlerFunc {
	rate, err := lm.NewRateFromFormatted(serverLimiter)
	if err != nil {
		log.Fatalf("RATELIMITER_ERROR: %s\n", err.Error())
	}
	store, err := lsredis.NewStoreWithOptions(RdpPool, lm.StoreOptions{Prefix: serverName})
	if err != nil {
		log.Fatalf("RATELIMITER_ERROR: %s\n", err.Error())
	}
	return lmgin.NewMiddleware(lm.New(store, rate))
}

func ServerEngine() Option {
	return func(cfg *Config) {
		gin.SetMode(gin.ReleaseMode)
		if cfg.AppDebug {
			gin.SetMode(gin.DebugMode)
		}
		// setup basic middleware
		engine := gin.New()
		engine.ForwardedByClientIP = true
		engine.Use(requestid.New(requestid.WithGenerator(func() string {
			myID := atomic.AddUint64(&reqIDCount, 1)
			return fmt.Sprintf("%s-%06d", reqIDPrefix, myID)
		})), logger(), gin.Recovery(), cors())
		if !cfg.AppDebug {
			// setup sentry middleware
			GinEngine.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
			GinEngine.Use(func(ctx *gin.Context) {
				if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
					hub.Scope().SetTag("CurrentServer", fmt.Sprintf(
						"[%s]%s", cfg.AppName, cfg.AppVersion))
				}
				ctx.Next()
			})
			// setup rate limit middleware
			serverInitial := fmt.Sprintf("%s %s", cfg.AppName, cfg.AppVersion)
			engine.Use(limiter(cfg.APILimiter, serverInitial))
		}
		GinEngine = engine
		log.Println("Gin engine ready!")
	}
}
