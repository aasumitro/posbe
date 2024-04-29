package config

import (
	"fmt"
	"log"
	"slices"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	lmgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	lsredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"go.uber.org/zap"
)

var allowOrigins = []string{
	"http://localhost:3000",
}

var allowHeaders = []string{
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
	"Authorization",
	"Cache-Control",
	"Origin",
	"Cookie",
}

func ServerEngine() Option {
	return func(cfg *Config) {
		engineOnce.Do(func() {
			log.Printf("Trying to init engine (GIN %s) . . . .",
				gin.Version)
			// set gin mode
			gin.SetMode(func() string {
				if cfg.AppDebug {
					return gin.DebugMode
				}
				return gin.ReleaseMode
			}())
			// set global variables
			GinEngine = gin.Default()
			// set cors middleware
			GinEngine.Use(cors.New(cors.Config{
				AllowOrigins:     allowOrigins,
				AllowMethods:     []string{"GET, POST, PATCH, DELETE"},
				AllowHeaders:     allowHeaders,
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				AllowOriginFunc: func(origin string) bool {
					return slices.Contains(allowOrigins, origin)
				},
				MaxAge: 12 * time.Hour,
			}))
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
				// setup rate limiter
				rate, err := limiter.NewRateFromFormatted("100-M")
				if err != nil {
					log.Fatalf("RATELIMITER_ERROR: %s", err.Error())
				}
				store, err := lsredis.NewStoreWithOptions(RedisPool,
					limiter.StoreOptions{Prefix: fmt.Sprintf(
						"%s{%s}", cfg.AppName, cfg.AppVersion)})
				if err != nil {
					log.Fatalf("RATELIMITER_ERROR: %s\n", err.Error())
				}
				GinEngine.ForwardedByClientIP = true
				GinEngine.Use(lmgin.NewMiddleware(limiter.New(store, rate)))
				// setup logger
				logger, err := zap.NewProduction()
				if err != nil {
					log.Fatalf("ZAP_LOGGER_ERROR: %s\n", err.Error())
				}
				defer func() { _ = logger.Sync() }()
				GinEngine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
				GinEngine.Use(ginzap.RecoveryWithZap(logger, true))
			}
		})
	}
}
