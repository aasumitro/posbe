package main

import (
	"context"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal"
	"github.com/spf13/viper"

	// swagger embed files
	_ "github.com/aasumitro/posbe/docs"
)

// @version                     0.0.1-dev
// @title                       POSBE API Specs documentation
// @description                 This is an auto-generated API Docs.
//
// @contact.name                @aasumitro
// @contact.url                 https://aasumitro.id
// @contact.email               hello@aasumitro.id
//
// @securityDefinitions.apikey  AuthToken
// @in                          header
// @name                        Authorization
//
// @license.name  MIT
// @license.url   https://github.com/aasumitro/posbe/blob/main/LICENSE

func main() {
	viper.SetConfigFile(".env")
	// viper.AutomaticEnv()
	mainCtx := context.Background()
	// load environment file
	config.LoadWith(mainCtx,
		config.SentryConnection(),
		config.PostgresConnection(),
		config.RedisConnection(),
		config.ServerEngine())
	// run server app
	internal.RunServer(mainCtx)
}
