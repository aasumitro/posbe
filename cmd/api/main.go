package main

import (
	"context"
	"github.com/aasumitro/posbe/docs"
	"github.com/aasumitro/posbe/internal/_default"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/catalog"
	"github.com/aasumitro/posbe/internal/store"
	"github.com/aasumitro/posbe/internal/transaction"
	"github.com/aasumitro/posbe/pkg/config"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

// @contact.name @aasumitro
// @contact.url https://aasumitro.id/
// @contact.email hello@aasumitro.id
// @license.name  MIT
// @license.url   https://github.com/aasumitro/posbe/blob/main/LICENSE

var (
	appEngine *gin.Engine
	ctx       = context.Background()
)

func init() {
	initConfig()

	initEngine()

	initSwaggerInfo()
}

func initConfig() {
	// Load app environment
	config.LoadConfig()

	// Init database connection
	config.Cfg.InitDbConn()

	// Init cache connection
	config.Cfg.InitRedisConn()

	if !config.Cfg.AppDebug {
		config.Cfg.InitCrashReporting()
		gin.SetMode(gin.ReleaseMode)
	}
}

func initEngine() {
	if config.Cfg.AppDebug {
		accessLogFile, _ := os.Create("./temps/access.log")
		gin.DefaultWriter = io.MultiWriter(accessLogFile, os.Stdout)

		errorLogFile, _ := os.Create("./temps/errors.log")
		gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stdout)
	}

	appEngine = gin.Default()

	if !config.Cfg.AppDebug {
		appEngine.Use(sentrygin.New(sentrygin.Options{}))
	}
}

func initSwaggerInfo() {
	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = config.Cfg.AppName
	docs.SwaggerInfo.Description = config.Cfg.AppDescription
	docs.SwaggerInfo.Version = config.Cfg.AppVersion
	docs.SwaggerInfo.Host = config.Cfg.AppUrl
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func main() {
	// Load registered modules
	loadModules()
	// start engine
	log.Fatal(appEngine.Run(config.Cfg.AppUrl))
}

func loadModules() {
	_default.InitDefaultModule(appEngine)
	account.InitAccountModule(ctx, appEngine)
	store.InitStoreModule(ctx, appEngine)
	catalog.InitCatalogModule(ctx, appEngine)
	transaction.InitTransactionModule(ctx, appEngine)
}
