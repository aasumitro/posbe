package service

import (
	"context"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/errors"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
)

// TODO
// ADD DATA CACHE
// 1. SEARCH FROM CACHE
// 2. DATA IS NOT CACHED
// 3. GET FROM STORAGE
// 4. STORE TO CACHE
// 5. RETURN DATA

// TODO WRAP IN-MEMORY & IN-STORAGE DB;
// getDataFromCacheRepo(ctx, func() {
//    return getDataFromStorageRepo()
//})
// if data not found on memory
// then load data from storage
// store data from storage to memory
// then return data to user

type accountService struct {
	ctx      context.Context
	roleRepo domain.ICRUDRepository[domain.Role]
	userRepo domain.ICRUDRepository[domain.User]
}

func (service accountService) RoleList() (roles []*domain.Role, errorData *utils.ServiceError) {
	data, err := service.roleRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Role](data, err)
}

func (service accountService) AddRole(data *domain.Role) (role *domain.Role, errorData *utils.ServiceError) {
	data, err := service.roleRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Role](data, err)
}

func (service accountService) EditRole(data *domain.Role) (role *domain.Role, errorData *utils.ServiceError) {
	data, err := service.roleRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Role](data, err)
}

func (service accountService) DeleteRole(data *domain.Role) *utils.ServiceError {
	role, err := service.roleRepo.Find(service.ctx, domain.FindWithId, data.ID)
	if err != nil {
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

	return nil
}

func (service accountService) UserList() (users []*domain.User, errorData *utils.ServiceError) {
	data, err := service.userRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.User](data, err)
}

func (service accountService) ShowUser(id int) (user *domain.User, errorData *utils.ServiceError) {
	data, err := service.userRepo.Find(service.ctx, domain.FindWithId, id)

	return utils.ValidateDataRow[domain.User](data, err)
}

func (service accountService) AddUser(data *domain.User) (user *domain.User, errorData *utils.ServiceError) {
	password := data.Password
	if password != "" {
		pwd, err := utils.HashPassword(password)
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
		pwd, err := utils.HashPassword(password)
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
	user, err := service.userRepo.Find(service.ctx, domain.FindWithId, data.ID)
	if err != nil {
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

//func (service accountService) VerifyUserCredentials(username, password string) (data any, errorData *utils.ServiceError) {
//	//TODO implement me
//	panic("implement me")
//}

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
