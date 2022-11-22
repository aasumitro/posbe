package main

import (
	"github.com/aasumitro/posbe/docs"
	"github.com/aasumitro/posbe/internal/_default"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/order"
	"github.com/aasumitro/posbe/internal/product"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

// @contact.name @aasumitro
// @contact.url https://aasumitro.id/
// @contact.email hello@aasumitro.id
// @license.name  MIT
// @license.url   https://github.com/aasumitro/posbe/blob/main/LICENSE

var (
	appConfig *config.Config
	ginEngine *gin.Engine
)

func init() {
	appConfig = &config.Config{}
	appConfig.InitDbConn()

	if !appConfig.GetAppDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine = gin.Default()

	docs.SwaggerInfo.BasePath = ginEngine.BasePath()
	docs.SwaggerInfo.Title = appConfig.GetAppName()
	docs.SwaggerInfo.Description = appConfig.GetAppDesc()
	docs.SwaggerInfo.Version = appConfig.GetAppVersion()
	docs.SwaggerInfo.Host = appConfig.GetAppUrl()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func main() {
	addModule()
	appConfig.DeferCloseDbConn()
	log.Fatal(ginEngine.Run(appConfig.GetAppUrl()))
}

func addModule() {
	_default.InitDefaultModule(ginEngine)
	account.InitAccountModule(appConfig, ginEngine)
	product.InitProductModule(appConfig, ginEngine)
	order.InitOrderModule(appConfig, ginEngine)
}
