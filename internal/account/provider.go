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

type IAccountRepository interface {
	GetAllRoles(ctx context.Context) ([]*model.Role, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	FindUserBy(ctx context.Context, key model.FindWith, val any) (*model.User, error)
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUserByID(ctx context.Context, user model.User) (*model.User, error)
	DeleteUserByID(ctx context.Context, user model.User) error
}

type IAccountService interface {
	Roles(ctx context.Context) ([]*model.Role, *utils.ServiceError)
	Users(ctx context.Context) ([]*model.User, *utils.ServiceError)
	UserByID(ctx context.Context, id int) (*model.User, *utils.ServiceError)
	CreateUser(ctx context.Context, data *NewUserForm) (*model.User, *utils.ServiceError)
	UpdateUser(ctx context.Context, data *UpdateUserForm) (*model.User, *utils.ServiceError)
	UpdateUserPassword(ctx context.Context, data *UpdatePasswordForm) *utils.ServiceError
	RemoveUser(ctx context.Context, data *model.User) *utils.ServiceError
	AuthenticateUser(ctx context.Context, form *LoginForm) (*model.User, *utils.ServiceError)
	RefreshToken(ctx context.Context, token string) (*model.User, *utils.ServiceError)
}

func NewAccountModuleProvider(router *gin.RouterGroup) {
	repository := NewAccountRepository(config.PgxPool)
	as := NewAccountService(repository)
	shouldCacheData(context.Background(), repository)
	NewAuthHandler(as, router)
	authn := router.Use(utils.AuthN())
	NewRoleHandler(as, authn)
	NewUserHandler(as, authn)
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
