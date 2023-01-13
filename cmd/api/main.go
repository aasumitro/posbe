package main

import (
	"context"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/docs"
	"github.com/aasumitro/posbe/internal/_default"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/catalog"
	"github.com/aasumitro/posbe/internal/store"
	"github.com/aasumitro/posbe/internal/transaction"
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
	configs.LoadConfig()

	// Init database connection
	configs.Cfg.InitDbConn()

	// Init cache connection
	configs.Cfg.InitRedisConn()

	if !configs.Cfg.AppDebug {
		configs.Cfg.InitCrashReporting()
		gin.SetMode(gin.ReleaseMode)
	}
}

func initEngine() {
	if configs.Cfg.AppDebug {
		accessLogFile, _ := os.Create("./temps/access.log")
		gin.DefaultWriter = io.MultiWriter(accessLogFile, os.Stdout)

		errorLogFile, _ := os.Create("./temps/errors.log")
		gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stdout)
	}

	appEngine = gin.Default()

	if !configs.Cfg.AppDebug {
		appEngine.Use(sentrygin.New(sentrygin.Options{}))
	}
}

func initSwaggerInfo() {
	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = configs.Cfg.AppName
	docs.SwaggerInfo.Description = configs.Cfg.AppDescription
	docs.SwaggerInfo.Version = configs.Cfg.AppVersion
	docs.SwaggerInfo.Host = configs.Cfg.AppURL
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func main() {
	// Load registered modules
	loadModules()
	// start engine
	log.Fatal(appEngine.Run(configs.Cfg.AppURL))
}

func loadModules() {
	_default.InitDefaultModule(appEngine)
	account.InitAccountModule(ctx, appEngine)
	store.InitStoreModule(ctx, appEngine)
	catalog.InitCatalogModule(ctx, appEngine)
	transaction.InitTransactionModule(ctx, appEngine)
}
