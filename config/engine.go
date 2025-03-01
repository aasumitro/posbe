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
		var zll zerolog.Level
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		ctx.Next()
		if raw != "" {
			path = path + "?" + raw
		}
		statusCode := ctx.Writer.Status()
		switch {
		case statusCode >= http.StatusInternalServerError:
			zll = zerolog.ErrorLevel
		case statusCode >= http.StatusBadRequest:
			zll = zerolog.WarnLevel
		default:
			zll = zerolog.InfoLevel
		}
		mtd := ctx.Request.Method
		rid := requestid.Get(ctx)

		// print log
		go global.Logger.WithLevel(zll).
			Str("client_ip", ctx.ClientIP()).
			Str("request_id", rid).
			Str("protocol", ctx.Request.Proto).
			Str("agent", ctx.Request.UserAgent()).
			Str("method", mtd).
			Int("status_code", statusCode).
			Int("body_size", ctx.Writer.Size()).
			Str("path", path).
			Str("latency", time.Since(start).String()).
			Str("error", ctx.Errors.ByType(gin.ErrorTypePrivate).String()).
			Msg(fmt.Sprintf("HTTPAccessLog | %7s | %3d | %s | rid: %s, uid: %s ",
				mtd, statusCode, path, rid, func() string {
					uid, ok := ctx.Get("user_id")
					if !ok {
						return "-"
					}
					return fmt.Sprintf("%0d", int(uid.(float64)))
				}()),
			)

		// store log
		go func() {
			uid, ok := ctx.Get("user_id")
			if !ok {
				log.Println("failed to get user_id")
				ctx.Next()
				return
			}
			rn, ok := ctx.Get("role_name")
			if !ok {
				log.Println("failed to get role name")
				ctx.Next()
				return
			}
			if _, err := PostgresPool.ExecContext(ctx,
				"INSERT INTO activity_logs (user_id, role, description, created_at) values ($1, $2, $3, $4)",
				int(uid.(float64)), rn.(string), fmt.Sprintf(
					"HTTPAccessLog | %7s | %3d | %s | rid: %s, uid: %d ",
					mtd, statusCode, path, rid, int(uid.(float64)),
				), time.Now().Unix(),
			); err != nil {
				log.Println("failed to store activity log", err.Error())
				ctx.Next()
				return
			}
			ctx.Next()
		}()
	}
}

func cors() gin.HandlerFunc {
	return gincors.New(gincors.Config{
		AllowOrigins:     corsAllowedOrigins,
		AllowMethods:     corsAllowedMethods,
		AllowHeaders:     corsAllowedHeaders,
		AllowCredentials: true,
		MaxAge:           corsMaxAge,
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
	store, err := lsredis.NewStoreWithOptions(RedisPool,
		lm.StoreOptions{Prefix: serverName})
	if err != nil {
		log.Fatalf("RATELIMITER_ERROR: %s\n", err.Error())
	}
	return lmgin.NewMiddleware(lm.New(store, rate))
}

func ServerEngine() Option {
	return func(cfg *Config) {
		engineOnce.Do(func() {
			log.Printf("Init router engine [GIN Framework %s]\n", gin.Version)
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
		})
	}
}
