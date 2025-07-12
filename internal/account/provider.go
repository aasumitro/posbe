package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aasumitro/posbe/config"
	repository "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/middleware"
	"github.com/aasumitro/posbe/internal/model"
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
	as := NewAccountService(
		roleRepository, userRepository)
	shouldCacheData(context.Background())
	NewAuthHandler(as, router)
	protectedRouter := router.Use(middleware.Auth())
	NewRoleHandler(as, protectedRouter)
	NewUserHandler(as, protectedRouter)
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
