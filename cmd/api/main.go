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
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	appConfig *config.Config
	appEngine *gin.Engine
	ctx       = context.Background()
)

func init() {
	initConfig()

	initEngine()

	initSwaggerInfo()
}

func main() {
	// Init database connection
	appConfig.InitDbConn()
	// Load registered modules
	loadModules()
	// start engine
	log.Fatal(appEngine.Run(appConfig.AppUrl))
}

func initConfig() {
	cfg, err := config.LoadConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal(".env file not found!, please copy .env.example and paste as .env")
		}

		log.Fatal(err.Error())
	}

	appConfig = cfg
}

func initEngine() {
	if !appConfig.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	accessLogFile, _ := os.Create("./temps/access.log")
	gin.DefaultWriter = io.MultiWriter(accessLogFile, os.Stdout)

	errorLogFile, _ := os.Create("./temps/errors.log")
	gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stdout)

	appEngine = gin.Default()
}

func initSwaggerInfo() {
	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = appConfig.AppName
	docs.SwaggerInfo.Description = appConfig.AppDescription
	docs.SwaggerInfo.Version = appConfig.AppVersion
	docs.SwaggerInfo.Host = appConfig.AppUrl
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func loadModules() {
	_default.InitDefaultModule(appEngine)
	account.InitAccountModule(ctx, appConfig, appEngine)
	store.InitStoreModule(ctx, appConfig, appEngine)
	catalog.InitCatalogModule(ctx, appConfig, appEngine)
	transaction.InitTransactionModule(ctx, appConfig, appEngine)
}
