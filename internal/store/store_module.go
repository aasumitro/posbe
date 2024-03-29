package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/internal/store/handler/http"
	repository "github.com/aasumitro/posbe/internal/store/repository/sql"
	"github.com/aasumitro/posbe/internal/store/service"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"strconv"
)

var (
	floorRepo     domain.ICRUDRepository[domain.Floor]
	tableRepo     domain.ICRUDAddOnRepository[domain.Table]
	roomRepo      domain.ICRUDAddOnRepository[domain.Room]
	storePrefRepo domain.IStorePrefRepository
)

func InitStoreModule(ctx context.Context, router *gin.Engine) {
	floorRepo = repository.NewFloorSQLRepository()
	tableRepo = repository.NewTableSQLRepository()
	roomRepo = repository.NewRoomSQLRepository()
	storePrefRepo = repository.NewStorePrefSQLRepository()
	storeService := service.NewStoreService(ctx, floorRepo, tableRepo, roomRepo)
	storePrefService := service.NewStorePrefService(ctx, storePrefRepo)
	shouldCacheData(ctx)
	routerGroup := router.Group("v1")
	protectedRouter := routerGroup.
		Use(middleware.Auth()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewFloorHandler(storeService, protectedRouter)
	http.NewTableHandler(storeService, protectedRouter)
	http.NewRoomHandler(storeService, protectedRouter)
	http.NewStorePrefHandler(storePrefService, protectedRouter)
}

func shouldCacheData(ctx context.Context) {
	// run this at first booting
	if err := configs.RedisPool.
		Get(ctx, "store_prefs").
		Err(); err != nil && err == redis.Nil {
		if prefs, err := storePrefRepo.All(ctx); err == nil {
			setting := *prefs

			// store room status
			if room, _ := strconv.ParseBool(setting["feature_room"].(string)); room {
				if rooms, err := roomRepo.All(ctx); err == nil {
					for _, d := range rooms {
						key := fmt.Sprintf("room_%d_status", d.ID)
						configs.RedisPool.Set(ctx, key, 0, 0)
					}
				}
			}

			// store table status
			if table, _ := strconv.ParseBool(setting["feature_table"].(string)); table {
				if tables, err := tableRepo.All(ctx); err == nil {
					for _, d := range tables {
						key := fmt.Sprintf("table_%d_status", d.ID)
						configs.RedisPool.Set(ctx, key, 0, 0)
					}
				}
			}

			jsonData, _ := json.Marshal(prefs)
			// store data to redis
			configs.RedisPool.Set(ctx, "store_prefs", jsonData, 0)
		}
	}
}
