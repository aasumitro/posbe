package sql_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/domain"
	repoSql "github.com/aasumitro/posbe/internal/store/repository/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

type floorRepositoryTestSuite struct {
	suite.Suite
	mock      sqlmock.Sqlmock
	floorRepo domain.ICRUDRepository[domain.Floor]
}

func (suite *floorRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.floorRepo = repoSql.NewFloorSQLRepository(db)
}

func (suite *floorRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_All_ExpectedReturnDataRows() {
	floors := suite.mock.
		NewRows([]string{"id", "name", "total_tables", "total_rooms", "created_at", "updated_at"}).
		AddRow(1, "test", 1, 1, "13123", "123123").
		AddRow(2, "test 2", 1, 2, "13123", "123123")
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at "
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "GROUP BY floors.id ORDER BY floors.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(floors)
	res, err := suite.floorRepo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_All_ExpectedReturnErrorFromQuery() {
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at "
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "GROUP BY floors.id ORDER BY floors.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.floorRepo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_All_ExpectedReturnErrorFromScan() {
	floors := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test").
		AddRow(nil, nil)
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at "
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "GROUP BY floors.id ORDER BY floors.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(floors)
	res, err := suite.floorRepo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Find_ExpectedSuccess() {
	floor := suite.mock.
		NewRows([]string{"id", "name", "total_tables", "total_rooms", "created_at", "updated_at"}).
		AddRow(1, "test", 1, 1, "13123", "123123")
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at"
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "WHERE floors.id = $1 GROUP BY floors.id LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(floor)
	res, err := suite.floorRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Find_ExpectedError() {
	floor := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(nil, nil)
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at"
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "WHERE floors.id = $1 GROUP BY floors.id LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(floor)
	res, err := suite.floorRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Create_ExpectedSuccess() {
	floor := &domain.Floor{ID: 1, Name: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, "test", "123123", "12312312")
	expectedQuery := regexp.QuoteMeta("INSERT INTO floors (name, created_at) values ($1, $2) RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(floor.Name, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.floorRepo.Create(context.TODO(), floor)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Create_ExpectedError() {
	floor := &domain.Floor{ID: 1, Name: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, nil, nil, nil)
	expectedQuery := regexp.QuoteMeta("INSERT INTO floors (name, created_at) values ($1, $2) RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(floor.Name, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.floorRepo.Create(context.TODO(), floor)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Update_ExpectedSuccess() {
	floor := &domain.Floor{ID: 1, Name: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, "test", "123123", "12312312")
	expectedQuery := regexp.QuoteMeta("UPDATE floors SET name = $1, updated_at = $2 WHERE id = $3 RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(floor.Name, time.Now().Unix(), floor.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.floorRepo.Update(context.TODO(), floor)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Update_ExpectedError() {
	floor := &domain.Floor{ID: 1, Name: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, nil, nil, nil)
	expectedQuery := regexp.QuoteMeta("UPDATE floors SET name = $1, updated_at = $2 WHERE id = $3 RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(floor.Name, time.Now().Unix(), floor.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.floorRepo.Update(context.TODO(), floor)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *floorRepositoryTestSuite) TestFloorRepository_Delete_ExpectedSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM floors")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	floor := &domain.Floor{ID: 1, Name: "test"}
	err := suite.floorRepo.Delete(context.TODO(), floor)
	require.Nil(suite.T(), err)
}

func TestFloorRepository(t *testing.T) {
	suite.Run(t, new(floorRepositoryTestSuite))
}
