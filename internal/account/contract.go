package account

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
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
