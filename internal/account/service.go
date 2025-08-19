package account

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/golang-jwt/jwt/v5"
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
			Key: model.UsersCacheKey, TTL: time.Minute * 30,
			CbF: func() ([]*model.User, error) {
				return service.repository.GetAllUsers(ctx)
			},
		},
	)

	return utils.HandleMultipleResults[model.User]("users", data, err)
}

func (service accountService) UserByID(
	ctx context.Context, id int,
) (*model.User, *utils.ServiceError) {
	data, err := service.repository.FindUserBy(ctx, model.FindWithID, id)

	return utils.HandleSingleResult[model.User]("user", data, err)
}

func (service accountService) CreateUser(
	ctx context.Context, data *model.User,
) (*model.User, *utils.ServiceError) {
	if strings.TrimSpace(data.Password) == "" {
		return nil, &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Password is required",
		}
	}

	pwd, err := utils.MakePassword(runtime.NumCPU(), data.Password)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	data.Password = pwd

	user, err := service.repository.InsertUser(ctx, *data)
	if err == nil {
		config.RdpPool.Del(ctx, model.UsersCacheKey)
	}

	return utils.HandleSingleResult[model.User]("user", user, err)
}

func (service accountService) UpdateUser(
	ctx context.Context, data *model.User,
) (user *model.User, errorData *utils.ServiceError) {
	data, err := service.repository.UpdateUserByID(ctx, *data)

	if err == nil {
		config.RdpPool.Del(ctx, model.UsersCacheKey)
	}

	return utils.HandleSingleResult[model.User]("user", data, err)
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
	valid, err := utils.ComparePassword(runtime.NumCPU(),
		user.Password, form.OldPassword)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if !valid {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "invalid password",
		}
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
	// get user data from database
	user, svcErr := service.getUserKV(ctx, model.FindWithID, data.ID)
	if svcErr != nil {
		return svcErr
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
) (user *model.User, errData *utils.ServiceError) {
	// get user data from database
	user, svcErr := service.getUserKV(ctx, model.FindWithUsername, form.Username)
	if svcErr != nil {
		return nil, svcErr
	}

	// validate password from input
	valid, err := utils.ComparePassword(runtime.NumCPU(), user.Password, form.Password)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if !valid {
		return nil, &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "invalid password",
		}
	}

	// jwt items
	secretKey := config.Instance.JWTSecretKey
	accessTokenDurationSecond := int64(3600)   // 1hr
	refreshTokenDurationSecond := int64(28800) // 8hrs
	claim := jwt.MapClaims{"id": user.ID, "role_id": user.Role.ID, "role_name": user.Role.Name}

	// generate access token
	accessToken, err := utils.NewJWT(claim, secretKey, accessTokenDurationSecond)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	user.AccessToken = accessToken

	// generate refresh token
	refreshToken, err := utils.NewJWT(claim, secretKey, refreshTokenDurationSecond)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	user.RefreshToken = refreshToken

	return user, nil
}

func (service accountService) RefreshToken(
	ctx context.Context, token string,
) (*model.User, *utils.ServiceError) {
	secretKey := config.Instance.JWTSecretKey
	accessTokenDurationSecond := int64(3600) // 1hr

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

	// generate access token
	claimMap := jwt.MapClaims{"id": user.ID, "role_id": user.Role.ID, "role_name": user.Role.Name}
	accessToken, err := utils.NewJWT(claimMap, secretKey, accessTokenDurationSecond)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	user.AccessToken = accessToken

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

func NewAccountService(repository IAccountRepository) IAccountService {
	return &accountService{repository: repository}
}
