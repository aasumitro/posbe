package sql_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/domain"
	repoSql "github.com/aasumitro/posbe/internal/store/repository/sql"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type roomRepositoryTestSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	roomRepo domain.ICRUDAddOnRepository[domain.Room]
}

func (suite *roomRepositoryTestSuite) SetupSuite() {
	var (
		err error
	)

	config.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.roomRepo = repoSql.NewRoomSQLRepository()
}

func (suite *roomRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_AllWhere_ExpectedReturnDataRows() {
	rooms := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123").
		AddRow(2, 1, "test 2", 1, 2, 3, 4, 5, 1, "13123", "123123")
	q := "SELECT * FROM rooms WHERE floor_id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rooms)
	res, err := suite.roomRepo.AllWhere(context.TODO(), domain.FindWithRelationId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_AllWhere_ExpectedReturnErrorFromQuery() {
	q := "SELECT * FROM rooms WHERE floor_id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.roomRepo.AllWhere(context.TODO(), domain.FindWithRelationId, 1)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_AllWhere_ExpectedReturnErrorFromScan() {
	rooms := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123").
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM rooms WHERE floor_id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rooms)
	res, err := suite.roomRepo.AllWhere(context.TODO(), domain.FindWithRelationId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_All_ExpectedReturnDataRows() {
	rooms := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123").
		AddRow(2, 1, "test 2", 1, 2, 3, 4, 5, 1, "13123", "123123")
	q := "SELECT * FROM rooms"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rooms)
	res, err := suite.roomRepo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_All_ExpectedReturnErrorFromQuery() {
	q := "SELECT * FROM rooms"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.roomRepo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_All_ExpectedReturnErrorFromScan() {
	rooms := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123").
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM rooms"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rooms)
	res, err := suite.roomRepo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Find_ExpectedSuccess() {
	rooms := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123")
	q := "SELECT * FROM rooms WHERE id = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rooms)
	res, err := suite.roomRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_FindRelation_ExpectedSuccess() {
	room := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123")
	q := "SELECT * FROM rooms WHERE floor_id = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(room)
	res, err := suite.roomRepo.Find(context.TODO(), domain.FindWithRelationId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Find_ExpectedError() {
	table := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM rooms WHERE id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(table)
	res, err := suite.roomRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Create_ExpectedSuccess() {
	room := &domain.Room{FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Price: 1}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123")
	q := "INSERT INTO rooms (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, price, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(room.FloorId, room.Name, room.XPos, room.YPos, room.WSize, room.HSize, room.Capacity, room.Price, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roomRepo.Create(context.TODO(), room)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Create_ExpectedError() {
	room := &domain.Room{FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Price: 1}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "INSERT INTO rooms (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, price, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(room.FloorId, room.Name, room.XPos, room.YPos, room.WSize, room.HSize, room.Capacity, room.Price, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roomRepo.Create(context.TODO(), room)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Update_ExpectedSuccess() {
	room := &domain.Room{ID: 1, FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Price: 1}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, 1, "13123", "123123")
	q := "UPDATE rooms SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, price = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(room.FloorId, room.Name, room.XPos, room.YPos, room.WSize, room.HSize, room.Capacity, room.Price, time.Now().Unix(), room.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roomRepo.Update(context.TODO(), room)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Update_ExpectedError() {
	room := &domain.Room{ID: 1, FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Price: 1}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "price", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "UPDATE rooms SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, price = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(room.FloorId, room.Name, room.XPos, room.YPos, room.WSize, room.HSize, room.Capacity, room.Price, time.Now().Unix(), room.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roomRepo.Update(context.TODO(), room)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roomRepositoryTestSuite) TestRoomRepository_Delete_ExpectedSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM rooms WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	room := &domain.Room{ID: 1}
	err := suite.roomRepo.Delete(context.TODO(), room)
	require.Nil(suite.T(), err)
}

func TestRoomRepository(t *testing.T) {
	suite.Run(t, new(roomRepositoryTestSuite))
}
