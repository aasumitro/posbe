package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/catalog"
	"github.com/aasumitro/posbe/internal/customer"
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
	// init event stream
	utils.InitEventStream(config.RedisPublisher,
		config.RedisSubscriber)
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
	utils.CloseEventStream(context.Background())
	if err := config.RedisCache.Close(); err != nil {
		log.Printf("Error shutting down redis connection: %v\n", err)
	}
	if err := config.RedisPublisher.Close(); err != nil {
		log.Printf("Error shutting down redis connection: %v\n", err)
	}
	if err := config.RedisSubscriber.Close(); err != nil {
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
	// stats route handler
	router.GET("/stats", func(ctx *gin.Context) {
		totalClients := 0
		if redisInfo, err := config.RedisPublisher.Info(
			ctx.Request.Context(), "clients",
		).Result(); err == nil {
			lines := strings.Split(redisInfo, "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "connected_clients:") {
					_, _ = fmt.Sscanf(line, "connected_clients:%d", &totalClients)
					break
				}
			}
		}

		channels := []string{order.SSEChannelKey}
		totalSubscribers := make(map[string]int64)
		if subscribers, err := config.RedisPublisher.PubSubNumSub(
			ctx.Request.Context(), channels...,
		).Result(); err == nil {
			totalSubscribers = subscribers
		}

		url := "https://www.google.com"
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer func() { _ = resp.Body.Close() }()
		elapsed := fmt.Sprintf("ping:google:%s", time.Since(start).String())

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		numGoroutines := runtime.NumGoroutine()

		pgStats := config.PgxPool.Stat()

		ctx.JSON(http.StatusOK, gin.H{
			"postgres": gin.H{
				"alive": config.PgxPool.Ping(ctx.Request.Context()) == nil,
				"metrics": gin.H{
					"total_conns":         pgStats.TotalConns(),
					"idle_conns":          pgStats.IdleConns(),
					"acquired_conns":      pgStats.AcquiredConns(),
					"max_conns":           pgStats.MaxConns(),
					"constructing_conns":  pgStats.ConstructingConns(),
					"acquire_count":       pgStats.AcquireCount(),
					"acquire_duration_ms": pgStats.AcquireDuration().Milliseconds(),
				},
			},
			"redis": gin.H{
				"alive": gin.H{
					"cache":      config.RedisCache.Ping(ctx.Request.Context()).Err() == nil,
					"publisher":  config.RedisPublisher.Ping(ctx.Request.Context()).Err() == nil,
					"subscriber": config.RedisSubscriber.Ping(ctx.Request.Context()).Err() == nil,
				},
				"metrics": gin.H{
					"total_connections":       totalClients,
					"subscribers_per_channel": totalSubscribers,
				},
			},
			"runtime": gin.H{
				"num_goroutines": numGoroutines,
				"alloc":          fmt.Sprintf("%dMB", m.Alloc/(1024*1024)),
				"total_alloc":    fmt.Sprintf("%dMB", m.TotalAlloc/(1024*1024)),
				"sys":            fmt.Sprintf("%dMB", m.Sys/(1024*1024)),
				"heap_alloc":     fmt.Sprintf("%dMB", m.HeapAlloc/(1024*1024)),
				"heap_sys":       fmt.Sprintf("%dMB", m.HeapSys/(1024*1024)),
				"stack_inuse":    fmt.Sprintf("%dMB", m.StackInuse/(1024*1024)),
				"stack_sys":      fmt.Sprintf("%dMB", m.StackSys/(1024*1024)),
				"mallocs":        fmt.Sprintf("%dKB", m.Mallocs/(1024*1024)),
				"frees":          fmt.Sprintf("%dKB", m.Frees/(1024)),
			},
			"network": elapsed,
		})
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
		customer.New(v1)
		// report.New(v1)
	}
}
