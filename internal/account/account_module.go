package account

import (
	"context"
	"github.com/aasumitro/posbe/internal/account/handler/http"
	repository "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/account/service"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/gin-gonic/gin"
)

func InitAccountModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	userRepository := repository.NewUserSQlRepository(config.GetDbConn())
	roleRepository := repository.NewRoleSQlRepository(config.GetDbConn())
	accountService := service.NewAccountService(ctx, roleRepository, userRepository)
	routerGroup := router.Group("v1")
	http.NewAuthHandler(accountService, routerGroup)
	protectedRouter := routerGroup
	//protectedRoute.Use() // TODO ADD AUTH MIDDLEWARE
	http.NewRoleHandler(accountService, protectedRouter)
	http.NewUserHandler(accountService, protectedRouter)
}
