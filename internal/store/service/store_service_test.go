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
	room   *domain.Room
	rooms  []*domain.Room
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
			ID:         2,
			Name:       "dolor",
			TotalRooms: 1,
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

	suite.room = &domain.Room{
		ID:   1,
		Name: "lorem ipsum",
	}

	suite.rooms = []*domain.Room{
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
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteFloor(suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	floorRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorWhenFindNotFound() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
	floorRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteFloor(suite.floor)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	floorRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteFloor_ShouldErrorHasTables() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
	svc := service.NewStoreService(context.TODO(), floorRepoMock, new(mocks.ICRUDAddOnRepository[domain.Table]),
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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

// ================== TABLE TEST CASE
func (suite *storeTestSuite) TestStoreService_TableList_ShouldSuccess() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteTable(suite.table)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenFindNotFound() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
	tableRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteTable(suite.table)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	tableRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteTable_ShouldErrorWhenDelete() {
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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

// =================== ROOM TEST CASE
func (suite *storeTestSuite) TestStoreService_RoomList_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("All", mock.Anything).
		Once().
		Return(suite.rooms, nil)
	data, err := svc.RoomList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.rooms)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_RoomList_ShouldError() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.RoomList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddRoom_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(suite.room, nil)
	data, err := svc.AddRoom(suite.room)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.room)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_AddRoom_ShouldError() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Create", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AddRoom(suite.room)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditRoom_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(suite.room, nil)
	data, err := svc.EditRoom(suite.room)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.room)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_EditRoom_ShouldError() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Update", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.EditRoom(suite.room)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldSuccess() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.rooms[1], nil)
	roomRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(nil)
	err := svc.DeleteRoom(suite.rooms[1])
	require.Nil(suite.T(), err)
	roomRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldErrorWhenFind() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteRoom(suite.room)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldErrorWhenFindNotFound() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteRoom(suite.room)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	roomRepoMock.AssertExpectations(suite.T())
}
func (suite *storeTestSuite) TestStoreService_DeleteRoom_ShouldErrorWhenDelete() {
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Floor]),
		new(mocks.ICRUDAddOnRepository[domain.Table]),
		roomRepoMock)
	roomRepoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(suite.rooms[1], nil)
	roomRepoMock.
		On("Delete", mock.Anything, mock.Anything).
		Once().
		Return(errors.New("UNEXPECTED"))
	err := svc.DeleteRoom(suite.rooms[1])
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	roomRepoMock.AssertExpectations(suite.T())
}

// =================== ADD ON TEST CASE
func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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

	data, err := svc.FloorsWith(domain.Table{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*domain.Floor{suite.floors[0]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithTable_ShouldErrorAllWhere() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
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

	data, err := svc.FloorsWith(domain.Table{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*domain.Floor{suite.floors[0]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithRoom_ShouldSuccess() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Table]),
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

	data, err := svc.FloorsWith(domain.Room{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*domain.Floor{suite.floors[1]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWithRoom_ShouldErrorAllWhere() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	roomRepoMock := new(mocks.ICRUDAddOnRepository[domain.Room])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Table]),
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

	data, err := svc.FloorsWith(domain.Room{})
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, []*domain.Floor{suite.floors[1]})
	floorRepoMock.AssertExpectations(suite.T())
}

func (suite *storeTestSuite) TestStoreService_FloorsWith_ShouldErrorFind() {
	floorRepoMock := new(mocks.ICRUDRepository[domain.Floor])
	tableRepoMock := new(mocks.ICRUDAddOnRepository[domain.Table])
	svc := service.NewStoreService(context.TODO(),
		floorRepoMock, tableRepoMock,
		new(mocks.ICRUDAddOnRepository[domain.Room]))
	floorRepoMock.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.FloorsWith(domain.Table{})
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	tableRepoMock.AssertExpectations(suite.T())
}

func TestStoreService(t *testing.T) {
	suite.Run(t, new(storeTestSuite))
}
