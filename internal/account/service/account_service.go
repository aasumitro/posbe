package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/aasumitro/posbe/common"
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
)

type accountService struct {
	roleRepo model.ICRUDRepository[model.Role]
	userRepo model.ICRUDRepository[model.User]
	pwd      utils.IPassword
}

var roleCacheKey = "roles"

func (service accountService) RoleList(
	ctx context.Context,
) (
	roles []*model.Role,
	errorData *utils.ServiceError,
) {
	helper := utils.RedisCache{Ctx: ctx, RdpConn: config.RedisPool}
	data, err := helper.CacheFirstData(&utils.CacheDataSupplied{
		Key: roleCacheKey,
		TTL: time.Hour * 1,
		CbF: func() (data any, err error) {
			return service.roleRepo.All(ctx)
		},
	})
	if data, ok := data.([]*model.Role); ok {
		roles = data
	}
	if data, ok := data.(string); ok {
		var r []*model.Role
		_ = json.Unmarshal([]byte(data), &r)
		roles = r
	}
	return utils.ValidateDataRows[model.Role](roles, err)
}

func (service accountService) AddRole(
	ctx context.Context,
	item *model.Role,
) (
	role *model.Role,
	errorData *utils.ServiceError,
) {
	data, err := service.roleRepo.Create(ctx, item)
	config.RedisPool.Del(ctx, roleCacheKey)
	return utils.ValidateDataRow[model.Role](data, err)
}

func (service accountService) EditRole(
	ctx context.Context,
	item *model.Role,
) (
	role *model.Role,
	errorData *utils.ServiceError,
) {
	data, err := service.roleRepo.Update(ctx, item)
	config.RedisPool.Del(ctx, roleCacheKey)
	return utils.ValidateDataRow[model.Role](data, err)
}

func (service accountService) DeleteRole(
	ctx context.Context,
	data *model.Role,
) *utils.ServiceError {
	role, err := service.roleRepo.Find(ctx, model.FindWithID, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
			Message: common.ErrorUnableToDelete,
		}
	}
	err = service.roleRepo.Delete(ctx, role)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	config.RedisPool.Del(ctx, roleCacheKey)
	return nil
}

func (service accountService) UserList(
	ctx context.Context,
) (
	users []*model.User,
	errorData *utils.ServiceError,
) {
	data, err := service.userRepo.All(ctx)
	return utils.ValidateDataRows[model.User](data, err)
}

func (service accountService) ShowUser(
	ctx context.Context,
	id int,
) (
	user *model.User,
	errorData *utils.ServiceError,
) {
	data, err := service.userRepo.Find(ctx, model.FindWithID, id)
	return utils.ValidateDataRow[model.User](data, err)
}

func (service accountService) AddUser(
	ctx context.Context,
	data *model.User,
) (
	user *model.User,
	errorData *utils.ServiceError,
) {
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
	data, err := service.userRepo.Create(ctx, data)
	return utils.ValidateDataRow[model.User](data, err)
}

func (service accountService) EditUser(
	ctx context.Context,
	data *model.User,
) (
	user *model.User,
	errorData *utils.ServiceError,
) {
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
	data, err := service.userRepo.Update(ctx, data)
	return utils.ValidateDataRow[model.User](data, err)
}

func (service accountService) DeleteUser(
	ctx context.Context,
	data *model.User,
) *utils.ServiceError {
	user, err := service.userRepo.Find(ctx, model.FindWithID, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
	err = service.userRepo.Delete(ctx, user)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service accountService) VerifyUserCredentials(
	ctx context.Context,
	username, password string,
) (
	data any,
	errorData *utils.ServiceError,
) {
	user, err := service.userRepo.Find(ctx, model.FindWithUsername, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
	roleRepo model.ICRUDRepository[model.Role],
	userRepo model.ICRUDRepository[model.User],
) model.IAccountService {
	return &accountService{
		roleRepo: roleRepo,
		userRepo: userRepo,
	}
}

// NewAccountServiceTest for testing purpose
func NewAccountServiceTest(
	roleRepo model.ICRUDRepository[model.Role],
	userRepo model.ICRUDRepository[model.User],
	pwd utils.IPassword,
) model.IAccountService {
	return &accountService{
		roleRepo: roleRepo,
		userRepo: userRepo,
		pwd:      pwd,
	}
}
