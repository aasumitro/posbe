package account

import (
	"context"
	"github.com/aasumitro/posbe/internal/account/handler/http"
	repository "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/account/service"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitAccountModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	userRepository := repository.NewUserSQlRepository(config.GetDbConn())
	roleRepository := repository.NewRoleSQlRepository(config.GetDbConn())
	accountService := service.NewAccountService(ctx, roleRepository, userRepository)
	routerGroup := router.Group("v1")
	http.NewAuthHandler(accountService, config, routerGroup)
	protectedRouter := routerGroup.
		Use(middleware.Auth(config.JWTSecretKey)).
		Use(middleware.ActivityObserver()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewRoleHandler(accountService, protectedRouter)
	http.NewUserHandler(accountService, protectedRouter)
}
