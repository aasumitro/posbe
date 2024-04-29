package model

import (
	"context"

	"github.com/aasumitro/posbe/pkg/utils"
)

type (
	LoginForm struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	User struct {
		ID       int    `json:"id"`
		RoleID   int    `json:"role_id,omitempty" form:"role_id" binding:"required"`
		Name     string `json:"name" form:"name" binding:"required"`
		Username string `json:"username" form:"username" binding:"required"`
		Email    string `json:"email" form:"email"`
		Phone    string `json:"phone" form:"phone"`
		Password string `json:"password,omitempty" form:"password" binding:"required"`
		Role     Role   `json:"role" binding:"-"`
	}

	Role struct {
		ID          int    `json:"id"`
		Name        string `json:"name" form:"name" binding:"required"`
		Description string `json:"description" form:"description"  binding:"required"`
		Usage       int    `json:"usage,omitempty"`
	}

	// IAccountService contract
	IAccountService interface {
		RoleList(ctx context.Context) (roles []*Role, errData *utils.ServiceError)
		AddRole(ctx context.Context, data *Role) (role *Role, errData *utils.ServiceError)
		EditRole(ctx context.Context, data *Role) (role *Role, errData *utils.ServiceError)
		DeleteRole(ctx context.Context, data *Role) *utils.ServiceError

		UserList(ctx context.Context) (users []*User, errData *utils.ServiceError)
		ShowUser(ctx context.Context, id int) (user *User, errData *utils.ServiceError)
		AddUser(ctx context.Context, data *User) (user *User, errData *utils.ServiceError)
		EditUser(ctx context.Context, data *User) (user *User, errData *utils.ServiceError)
		DeleteUser(ctx context.Context, data *User) *utils.ServiceError

		VerifyUserCredentials(ctx context.Context, username, password string) (data any, errData *utils.ServiceError)
	}
)
