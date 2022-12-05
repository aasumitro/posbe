package store

import (
	"context"
	"github.com/aasumitro/posbe/internal/store/handler/http"
	repository "github.com/aasumitro/posbe/internal/store/repository/sql"
	"github.com/aasumitro/posbe/internal/store/service"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitStoreModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	floorRepo := repository.NewFloorSQLRepository(config.GetDbConn())
	tableRepo := repository.NewTableSQLRepository(config.GetDbConn())
	roomRepo := repository.NewRoomSQLRepository(config.GetDbConn())
	storePrefRepo := repository.NewStorePrefSQLRepository(config.GetDbConn())
	storeService := service.NewStoreService(ctx, floorRepo, tableRepo, roomRepo)
	storePrefService := service.NewStorePrefService(ctx, storePrefRepo)
	routerGroup := router.Group("v1")
	protectedRouter := routerGroup.
		Use(middleware.Auth(config.JWTSecretKey)).
		Use(middleware.ActivityObserver()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewFloorHandler(storeService, protectedRouter)
	http.NewTableHandler(storeService, protectedRouter)
	http.NewRoomHandler(storeService, protectedRouter)
	http.NewStorePrefHandler(storePrefService, protectedRouter)
}
