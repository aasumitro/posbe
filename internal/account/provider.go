package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func New(router *gin.RouterGroup) {
	repository := NewAccountRepository(config.PgxPool)
	as := NewAccountService(repository)
	shouldCacheData(context.Background(), repository)
	NewAuthHandler(as, router)
	authn := router.Group("")
	authn.Use(utils.AuthN())
	{
		NewRoleHandler(as, authn)
		NewUserHandler(as, authn)
	}
}

func shouldCacheData(ctx context.Context, repository IAccountRepository) {
	// at first booting validate roles
	err := config.RdpPool.Get(ctx, model.RolesCacheKey).Err()
	if errors.Is(err, redis.Nil) && err != nil {
		return
	}
	// if roles dint found then get data from database
	roles, err := repository.GetAllRoles(ctx)
	if err != nil {
		return
	}
	// encode data from storage
	jsonData, err := json.Marshal(roles)
	if err != nil {
		return
	}
	// store data to redis
	config.RdpPool.Set(ctx, model.RolesCacheKey, jsonData, 0)
}
