package domain

import "github.com/aasumitro/posbe/pkg/utils"

type (
	User struct {
		ID       int    `json:"id"`
		RoleId   int    `json:"role_id,omitempty"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"-"`
		Role     Role   `json:"role"`
	}

	Role struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
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
		//LoggedUserOut()
	}
)
