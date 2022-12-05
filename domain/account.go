package domain

import (
	"github.com/aasumitro/posbe/pkg/utils"
)

type (
	LoginForm struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	User struct {
		ID       int    `json:"id"`
		RoleId   int    `json:"role_id,omitempty" form:"role_id" binding:"required"`
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
		RoleList() (roles []*Role, errData *utils.ServiceError)
		AddRole(data *Role) (role *Role, errData *utils.ServiceError)
		EditRole(data *Role) (role *Role, errData *utils.ServiceError)
		DeleteRole(data *Role) *utils.ServiceError

		UserList() (users []*User, errData *utils.ServiceError)
		ShowUser(id int) (user *User, errData *utils.ServiceError)
		AddUser(data *User) (user *User, errData *utils.ServiceError)
		EditUser(data *User) (user *User, errData *utils.ServiceError)
		DeleteUser(data *User) *utils.ServiceError

		VerifyUserCredentials(username, password string) (data any, errData *utils.ServiceError)
	}
)
