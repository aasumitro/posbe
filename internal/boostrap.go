package internal

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/catalog"
	"github.com/aasumitro/posbe/internal/order"
	"github.com/aasumitro/posbe/internal/store"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer(ctx context.Context) {
	// Create a context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(ctx,
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// router engine
	routerEngine := config.GinEngine
	// register public routes
	registerPublicRoutes(routerEngine)
	// register providers
	registerAPIModules(routerEngine)
	// server defines parameters for running an HTTP server.
	server := &http.Server{
		Addr:              config.Instance.AppPort,
		Handler:           routerEngine,
		ReadHeaderTimeout: time.Second * 10,
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
	config.PgxPool.Close()
	// Close redis connections
	if err := config.RdpPool.Close(); err != nil {
		log.Printf("Error shutting down redis connection: %v\n", err)
	}
	// notify user of shutdown
	log.Println("Server exiting")
}

func registerPublicRoutes(engine *gin.Engine) {
	router := engine
	// no route handler
	router.NoMethod(func(ctx *gin.Context) {
		ctx.String(http.StatusBadRequest,
			"HTTP_METHOD_NOT_FOUND")
	})
	// no route handler
	router.NoRoute(func(ctx *gin.Context) {
		if strings.Contains(ctx.FullPath(), "/api/") {
			ctx.String(http.StatusNotFound,
				"route that you are looking for is not found")
			return
		}
	})
	// main route handler
	router.GET(utils.EmptyPath, func(ctx *gin.Context) {
		ctx.String(http.StatusOK,
			"hello world")
		return
	})
	// asset route
	router.Static("/assets", "./uploads")
	// swagger docs routes
	router.GET("/api-specs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(4)))
}

func registerAPIModules(engine *gin.Engine) {
	v1Path := "api/v1"
	v1 := engine.Group(v1Path)
	{
		account.New(v1)
		store.New(v1)
		catalog.New(v1)
		order.New(v1)
		// report.New(v1)
	}
}
