package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/store/handler/http"
	repository "github.com/aasumitro/posbe/internal/store/repository/sql"
	"github.com/aasumitro/posbe/internal/store/service"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	floorRepo     model.ICRUDRepository[model.Floor]
	tableRepo     model.ICRUDAddOnRepository[model.Table]
	roomRepo      model.ICRUDAddOnRepository[model.Room]
	storePrefRepo model.IStorePrefRepository
)

func NewStoreModuleProvider(router *gin.RouterGroup) {
	floorRepo = repository.NewFloorSQLRepository()
	tableRepo = repository.NewTableSQLRepository()
	roomRepo = repository.NewRoomSQLRepository()
	storePrefRepo = repository.NewStorePrefSQLRepository()
	storeService := service.NewStoreService(floorRepo, tableRepo, roomRepo)
	storePrefService := service.NewStorePrefService(storePrefRepo)
	shouldCacheData(context.Background())
	protectedRouter := router.
		Use(middleware.Auth()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewFloorHandler(storeService, protectedRouter)
	http.NewTableHandler(storeService, protectedRouter)
	http.NewRoomHandler(storeService, protectedRouter)
	http.NewStorePrefHandler(storePrefService, protectedRouter)
}

func shouldCacheData(ctx context.Context) {
	// run this at first booting
	if err := config.RedisPool.
		Get(ctx, "store_prefs").
		Err(); err != nil && errors.Is(err, redis.Nil) {
		if prefs, err := storePrefRepo.All(ctx); err == nil {
			setting := *prefs

			// store room status
			if room, _ := strconv.ParseBool(setting["feature_room"].(string)); room {
				if rooms, err := roomRepo.All(ctx); err == nil {
					for _, d := range rooms {
						key := fmt.Sprintf("room_%d_status", d.ID)
						config.RedisPool.Set(ctx, key, 0, 0)
					}
				}
			}

			// store table status
			if table, _ := strconv.ParseBool(setting["feature_table"].(string)); table {
				if tables, err := tableRepo.All(ctx); err == nil {
					for _, d := range tables {
						key := fmt.Sprintf("table_%d_status", d.ID)
						config.RedisPool.Set(ctx, key, 0, 0)
					}
				}
			}

			jsonData, _ := json.Marshal(prefs)
			// store data to redis
			config.RedisPool.Set(ctx, "store_prefs", jsonData, 0)
		}
	}
}
