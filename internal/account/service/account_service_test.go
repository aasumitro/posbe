package service_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	svcErr "github.com/aasumitro/posbe/commons"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/internal/account/service"
	mocks2 "github.com/aasumitro/posbe/mocks"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type accountTestSuite struct {
	suite.Suite
	role   *domain.Role
	roles  []*domain.Role
	user   *domain.User
	users  []*domain.User
	svcErr *utils.ServiceError
}

func (suite *accountTestSuite) SetupSuite() {
	suite.role = &domain.Role{
		ID:          1,
		Name:        "lorem",
		Description: "lorem ipsum",
		Usage:       1,
	}

	suite.roles = []*domain.Role{
		suite.role,
		{
			ID:          2,
			Name:        "dolor",
			Description: "Dolor Sit Amet",
		},
	}

	suite.user = &domain.User{
		ID:       1,
		RoleID:   1,
		Name:     "lorem ipsum",
		Username: "lorem",
		Email:    "lorem@ipsum.id",
		Phone:    "+628227111111",
		Role:     *suite.role,
		Password: "2ad1a22d5b3c9396d16243d2fe7f067976363715e322203a456278bb80b0b4a4.7ab4dcccfcd9d36efc68f1626d2fb80804a6508f9c3a7b44f430ba082b6870d2",
	}

	suite.users = []*domain.User{
		suite.user,
		{
			ID:       2,
			RoleID:   1,
			Name:     "dolor amet",
			Username: "dolor",
			Email:    "dolor@amet.id",
			Phone:    "+628227222222",
			Role:     *suite.role,
			Password: "secret",
		},
	}

	suite.svcErr = &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	}

	configs.RedisPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess_ReturnModel() {
	configs.RedisPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
	cacheMock := new(mocks2.Cache)
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)

	roleRepoMock.
		On("All", mock.Anything).
		Return(suite.roles, nil).Once()
	cacheMock.On("CacheFirstData", mock.Anything).
		Return(suite.roles, nil).Once()

	data, err := accSvc.RoleList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess_ReturnString() {
	configs.RedisPool = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
	cacheMock := new(mocks2.Cache)
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("All", mock.Anything).
		Return(nil, nil).Once()
	jsonData, _ := json.Marshal(suite.roles)
	configs.RedisPool.Set(context.TODO(), "roles", jsonData, 1)
	cacheMock.On("CacheFirstData", &utils.CacheDataSupplied{
		Key: "roles",
		TTL: time.Hour * 1,
		CbF: nil,
	}).Return(jsonData, nil).Once()
	data, err := accSvc.RoleList()
	suite.T().Log(data)
	suite.T().Log(err)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.RoleList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddRole_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.role, nil)
	data, err := accSvc.AddRole(suite.role)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.role)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddRole_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.AddRole(suite.role)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditRole_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.role, nil)
	data, err := accSvc.EditRole(suite.role)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.role)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditRole_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.EditRole(suite.role)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteRole_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.roles[1], nil)
	roleRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := accSvc.DeleteRole(suite.roles[1])
	require.Nil(suite.T(), err)
	roleRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestService_DeleteRole_ShouldErrorWhenFindNotFound() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	svc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteRole(suite.role)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	roleRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_DeleteRole_ShouldErrorInternal() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := accSvc.DeleteRole(suite.role)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteRole_ShouldErrorUsage() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.role, nil)
	err := accSvc.DeleteRole(suite.role)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{
		Code:    http.StatusForbidden,
		Message: svcErr.ErrorUnableToDelete,
	})
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteRole_ShouldErrorWhenDelete() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.roles[1], nil)
	roleRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := accSvc.DeleteRole(suite.roles[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_UserList_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.users, nil)
	data, err := accSvc.UserList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_UserList_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.UserList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_ShowUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	data, err := accSvc.ShowUser(suite.user.ID)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.user)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_ShowUser_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.ShowUser(suite.user.ID)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.users[1], nil)
	data, err := accSvc.AddUser(suite.users[1])
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users[1])
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError_Password() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	pwdMock := new(mocks2.IPassword)
	pwdMock.
		On("HashPassword").
		Return("", errors.New("UNEXPECTED")).
		Once()
	accSvc := service.NewAccountServiceTest(context.TODO(),
		roleRepoMock, userRepoMock, pwdMock)
	data, err := accSvc.AddUser(suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	userRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.AddUser(suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.users[1], nil)
	data, err := accSvc.EditUser(suite.users[1])
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.users[1])
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError_Password() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	pwdMock := new(mocks2.IPassword)
	pwdMock.
		On("HashPassword").
		Return("", errors.New("UNEXPECTED")).
		Once()
	accSvc := service.NewAccountServiceTest(context.TODO(),
		roleRepoMock, userRepoMock, pwdMock)
	data, err := accSvc.EditUser(suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	userRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := accSvc.EditUser(suite.users[1])
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	userRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := accSvc.DeleteUser(suite.user)
	require.Nil(suite.T(), err)
	roleRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestService_DeleteUser_ShouldErrorWhenFindNotFound() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	svc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteUser(suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	userRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenFind() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := accSvc.DeleteUser(suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenDelete() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	userRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := accSvc.DeleteUser(suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldSuccess() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	user, err := accSvc.VerifyUserCredentials(suite.user.Username, "secret")
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorFind() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	user, err := accSvc.VerifyUserCredentials(suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorWhenFindNotFound() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	svc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	user, err := svc.VerifyUserCredentials(suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	userRepoMock.AssertExpectations(suite.T())
}
func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorComparePassword() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	pwdUtil := new(mocks2.IPassword)
	accSvc := service.NewAccountServiceTest(context.TODO(),
		roleRepoMock, userRepoMock, pwdUtil)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	pwdUtil.
		On("ComparePasswords").
		Return(false, errors.New("LOREM")).
		Once()
	user, err := accSvc.VerifyUserCredentials(suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_VerifyUserCredentials_ShouldErrorPassword() {
	roleRepoMock := new(mocks2.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks2.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	userRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.user, nil)
	user, err := accSvc.VerifyUserCredentials(suite.user.Username, suite.user.Password)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), user)
	roleRepoMock.AssertExpectations(suite.T())
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
