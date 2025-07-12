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
	mocks2 "github.com/aasumitro/posbe/mocks"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
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
		Password: "2ad1a22d5b3c9396d16243d2fe7f067976363715e322203a456278bb80b0b4a4.7ab4dcccfcd9d36efc68f1626d2fb80804a6508f9c3a7b44f430ba082b6870d2",
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
			Password: "secret",
		},
	}

	suite.svcErr = &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	}

	config.RedisPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess_ReturnModel() {
	config.RedisPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
	cacheMock := new(mocks2.Cache)
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("All", mock.Anything).
		Return(suite.roles, nil).Once()
	cacheMock.On("CacheFirstData", mock.Anything).
		Return(suite.roles, nil).Once()
	data, err := accSvc.RoleList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess_ReturnString() {
	config.RedisPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
	cacheMock := new(mocks2.Cache)
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("All", mock.Anything).
		Return(nil, nil).Once()
	jsonData, _ := json.Marshal(suite.roles)
	config.RedisPool.Set(context.TODO(), "roles", jsonData, 1)
	cacheMock.On("CacheFirstData", &utils.CacheDataSupplied{
		Key: "roles",
		TTL: time.Hour * 1,
		CbF: nil,
	}).Return(jsonData, nil).Once()
	data, err := accSvc.RoleList(context.TODO())
	suite.T().Log(data)
	suite.T().Log(err)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.RoleList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_UserList_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.users, nil)
	data, err := accSvc.UserList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_UserList_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.UserList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_ShowUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	data, err := accSvc.ShowUser(context.TODO(), suite.user.ID)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.user)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_ShowUser_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.ShowUser(context.TODO(), suite.user.ID)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.users[1], nil)
	data, err := accSvc.AddUser(context.TODO(), suite.users[1])
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users[1])
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError_Password() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	pwdMock := new(mocks2.IPassword)
	pwdMock.
		On("HashPassword").
		Return("", errors.New("UNEXPECTED")).
		Once()
	accSvc := account.NewAccountServiceTest(
		roleRepoMock, userRepoMock, pwdMock)
	data, err := accSvc.AddUser(context.TODO(), suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	userRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.AddUser(context.TODO(), suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.users[1], nil)
	data, err := accSvc.EditUser(context.TODO(), suite.users[1])
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users[1])
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError_Password() {
	// TODO: fix this
	suite.T().Skip()
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	pwdMock := new(mocks2.IPassword)
	pwdMock.
		On("HashPassword").
		Return("", errors.New("UNEXPECTED")).
		Once()
	accSvc := account.NewAccountServiceTest(
		roleRepoMock, userRepoMock, pwdMock)
	data, err := accSvc.EditUser(context.TODO(), suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	userRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.EditUser(context.TODO(), suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	userRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := accSvc.DeleteUser(context.TODO(), suite.user)
	require.Nil(suite.T(), err)
	roleRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestService_DeleteUser_ShouldErrorWhenFindNotFound() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	svc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteUser(context.TODO(), suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	userRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenFind() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := accSvc.DeleteUser(context.TODO(), suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenDelete() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	userRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := accSvc.DeleteUser(context.TODO(), suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	user, err := accSvc.VerifyUserCredentials(context.TODO(),
		suite.user.Username, "secret")
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorFind() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	user, err := accSvc.VerifyUserCredentials(context.TODO(),
		suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorWhenFindNotFound() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	svc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	user, err := svc.VerifyUserCredentials(context.TODO(),
		suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	userRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorComparePassword() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	pwdUtil := new(mocks2.IPassword)
	accSvc := account.NewAccountServiceTest(
		roleRepoMock, userRepoMock, pwdUtil)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	pwdUtil.
		On("ComparePasswords").
		Return(false, errors.New("LOREM")).
		Once()
	user, err := accSvc.VerifyUserCredentials(context.TODO(),
		suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorPassword() {
	roleRepoMock := new(mocks2.ICRUDRepository[model.Role])
	userRepoMock := new(mocks2.ICRUDRepository[model.User])
	accSvc := account.NewAccountService(
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	user, err := accSvc.VerifyUserCredentials(context.TODO(),
		suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
