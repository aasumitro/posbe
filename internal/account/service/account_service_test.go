package service_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/domain/mocks"
	"github.com/aasumitro/posbe/internal/account/service"
	svcErr "github.com/aasumitro/posbe/pkg/errors"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type accountTestSuite struct {
	suite.Suite
	Db     *sql.DB
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
		RoleId:   1,
		Name:     "lorem ipsum",
		Username: "lorem",
		Email:    "lorem@ipsum.id",
		Phone:    "+628227111111",
		Role:     *suite.role,
	}

	suite.users = []*domain.User{
		suite.user,
		{
			ID:       2,
			RoleId:   1,
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
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldSuccess() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
	accSvc := service.NewAccountService(context.TODO(),
		roleRepoMock, userRepoMock)
	roleRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.roles, nil)

	data, err := accSvc.RoleList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.roles)
	roleRepoMock.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestAccountService_RoleList_ShouldError() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

func (suite *accountTestSuite) TestAccountService_DeleteRole_ShouldErrorInternal() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

// TODO: Create Password Coverage
func (suite *accountTestSuite) TestAccountService_AddUser_ShouldSuccess() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

func (suite *accountTestSuite) TestAccountService_AddUser_ShouldError() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

func (suite *accountTestSuite) TestAccountService_EditUser_ShouldError() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

// END TODO

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldSuccess() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

func (suite *accountTestSuite) TestAccountService_DeleteUser_ShouldErrorWhenFind() {
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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
	roleRepoMock := new(mocks.ICRUDRepository[domain.Role])
	userRepoMock := new(mocks.ICRUDRepository[domain.User])
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

func (suite *accountTestSuite) TestAccountService_T_K() {
	// TODO
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}