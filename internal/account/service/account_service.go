package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/errors"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
	"time"
)

type accountService struct {
	ctx      context.Context
	roleRepo domain.ICRUDRepository[domain.Role]
	userRepo domain.ICRUDRepository[domain.User]
	pwd      utils.IPassword
}

var (
	roleCacheKey = "roles"
)

func (service accountService) RoleList() (roles []*domain.Role, errorData *utils.ServiceError) {
	helper := utils.RedisCache{Ctx: service.ctx, RdpConn: configs.RedisPool}
	data, err := helper.CacheFirstData(&utils.CacheDataSupplied{
		Key: roleCacheKey,
		TTL: time.Hour * 1,
		CbF: func() (data any, err error) {
			return service.roleRepo.All(service.ctx)
		},
	})

	if data, ok := data.([]*domain.Role); ok {
		roles = data
	}

	if data, ok := data.(string); ok {
		var r []*domain.Role
		_ = json.Unmarshal([]byte(data), &r)
		roles = r
	}

	return utils.ValidateDataRows[domain.Role](roles, err)
}

func (service accountService) AddRole(data *domain.Role) (role *domain.Role, errorData *utils.ServiceError) {
	data, err := service.roleRepo.Create(service.ctx, data)

	configs.RedisPool.Del(service.ctx, roleCacheKey)

	return utils.ValidateDataRow[domain.Role](data, err)
}

func (service accountService) EditRole(data *domain.Role) (role *domain.Role, errorData *utils.ServiceError) {
	data, err := service.roleRepo.Update(service.ctx, data)

	configs.RedisPool.Del(service.ctx, roleCacheKey)

	return utils.ValidateDataRow[domain.Role](data, err)
}

func (service accountService) DeleteRole(data *domain.Role) *utils.ServiceError {
	role, err := service.roleRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if role.Usage >= 1 {
		return &utils.ServiceError{
			Code:    http.StatusForbidden,
			Message: errors.ErrorUnableToDelete,
		}
	}

	err = service.roleRepo.Delete(service.ctx, role)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	configs.RedisPool.Del(service.ctx, roleCacheKey)

	return nil
}

func (service accountService) UserList() (users []*domain.User, errorData *utils.ServiceError) {
	data, err := service.userRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.User](data, err)
}

func (service accountService) ShowUser(id int) (user *domain.User, errorData *utils.ServiceError) {
	data, err := service.userRepo.Find(service.ctx, domain.FindWithID, id)

	return utils.ValidateDataRow[domain.User](data, err)
}

func (service accountService) AddUser(data *domain.User) (user *domain.User, errorData *utils.ServiceError) {
	password := data.Password
	if password != "" {
		u := utils.Password{Stored: "", Supplied: password}
		pwd, err := u.HashPassword()
		if service.pwd != nil {
			pwd, err = service.pwd.HashPassword()
		}
		if err != nil {
			return nil, &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

		data.Password = pwd
	}

	data, err := service.userRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.User](data, err)
}

func (service accountService) EditUser(data *domain.User) (user *domain.User, errorData *utils.ServiceError) {
	password := data.Password
	if password != "" {
		u := utils.Password{Stored: "", Supplied: password}
		pwd, err := u.HashPassword()
		if service.pwd != nil {
			pwd, err = service.pwd.HashPassword()
		}
		if err != nil {
			return nil, &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

		data.Password = pwd
	}

	data, err := service.userRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.User](data, err)
}

func (service accountService) DeleteUser(data *domain.User) *utils.ServiceError {
	user, err := service.userRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	err = service.userRepo.Delete(service.ctx, user)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service accountService) VerifyUserCredentials(username, password string) (data any, errorData *utils.ServiceError) {
	user, err := service.userRepo.Find(service.ctx, domain.FindWithUsername, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}

		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	u := utils.Password{Stored: user.Password, Supplied: password}
	ok, err := u.ComparePasswords()
	if service.pwd != nil {
		ok, err = service.pwd.ComparePasswords()
	}
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if !ok {
		return nil, &utils.ServiceError{
			Code:    http.StatusUnprocessableEntity,
			Message: "Password Not Match",
		}
	}

	user.Password = ""

	return user, nil
}

func NewAccountService(
	ctx context.Context,
	roleRepo domain.ICRUDRepository[domain.Role],
	userRepo domain.ICRUDRepository[domain.User],
) domain.IAccountService {
	return &accountService{
		ctx:      ctx,
		roleRepo: roleRepo,
		userRepo: userRepo,
	}
}

// NewAccountServiceTest for testing purpose
func NewAccountServiceTest(
	ctx context.Context,
	roleRepo domain.ICRUDRepository[domain.Role],
	userRepo domain.ICRUDRepository[domain.User],
	pwd utils.IPassword,
) domain.IAccountService {
	return &accountService{
		ctx:      ctx,
		roleRepo: roleRepo,
		userRepo: userRepo,
		pwd:      pwd,
	}
}
