package account_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/aasumitro/posbe/mocks"
	"github.com/alicebob/miniredis/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type accountTestSuite struct {
	suite.Suite
	role   *model.Role
	roles  []*model.Role
	user   *model.User
	users  []*model.User
	svcErr *utils.ServiceError
}

func (suite *accountTestSuite) SetupSuite() {
	suite.role = &model.Role{
		ID:          1,
		Name:        "lorem",
		Description: "lorem ipsum",
		Usage:       1,
	}

	suite.roles = []*model.Role{
		suite.role,
		{
			ID:          2,
			Name:        "dolor",
			Description: "Dolor Sit Amet",
		},
	}

	suite.user = &model.User{
		ID:       1,
		RoleID:   1,
		Name:     "lorem ipsum",
		Username: "lorem",
		Email:    "lorem@ipsum.id",
		Role:     *suite.role,
		Password: "4c388700ed16de8681d5f5b4785c94f60af20dd765a5acade2b0f8e86357315a.dda885bab15f14eebdeabfde43173ee0de14db0a3a8e65438974678a4b3a5135",
	}

	suite.users = []*model.User{
		suite.user,
		{
			ID:       2,
			RoleID:   1,
			Name:     "dolor amet",
			Username: "dolor",
			Email:    "dolor@amet.id",
			Role:     *suite.role,
			Password: "4c388700ed16de8681d5f5b4785c94f60af20dd765a5acade2b0f8e86357315a.dda885bab15f14eebdeabfde43173ee0de14db0a3a8e65438974678a4b3a5135",
		},
	}

	suite.svcErr = &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	}

	config.RdpPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})

	viper.SetConfigFile("../../misc/conf/.env.example")
	viper.SetConfigType("dotenv")
	config.LoadWith(context.Background())
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess_ReturnModel() {
	cacheMock := new(mocks.MockICacheUtil[model.Role])
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.
		On("GetAllRoles", mock.Anything).
		Return(suite.roles, nil).Once()
	cacheMock.On("CacheFirstData", mock.Anything).
		Return(suite.roles, nil).Once()
	data, err := accSvc.Roles(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess_ReturnString() {
	cacheMock := new(mocks.MockICacheUtil[model.Role])
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("GetAllRoles", mock.Anything).
		Return(nil, nil).Once()
	jsonData, err := json.Marshal(suite.roles)
	require.Nil(suite.T(), err)
	config.RdpPool.Set(context.TODO(), "roles", jsonData, 1)
	cacheMock.On("CacheFirstData", config.RdpPool, &utils.CacheDataSupplied[[]*model.Role]{
		Key: "roles",
		TTL: time.Hour * 1,
		CbF: nil,
	}).Return(jsonData, nil).Once()
	data, svcErr := accSvc.Roles(context.TODO())
	suite.T().Log(data)
	suite.T().Log(svcErr)
	require.Nil(suite.T(), svcErr)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
}
func (suite *accountTestSuite) TestAccountService_RoleList_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("GetAllRoles", mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.Roles(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_UserList_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("GetAllUsers", mock.Anything).
		Once().Return(suite.users, nil)
	data, err := accSvc.Users(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_UserList_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("GetAllUsers", mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.Users(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_ShowUser_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	data, err := accSvc.UserByID(context.TODO(), suite.user.ID)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.user)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_ShowUser_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.UserByID(context.TODO(), suite.user.ID)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("InsertUser", mock.Anything, mock.Anything).
		Once().Return(suite.users[1], nil)
	data, err := accSvc.CreateUser(context.TODO(), &account.NewUserForm{
		RoleID:   1,
		Name:     "Test User",
		Username: "test_user",
		Email:    "test@user.id",
		Password: "secret",
	})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users[1])
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError_Password() {
	suite.T().Skip() // need to inject pwd mock before run this
	repoMock := new(mocks.MockIAccountRepository)
	pwdMock := new(mocks.MockIPasswordUtil)
	accSvc := account.NewAccountService(repoMock)
	pwdMock.
		On("HashPassword").
		Return("", errors.New("UNEXPECTED")).
		Once()
	data, err := accSvc.CreateUser(context.TODO(), &account.NewUserForm{
		RoleID:   1,
		Name:     "Test User",
		Username: "test_user",
		Email:    "test@user.id",
		Password: "secret",
	})
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	pwdMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("InsertUser", mock.Anything, mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.CreateUser(context.TODO(), &account.NewUserForm{
		RoleID:   1,
		Name:     "Test User",
		Username: "test_user",
		Email:    "test@user.id",
		Password: "secret",
	})
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("UpdateUserByID", mock.Anything, mock.Anything).
		Once().Return(suite.users[1], nil)
	data, err := accSvc.UpdateUser(context.TODO(), &account.UpdateUserForm{
		ID:       1,
		RoleID:   1,
		Name:     "Test User",
		Username: "test_user",
		Email:    "test@user.id",
	})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users[1])
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError_Password() {
	suite.T().Skip() // need to inject pwd mock before run this
	pwdMock := new(mocks.MockIPasswordUtil)
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	pwdMock.On("HashPassword").
		Return("", errors.New("UNEXPECTED")).Once()
	data, err := accSvc.UpdateUser(context.TODO(), &account.UpdateUserForm{
		ID:       1,
		RoleID:   1,
		Name:     "Test User",
		Username: "test_user",
		Email:    "test@user.id",
	})
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("UpdateUserByID", mock.Anything, mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.UpdateUser(context.TODO(), &account.UpdateUserForm{
		ID:       1,
		RoleID:   1,
		Name:     "Test User",
		Username: "test_user",
		Email:    "test@user.id",
	})
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditPassword_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	repoMock.On("UpdateUserByID", mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	err := accSvc.UpdateUserPassword(context.TODO(), &account.UpdatePasswordForm{
		ID:          1,
		Password:    "secret",
		NewPassword: "new_password",
	})
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_EditPassword_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	suite.T().Run("failed get user", func(t *testing.T) {
		repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
			Once().Return(nil, sql.ErrNoRows)
		err := accSvc.UpdateUserPassword(context.TODO(), &account.UpdatePasswordForm{
			ID:          1,
			Password:    "secret",
			NewPassword: "new_password",
		})
		require.NotNil(suite.T(), err)
		repoMock.AssertExpectations(suite.T())
	})
	suite.T().Run("failed compare password", func(t *testing.T) {
		repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
			Once().Return(suite.user, nil)
		err := accSvc.UpdateUserPassword(context.TODO(), &account.UpdatePasswordForm{
			ID:          1,
			Password:    "testfailed",
			NewPassword: "new_password",
		})
		require.NotNil(suite.T(), err)
		repoMock.AssertExpectations(suite.T())
	})
	suite.T().Run("failed update password", func(t *testing.T) {
		repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
			Once().Return(suite.user, nil)
		repoMock.On("UpdateUserByID", mock.Anything, mock.Anything).
			Once().Return(nil, errors.New("UNEXPECTED"))
		err := accSvc.UpdateUserPassword(context.TODO(), &account.UpdatePasswordForm{
			ID:          1,
			Password:    "secret",
			NewPassword: "new_password",
		})
		require.NotNil(suite.T(), err)
		repoMock.AssertExpectations(suite.T())
	})
}

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	repoMock.On("DeleteUserByID", mock.Anything, mock.Anything).
		Once().Return(nil)
	err := accSvc.RemoveUser(context.TODO(), suite.user)
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestService_DeleteUser_ShouldErrorWhenFindNotFound() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(nil, sql.ErrNoRows)
	err := accSvc.RemoveUser(context.TODO(), suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "user not found"})
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenFind() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	err := accSvc.RemoveUser(context.TODO(), suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenDelete() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	repoMock.On("DeleteUserByID", mock.Anything, mock.Anything).
		Once().Return(errors.New("UNEXPECTED"))
	err := accSvc.RemoveUser(context.TODO(), suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AuthenticateUser_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	user, err := accSvc.AuthenticateUser(context.TODO(), &account.LoginForm{
		Username: "test",
		Password: "secret",
	})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), user)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_AuthenticateUser_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)

	suite.T().Run("should return error no user", func(t *testing.T) {
		repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
			Once().Return(nil, nil)
		user, err := accSvc.AuthenticateUser(context.TODO(), &account.LoginForm{
			Username: "test", Password: "secret"})
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), user)
		repoMock.AssertExpectations(suite.T())
	})
	suite.T().Run("should return error password", func(t *testing.T) {
		repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
			Once().Return(suite.user, nil)
		user, err := accSvc.AuthenticateUser(context.TODO(), &account.LoginForm{
			Username: "test", Password: "secret_failed"})
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), user)
		repoMock.AssertExpectations(suite.T())
	})

	// need to inject the mocks -- so we skip this for now!
	// suite.T().Run("should return error generate token", func(t *testing.T) {})
}

func (suite *accountTestSuite) TestAccountService_RefreshToken_ShouldSuccess() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	claim := jwt.MapClaims{"id": suite.user.ID, "role_id": suite.user.Role.ID, "role_name": suite.user.Role.Name}
	newToken, tokenErr := utils.NewJWT(claim, config.Instance.JWTSecretKey, 30)
	require.Nil(suite.T(), tokenErr)
	repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
		Once().Return(suite.user, nil)
	user, err := accSvc.RefreshToken(context.TODO(), newToken)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), user)
	repoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_RefreshToken_ShouldError() {
	repoMock := new(mocks.MockIAccountRepository)
	accSvc := account.NewAccountService(repoMock)
	claim := jwt.MapClaims{"id": suite.user.ID, "role_id": suite.user.Role.ID, "role_name": suite.user.Role.Name}

	suite.T().Run("should err parse jwt", func(t *testing.T) {
		newToken, tokenErr := utils.NewJWT(claim, "wrong_key", 30)
		require.Nil(suite.T(), tokenErr)
		user, err := accSvc.RefreshToken(context.TODO(), newToken)
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), user)
	})

	suite.T().Run("should err claim id", func(t *testing.T) {
		wrongClaim := jwt.MapClaims{}
		newToken, tokenErr := utils.NewJWT(wrongClaim, config.Instance.JWTSecretKey, 30)
		require.Nil(suite.T(), tokenErr)
		user, err := accSvc.RefreshToken(context.TODO(), newToken)
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), user)
	})

	suite.T().Run("should err get user", func(t *testing.T) {
		newToken, tokenErr := utils.NewJWT(claim, config.Instance.JWTSecretKey, 30)
		require.Nil(suite.T(), tokenErr)
		repoMock.On("FindUserBy", mock.Anything, mock.Anything, mock.Anything).
			Once().Return(nil, errors.New("UNEXPECTED"))
		user, err := accSvc.RefreshToken(context.TODO(), newToken)
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), user)
	})

	// skip for now because we need to inject the utils mock
	// suite.T().Run("should err generate token", func(t *testing.T) {})
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
