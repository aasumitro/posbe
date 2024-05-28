package service_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	svcErr "github.com/aasumitro/posbe/common"
	"github.com/aasumitro/posbe/internal/store/service"
	"github.com/aasumitro/posbe/mocks"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storeTestSuite struct {
	suite.Suite
	Db     *sql.DB
	floor  *model.Floor
	floors []*model.Floor
	table  *model.Table
	tables []*model.Table
	room   *model.Room
	rooms  []*model.Room
	svcErr *utils.ServiceError
}

func (suite *storeTestSuite) SetupSuite() {
	suite.floor = &model.Floor{
		ID:          1,
		Name:        "lorem",
		TotalTables: 1,
	}

	suite.floors = []*model.Floor{
		suite.floor,
		{
			ID:         2,
			Name:       "dolor",
			TotalRooms: 1,
		},
	}

	suite.table = &model.Table{
		ID:   1,
		Name: "lorem ipsum",
	}

	suite.tables = []*model.Table{
		suite.table,
		{
			ID:   2,
			Name: "dolor amet",
		},
	}

	suite.room = &model.Room{
		ID:   1,
		Name: "lorem ipsum",
	}

	suite.rooms = []*model.Room{
		suite.room,
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

// ================= FLOOR TEST CASE
func (suite *storeTestSuite) TestStoreService_FloorList_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)
	data, err := svc.FloorList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floors)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorList_ShouldError() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.FloorList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddFloor_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.floor, nil)
	data, err := svc.AddFloor(context.TODO(), suite.floor)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floor)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddFloor_ShouldError() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AddFloor(context.TODO(), suite.floor)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditFloor_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.floor, nil)
	data, err := svc.EditFloor(context.TODO(), suite.floor)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.floor)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditFloor_ShouldError() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.EditFloor(context.TODO(), suite.floor)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.floors[1], nil)
	floorRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := svc.DeleteFloor(context.TODO(), suite.floors[1])
	require.Nil(suite.T(), err)
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorWhenFind() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteFloor(context.TODO(), suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorWhenFindNotFound() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteFloor(context.TODO(), suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	floorRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorHasTables() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.floor, nil)
	err := svc.DeleteFloor(context.TODO(), suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{
		Code:    http.StatusForbidden,
		Message: svcErr.ErrorUnableToDelete,
	})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorWhenDelete() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	svc := service.NewStoreService(floorRepoMock, new(mocks.ICRUDAddOnRepository[model.Table]),
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.floors[1], nil)
	floorRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := svc.DeleteFloor(context.TODO(), suite.floors[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}

// ================== TABLE TEST CASE
func (suite *storeTestSuite) TestStoreService_TableList_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.tables, nil)
	data, err := svc.TableList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.tables)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_TableList_ShouldError() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.TableList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddTable_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.table, nil)
	data, err := svc.AddTable(context.TODO(), suite.table)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.table)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddTable_ShouldError() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AddTable(context.TODO(), suite.table)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditTable_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.table, nil)
	data, err := svc.EditTable(context.TODO(), suite.table)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.table)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditTable_ShouldError() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.EditTable(context.TODO(), suite.table)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.tables[1], nil)
	tableRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := svc.DeleteTable(context.TODO(), suite.tables[1])
	require.Nil(suite.T(), err)
	tableRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenFind() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteTable(context.TODO(), suite.table)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenFindNotFound() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteTable(context.TODO(), suite.table)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	tableRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenDelete() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.tables[1], nil)
	tableRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := svc.DeleteTable(context.TODO(), suite.tables[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

// =================== ROOM TEST CASE
func (suite *storeTestSuite) TestStoreService_RoomList_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.rooms, nil)
	data, err := svc.RoomList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.rooms)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_RoomList_ShouldError() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.RoomList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddRoom_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.room, nil)
	data, err := svc.AddRoom(context.TODO(), suite.room)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.room)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddRoom_ShouldError() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AddRoom(context.TODO(), suite.room)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditRoom_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.room, nil)
	data, err := svc.EditRoom(context.TODO(), suite.room)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.room)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditRoom_ShouldError() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.EditRoom(context.TODO(), suite.room)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.rooms[1], nil)
	roomRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := svc.DeleteRoom(context.TODO(), suite.rooms[1])
	require.Nil(suite.T(), err)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldErrorWhenFind() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteRoom(context.TODO(), suite.room)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldErrorWhenFindNotFound() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteRoom(context.TODO(), suite.room)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	roomRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldErrorWhenDelete() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		new(mocks.ICRUDRepository[model.Floor]),
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.rooms[1], nil)
	roomRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := svc.DeleteRoom(context.TODO(), suite.rooms[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

// =================== ADD ON TEST CASE
func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		floorRepoMock, tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)

	for _, floor := range suite.floors {
		if floor.TotalTables >= 1 {
			require.NotNil(suite.T(), floor)
			tableRepoMock.
				On("AllWhere", mock.Anything, mock.Anything, mock.Anything).
				Once().
				Return(suite.tables, nil)
		}
	}

	data, err := svc.FloorsWith(context.TODO(), model.Table{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*model.Floor{suite.floors[0]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldErrorAllWhere() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		floorRepoMock, tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)

	for _, floor := range suite.floors {
		if floor.TotalTables >= 1 {
			require.NotNil(suite.T(), floor)
			tableRepoMock.
				On("AllWhere", mock.Anything, mock.Anything, mock.Anything).
				Once().
				Return(nil, errors.New("UNEXPECTED"))
		}
	}

	data, err := svc.FloorsWith(context.TODO(), model.Table{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*model.Floor{suite.floors[0]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithRoom_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		floorRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)

	for _, floor := range suite.floors {
		if floor.TotalRooms >= 1 {
			require.NotNil(suite.T(), floor)
			roomRepoMock.
				On("AllWhere", mock.Anything, mock.Anything, mock.Anything).
				Once().
				Return(suite.rooms, nil)
		}
	}

	data, err := svc.FloorsWith(context.TODO(), model.Room{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*model.Floor{suite.floors[1]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithRoom_ShouldErrorAllWhere() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	roomRepoMock := new(mocks.ICRUDAddOnRepository[model.Room])
	svc := service.NewStoreService(
		floorRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Table]),
		roomRepoMock)
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.floors, nil)

	for _, floor := range suite.floors {
		if floor.TotalRooms >= 1 {
			require.NotNil(suite.T(), floor)
			roomRepoMock.
				On("AllWhere", mock.Anything, mock.Anything, mock.Anything).
				Once().
				Return(nil, errors.New("UNEXPECTED"))
		}
	}

	data, err := svc.FloorsWith(context.TODO(), model.Room{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*model.Floor{suite.floors[1]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWith_ShouldErrorFind() {
	floorRepoMock := new(mocks.ICRUDRepository[model.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[model.Table])
	svc := service.NewStoreService(
		floorRepoMock, tableRepoMock,
		new(mocks.ICRUDAddOnRepository[model.Room]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.FloorsWith(context.TODO(), model.Table{})
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func TestStoreService(t *testing.T) {
	suite.Run(t, new(storeTestSuite))
}
