package sql_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/domain"
	repoSql "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type unitRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.ICRUDRepository[domain.Unit]
}

func (suite *unitRepositoryTestSuite) SetupSuite() {
	var (
		err error
	)

	config.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.repo = repoSql.NewUnitSQLRepository()
}

func (suite *unitRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *unitRepositoryTestSuite) TestRepository_All_ExpectReturnRows() {
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, "test", "test", "test").
		AddRow(2, "test 2", "test 2", "test 2")
	query := "SELECT * FROM units"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *unitRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromQuery() {
	query := "SELECT * FROM units"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnError(errors.New(""))
	res, err := suite.repo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *unitRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromScan() {
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, "test", "test", "test").
		AddRow(nil, nil, nil, nil)
	query := "SELECT * FROM units"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *unitRepositoryTestSuite) TestRepository_Find_ExpectReturnRow() {
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, "test", "test", "test")
	query := "SELECT * FROM units WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *unitRepositoryTestSuite) TestRepository_Find_ExpectReturnError() {
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(nil, nil, nil, nil)
	query := "SELECT * FROM units WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *unitRepositoryTestSuite) TestRepository_Created_ExpectSuccess() {
	unit := &domain.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, "test", "test", "test")
	query := "INSERT INTO units (magnitude, name, symbol) VALUES ($1, $2, $3) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(unit.Magnitude, unit.Name, unit.Symbol).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), unit)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *unitRepositoryTestSuite) TestRepository_Created_ExpectError() {
	unit := &domain.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, nil, nil, nil)
	query := "INSERT INTO units (magnitude, name, symbol) VALUES ($1, $2, $3) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(unit.Magnitude, unit.Name, unit.Symbol).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), unit)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *unitRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	unit := &domain.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, "test", "test", "test")
	query := "UPDATE units SET magnitude = $1, name = $2, symbol = $3 WHERE id = $4 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(unit.Magnitude, unit.Name, unit.Symbol, unit.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), unit)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *unitRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	unit := &domain.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	data := suite.mock.
		NewRows([]string{"id", "magnitude", "name", "symbol"}).
		AddRow(1, nil, nil, nil)
	query := "UPDATE units SET magnitude = $1, name = $2, symbol = $3 WHERE id = $4 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(unit.Magnitude, unit.Name, unit.Symbol, unit.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), unit)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *unitRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM units WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	data := &domain.Unit{ID: 1}
	err := suite.repo.Delete(context.TODO(), data)
	require.Nil(suite.T(), err)
}

func TestUnitRepository(t *testing.T) {
	suite.Run(t, new(unitRepositoryTestSuite))
}
