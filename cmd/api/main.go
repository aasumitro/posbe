package main

import (
	"github.com/aasumitro/posbe/docs"
	"github.com/aasumitro/posbe/internal/_default"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/order"
	"github.com/aasumitro/posbe/internal/product"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

// @contact.name @aasumitro
// @contact.url https://aasumitro.id/
// @contact.email hello@aasumitro.id
// @license.name  MIT
// @license.url   https://github.com/aasumitro/posbe/blob/main/LICENSE

var (
	appConfig *config.Config
	appEngine *gin.Engine
)

func init() {
	initConfig()

	if !appConfig.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	appEngine = gin.Default()

	initSwaggerInfo()
}

func main() {
	// Init database connection
	appConfig.InitDbConn()
	// Load registered modules
	loadModules()
	// Defer close database
	appConfig.DeferCloseDbConn()
	// start server engine
	log.Fatal(appEngine.Run(appConfig.AppUrl))
}

func initConfig() {
	cfg, err := config.LoadConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(".env file not found!, please copy .env.example and paste as .env")
		} else {
			panic(err.Error())
		}
	}
	appConfig = cfg
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
	account.InitAccountModule(appConfig, appEngine)
	product.InitProductModule(appConfig, appEngine)
	order.InitOrderModule(appConfig, appEngine)
}
