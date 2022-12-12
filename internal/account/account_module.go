package account

import (
	"context"
	"encoding/json"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/internal/account/handler/http"
	repository "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/account/service"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var (
	userRepository domain.ICRUDRepository[domain.User]
	roleRepository domain.ICRUDRepository[domain.Role]
)

func InitAccountModule(ctx context.Context, router *gin.Engine) {
	userRepository = repository.NewUserSQlRepository()
	roleRepository = repository.NewRoleSQlRepository()
	accountService := service.NewAccountService(ctx, roleRepository, userRepository)
	shouldCacheData(ctx)
	routerGroup := router.Group("v1")
	http.NewAuthHandler(accountService, routerGroup)
	protectedRouter := routerGroup.
		Use(middleware.Auth()).
		Use(middleware.ActivityObserver())
	http.NewRoleHandler(accountService, protectedRouter)
	http.NewUserHandler(accountService, protectedRouter)
}

func shouldCacheData(ctx context.Context) {
	// run this at first booting
	if err := config.RedisCache.
		Get(ctx, "roles").
		Err(); err != nil && err == redis.Nil {
		if roles, err := roleRepository.All(ctx); err == nil {
			// encode given data
			jsonData, _ := json.Marshal(roles)
			// store data to redis
			config.RedisCache.Set(ctx,
				"roles", jsonData, 0)
		}
	}
}
