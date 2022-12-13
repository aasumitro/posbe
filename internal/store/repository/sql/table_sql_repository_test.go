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

type tableRepositoryTestSuite struct {
	suite.Suite
	mock      sqlmock.Sqlmock
	tableRepo domain.ICRUDAddOnRepository[domain.Table]
}

func (suite *tableRepositoryTestSuite) SetupSuite() {
	var (
		err error
	)

	config.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.tableRepo = repoSql.NewTableSQLRepository()
}

func (suite *tableRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *tableRepositoryTestSuite) TestTableRepository_AllWhere_ExpectedReturnDataRows() {
	tables := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123").
		AddRow(2, 1, "test 2", 1, 2, 3, 4, 5, "sq", "13123", "123123")
	q := "SELECT * FROM tables WHERE floor_id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(tables)
	res, err := suite.tableRepo.AllWhere(context.TODO(), domain.FindWithRelationId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_AllWhere_ExpectedReturnErrorFromQuery() {
	q := "SELECT * FROM tables WHERE floor_id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.tableRepo.AllWhere(context.TODO(), domain.FindWithRelationId, 1)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_AllWhere_ExpectedReturnErrorFromScan() {
	tables := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123").
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM tables WHERE floor_id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(tables)
	res, err := suite.tableRepo.AllWhere(context.TODO(), domain.FindWithRelationId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_All_ExpectedReturnDataRows() {
	tables := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123").
		AddRow(2, 1, "test 2", 1, 2, 3, 4, 5, "sq", "13123", "123123")
	q := "SELECT * FROM tables"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(tables)
	res, err := suite.tableRepo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_All_ExpectedReturnErrorFromQuery() {
	q := "SELECT * FROM tables"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.tableRepo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_All_ExpectedReturnErrorFromScan() {
	tables := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123").
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM tables"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(tables)
	res, err := suite.tableRepo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Find_ExpectedSuccess() {
	table := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123")
	q := "SELECT * FROM tables WHERE id = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(table)
	res, err := suite.tableRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_FindRelation_ExpectedSuccess() {
	table := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123")
	q := "SELECT * FROM tables WHERE floor_id = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(table)
	res, err := suite.tableRepo.Find(context.TODO(), domain.FindWithRelationId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Find_ExpectedError() {
	table := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM tables WHERE id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(table)
	res, err := suite.tableRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Create_ExpectedSuccess() {
	table := &domain.Table{FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Type: "sq"}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123")
	q := "INSERT INTO tables (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, type, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(table.FloorId, table.Name, table.XPos, table.YPos, table.WSize, table.HSize, table.Capacity, table.Type, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.tableRepo.Create(context.TODO(), table)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Create_ExpectedError() {
	table := &domain.Table{FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Type: "sq"}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "INSERT INTO tables (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, type, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(table.FloorId, table.Name, table.XPos, table.YPos, table.WSize, table.HSize, table.Capacity, table.Type, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.tableRepo.Create(context.TODO(), table)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Update_ExpectedSuccess() {
	table := &domain.Table{ID: 1, FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Type: "sq"}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(1, 1, "test", 1, 2, 3, 4, 5, "sq", "13123", "123123")
	q := "UPDATE tables SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, type = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(table.FloorId, table.Name, table.XPos, table.YPos, table.WSize, table.HSize, table.Capacity, table.Type, time.Now().Unix(), table.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.tableRepo.Update(context.TODO(), table)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Update_ExpectedError() {
	table := &domain.Table{ID: 1, FloorId: 1, Name: "test", XPos: 1, YPos: 1, WSize: 1, HSize: 1, Capacity: 1, Type: "sq"}
	rows := suite.mock.
		NewRows([]string{"id", "floor_id", "name", "x_pos", "y_pos", "w_size", "h_size", "capacity", "type", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "UPDATE tables SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, type = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(table.FloorId, table.Name, table.XPos, table.YPos, table.WSize, table.HSize, table.Capacity, table.Type, time.Now().Unix(), table.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.tableRepo.Update(context.TODO(), table)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *tableRepositoryTestSuite) TestTableRepository_Delete_ExpectedSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM table")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	table := &domain.Table{ID: 1}
	err := suite.tableRepo.Delete(context.TODO(), table)
	require.Nil(suite.T(), err)
}

func TestTableRepository(t *testing.T) {
	suite.Run(t, new(tableRepositoryTestSuite))
}
