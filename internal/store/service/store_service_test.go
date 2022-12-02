package service_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/domain/mocks"
	"github.com/aasumitro/posbe/internal/store/service"
	svcErr "github.com/aasumitro/posbe/pkg/errors"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type storeTestSuite struct {
	suite.Suite
	Db     *sql.DB
	floor  *domain.Floor
	floors []*domain.Floor
	table  *domain.Table
	tables []*domain.Table
	svcErr *utils.ServiceError
}

func (suite *storeTestSuite) SetupSuite() {
	suite.floor = &domain.Floor{
		ID:          1,
		Name:        "lorem",
		TotalTables: 1,
	}

	suite.floors = []*domain.Floor{
		suite.floor,
		{
			ID:   2,
			Name: "dolor",
		},
	}

	suite.table = &domain.Table{
		ID:   1,
		Name: "lorem ipsum",
	}

	suite.tables = []*domain.Table{
		suite.table,
		{
			ID:   2,
			Name: "dolor amet",
		},
	}

	suite.svcErr = &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	}
}

func (suite *storeTestSuite) TestStoreService_FloorList_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)
	data, err := svc.FloorList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floors)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorList_ShouldError() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.FloorList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddFloor_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.floor, nil)
	data, err := svc.AddFloor(suite.floor)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floor)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddFloor_ShouldError() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AddFloor(suite.floor)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditFloor_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.floor, nil)
	data, err := svc.EditFloor(suite.floor)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floor)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditFloor_ShouldError() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.EditFloor(suite.floor)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.floors[1], nil)
	floorRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := svc.DeleteFloor(suite.floors[1])
	require.Nil(suite.T(), err)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorWhenFind() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteFloor(suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorHasTables() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.floor, nil)
	err := svc.DeleteFloor(suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{
		Code:    http.StatusForbidden,
		Message: svcErr.ErrorUnableToDelete,
	})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorWhenDelete() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.floors[1], nil)
	floorRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := svc.DeleteFloor(suite.floors[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_TableList_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.tables, nil)
	data, err := svc.TableList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.tables)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_TableList_ShouldError() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.TableList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddTable_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.table, nil)
	data, err := svc.AddTable(suite.table)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.table)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddTable_ShouldError() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AddTable(suite.table)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditTable_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.table, nil)
	data, err := svc.EditTable(suite.table)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.table)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditTable_ShouldError() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.EditTable(suite.table)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.tables[1], nil)
	tableRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := svc.DeleteTable(suite.tables[1])
	require.Nil(suite.T(), err)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenFind() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteTable(suite.table)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenDelete() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]), tableRepoMock)
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.tables[1], nil)
	tableRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := svc.DeleteTable(suite.tables[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, tableRepoMock)
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)

	for _, data := range suite.floors {
		require.NotNil(suite.T(), data)
		tableRepoMock.
			On("AllWhere", mock.Anything, mock.Anything, mock.Anything).
			Once().
			Return(suite.tables, nil)
	}

	data, err := svc.FloorsWithTables()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floors)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldErrorFind() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, tableRepoMock)
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.FloorsWithTables()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldErrorAllWhere() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, tableRepoMock)
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)

	for _, data := range suite.floors {
		require.NotNil(suite.T(), data)
		tableRepoMock.
			On("AllWhere", mock.Anything, mock.Anything, mock.Anything).
			Once().
			Return(nil, errors.New("UNEXPECTED"))
	}

	data, err := svc.FloorsWithTables()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floors)
	floorRepoMock.AssertExpectations(suite.T())
}

func TestStoreService(t *testing.T) {
	suite.Run(t, new(storeTestSuite))
}
