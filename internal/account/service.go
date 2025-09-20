package account

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"runtime"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type accountService struct {
	repository IAccountRepository
}

func (service accountService) Roles(
	ctx context.Context,
) ([]*model.Role, *utils.ServiceError) {
	data, err := utils.CacheFirstData(ctx, config.RdpPool,
		&utils.CacheDataSupplied[[]*model.Role]{
			Key: model.RolesCacheKey, TTL: time.Hour * 1,
			CbF: func() ([]*model.Role, error) {
				return service.repository.GetAllRoles(ctx)
			},
		},
	)

	return utils.HandleMultipleResults[model.Role]("roles", data, err)
}

func (service accountService) Users(
	ctx context.Context,
) ([]*model.User, *utils.ServiceError) {
	data, err := utils.CacheFirstData(ctx, config.RdpPool,
		&utils.CacheDataSupplied[[]*model.User]{
			Key: model.UsersCacheKey, TTL: time.Minute * 10,
			CbF: func() ([]*model.User, error) {
				return service.repository.GetAllUsers(ctx)
			},
		},
	)

	return utils.HandleMultipleResults[model.User]("users", data, err)
}

func (service accountService) UserByID(
	ctx context.Context, id int64,
) (*model.User, *utils.ServiceError) {
	data, err := service.repository.FindUserBy(ctx, model.FindWithID, id)

	return utils.HandleSingleResult[model.User]("user", data, err)
}

func (service accountService) CreateUser(
	ctx context.Context, data *NewUserForm,
) (*model.User, *utils.ServiceError) {
	hashPassword, err := utils.MakePassword(runtime.NumCPU(), data.Password)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	newUser := model.User{RoleID: data.RoleID,
		Name: data.Name, Username: data.Username,
		Email: data.Email, Password: hashPassword}
	user, err := service.repository.InsertUser(ctx, newUser)
	if err == nil {
		config.RdpPool.Del(ctx, model.UsersCacheKey)
	}

	return utils.HandleSingleResult[model.User]("user", user, err)
}

func (service accountService) UpdateUser(
	ctx context.Context, data *UpdateUserForm,
) (*model.User, *utils.ServiceError) {
	user, err := service.ValidateAdmin(ctx, data.ID, data.RoleID)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	updateUser := model.User{ID: data.ID, RoleID: data.RoleID,
		Name: data.Name, Username: data.Username, Email: data.Email}
	user, err = service.repository.UpdateUserByID(ctx, updateUser)

	if err == nil {
		config.RdpPool.Del(ctx, model.UsersCacheKey)
	}

	return utils.HandleSingleResult[model.User]("user", user, err)
}

func (service accountService) UpdateUserPassword(
	ctx context.Context, form *UpdatePasswordForm,
) *utils.ServiceError {
	// get user data from database
	user, svcErr := service.getUserKV(ctx, model.FindWithID, form.ID)
	if svcErr != nil {
		return svcErr
	}

	// validate password from input
	if err := user.ComparePassword(form.Password); err != nil {
		svcErr = &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		if errors.Is(err, model.ErrInvalidCredential) {
			svcErr.Code = http.StatusBadRequest
		}
		return svcErr
	}

	// generate new password
	pwd, err := utils.MakePassword(runtime.NumCPU(), form.NewPassword)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// call repo for update action
	data := model.User{ID: form.ID, Password: pwd}
	if _, err := service.repository.UpdateUserByID(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service accountService) RemoveUser(
	ctx context.Context, data *model.User,
) *utils.ServiceError {
	user, err := service.ValidateAdmin(ctx, data.ID, 0)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	// call repo for delete action
	if err := service.repository.DeleteUserByID(ctx, *user); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// clear cache
	config.RdpPool.Del(ctx, model.UsersCacheKey)

	return nil
}

func (service accountService) AuthenticateUser(
	ctx context.Context, form *LoginForm,
) (*model.User, *utils.ServiceError) {
	// get user data from database
	user, svcErr := service.getUserKV(ctx, model.FindWithUsername, form.Username)
	if svcErr != nil {
		return nil, svcErr
	}

	// validate password from input
	if err := user.ComparePassword(form.Password); err != nil {
		svcErr = &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		if errors.Is(err, model.ErrInvalidCredential) {
			svcErr.Code = http.StatusBadRequest
		}
		return nil, svcErr
	}

	// generate tokens
	secretKey := config.Instance.JWTSecretKey
	if err := user.GenerateToken(secretKey, true); err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return user, nil
}

func (service accountService) RefreshToken(
	ctx context.Context, token string,
) (*model.User, *utils.ServiceError) {
	secretKey := config.Instance.JWTSecretKey

	claim, err := utils.ParseJWT(token, secretKey)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	userID, ok := claim["id"].(float64)
	if !ok {
		return nil, &utils.ServiceError{
			Code:    http.StatusUnauthorized,
			Message: "invalid token",
		}
	}

	user, svcErr := service.getUserKV(ctx, model.FindWithID, int(userID))
	if svcErr != nil {
		return nil, svcErr
	}

	// renew access token
	if err := user.GenerateToken(secretKey, false); err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return user, nil
}

func (service accountService) getUserKV(
	ctx context.Context, k model.FindWith, v any,
) (*model.User, *utils.ServiceError) {
	user, err := service.repository.FindUserBy(ctx, k, v)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: "user not found",
			}
		}

		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if user == nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusNotFound,
			Message: "user not found",
		}
	}

	return user, nil
}

func (service accountService) ValidateAdmin(
	ctx context.Context, id int64, roleID int,
) (*model.User, error) {
	// get user data from database
	user, svcErr := service.getUserKV(ctx, model.FindWithID, id)
	if svcErr != nil {
		return nil, errors.New("failed to get user by given id")
	}

	users, err := service.Users(ctx)
	if err != nil {
		return nil, errors.New("failed to get user list")
	}

	adminCount := 0
	for _, u := range users {
		if u.Role.Name == "admin" {
			adminCount++
		}
	}
	isAdmin := user.Role.Name == "admin"

	// Delete action
	if roleID == 0 {
		if isAdmin && adminCount <= 1 {
			return nil, errors.New("cannot delete this admin user")
		}
		return user, nil
	}

	// Update action
	if isAdmin && adminCount == 1 && roleID != user.RoleID {
		return nil, errors.New("cannot change this admin user role")
	}

	return user, nil
}

func NewAccountService(repository IAccountRepository) IAccountService {
	return &accountService{repository: repository}
}
