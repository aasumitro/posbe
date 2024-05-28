package internal

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/aasumitro/posbe/common"
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/catalog"
	"github.com/aasumitro/posbe/internal/store"
	"github.com/aasumitro/posbe/internal/transaction"
	"github.com/aasumitro/posbe/web"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	healthcheckconfig "github.com/tavsec/gin-healthcheck/config"
)

type embeddedFile struct {
	fs.File
}

func (f *embeddedFile) Close() error {
	return nil
}

func (f *embeddedFile) Seek(offset int64, whence int) (int64, error) {
	return f.File.(io.Seeker).Seek(offset, whence)
}

func RunServer(ctx context.Context) {
	// Create a context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(ctx,
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// router engine
	routerEngine := config.GinEngine
	// register public routes
	registerPublicRoutes(ctx, routerEngine)
	// register providers
	registerAPIModuleV1(routerEngine)
	// server defines parameters for running an HTTP server.
	server := &http.Server{
		Addr:              config.Instance.AppPort,
		Handler:           routerEngine,
		ReadHeaderTimeout: time.Second * common.ServerReadTimeout,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(
			err, http.ErrServerClosed,
		) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	timeToHandle := 10
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(timeToHandle)*time.Second)
	defer cancel()
	// Shutdown server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %s\n", err)
	}
	// Close database connections
	if err := config.PostgresPool.Close(); err != nil {
		log.Printf("Error disconnect mongodb connection: %v\n", err)
	}
	// Close redis connections
	if err := config.RedisPool.Close(); err != nil {
		log.Printf("Error shutting down redis connection: %v\n", err)
	}
	// notify user of shutdown
	log.Println("Server exiting")
}

func registerPublicRoutes(
	sgCtx context.Context,
	engine *gin.Engine,
) {
	router := engine
	// no route handler
	router.NoMethod(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound,
			"HTTP_METHOD_NOT_FOUND")
	})
	// no route handler
	router.NoRoute(func(ctx *gin.Context) {
		if !strings.Contains(ctx.FullPath(), "/fe/") {
			ctx.String(http.StatusNotFound,
				"route that you are looking for is not found")
			return
		}
		file, err := web.SPAAssets().Open("index.html")
		if err != nil {
			ctx.String(http.StatusInternalServerError,
				"failed to open spa file: ", err.Error())
			return
		}
		defer func() { _ = file.Close() }()
		fileInfo, err := file.Stat()
		if err != nil {
			ctx.String(http.StatusInternalServerError,
				"failed to get spa file info: ", err.Error())
			return
		}
		http.ServeContent(
			ctx.Writer, ctx.Request, fileInfo.Name(),
			fileInfo.ModTime(), &embeddedFile{file})
	})
	// main route handler
	router.GET(common.EmptyPath, func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/fe")
	})
	// client (web ui) route handler
	router.StaticFS("/fe", http.FS(web.SPAAssets()))
	// swagger docs routes
	router.GET("/api-specs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(
				common.SwaggerDefaultModelsExpandDepth)))
	// health check routes
	redisCheck := checks.NewRedisCheck(config.RedisPool)
	healthConfig := healthcheckconfig.DefaultConfig()
	healthConfig.HealthPath = "/health"
	_ = healthcheck.New(router, healthConfig, []checks.Check{
		&redisCheck, checks.NewContextCheck(sgCtx, "signals"),
		checks.NewPingCheck("https://www.google.com",
			"GET", common.HealthCheckPingTimeout, nil, nil),
		checks.SqlCheck{Sql: config.PostgresPool},
	})
}

func registerAPIModuleV1(engine *gin.Engine) {
	routerGroup := engine.Group("api/v1")
	account.NewAccountModuleProvider(routerGroup)
	store.NewStoreModuleProvider(routerGroup)
	catalog.NewCatalogModuleProvider(routerGroup)
	transaction.NewTransactionModuleProvider(routerGroup)
}
