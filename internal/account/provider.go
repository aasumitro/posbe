package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/account/handler/http"
	repository "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/account/service"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	userRepository model.ICRUDRepository[model.User]
	roleRepository model.ICRUDRepository[model.Role]
)

func NewAccountModuleProvider(router *gin.RouterGroup) {
	userRepository = repository.NewUserSQLRepository()
	roleRepository = repository.NewRoleSQLRepository()
	accountService := service.NewAccountService(
		roleRepository, userRepository)
	shouldCacheData(context.Background())
	http.NewAuthHandler(accountService, router)
	protectedRouter := router.
		Use(middleware.Auth()).
		Use(middleware.ActivityObserver())
	http.NewRoleHandler(accountService, protectedRouter)
	http.NewUserHandler(accountService, protectedRouter)
}

func shouldCacheData(ctx context.Context) {
	// run this at first booting
	if err := config.RedisPool.
		Get(ctx, "roles").
		Err(); errors.Is(err, redis.Nil) && err != nil {
		if roles, err := roleRepository.All(ctx); err == nil {
			// encode given data
			jsonData, _ := json.Marshal(roles)
			// store data to redis
			config.RedisPool.Set(ctx,
				"roles", jsonData, 0)
		}
	}
}
