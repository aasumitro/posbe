package sql_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/domain"
	repoSql "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type addonRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.ICRUDRepository[domain.Addon]
}

func (suite *addonRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.repo = repoSql.NewAddonSQLRepository(db)
}

func (suite *addonRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *addonRepositoryTestSuite) TestRepository_All_ExpectReturnRows() {
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, "test", "test", 1).
		AddRow(2, "test 2", "test 2", 1)
	query := "SELECT * FROM addons"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *addonRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromQuery() {
	query := "SELECT * FROM addons"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnError(errors.New(""))
	res, err := suite.repo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *addonRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromScan() {
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, "test", "test", 1).
		AddRow(nil, nil, nil, nil)
	query := "SELECT * FROM addons"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *addonRepositoryTestSuite) TestRepository_Find_ExpectReturnRow() {
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, "test", "test", 1)
	query := "SELECT * FROM addons WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *addonRepositoryTestSuite) TestRepository_Find_ExpectReturnError() {
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(nil, nil, nil, nil)
	query := "SELECT * FROM addons WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *addonRepositoryTestSuite) TestRepository_Created_ExpectSuccess() {
	addon := &domain.Addon{ID: 1, Name: "test", Description: "test", Price: 1}
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, "test", "test", 1)
	query := "INSERT INTO addons (name, description, price) VALUES ($1, $2, $3) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(addon.Name, addon.Description, addon.Price).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), addon)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *addonRepositoryTestSuite) TestRepository_Created_ExpectError() {
	addon := &domain.Addon{ID: 1, Name: "test", Description: "test", Price: 1}
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, nil, nil, nil)
	query := "INSERT INTO addons (name, description, price) VALUES ($1, $2, $3) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(addon.Name, addon.Description, addon.Price).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), addon)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *addonRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	addon := &domain.Addon{ID: 1, Name: "test", Description: "test", Price: 1}
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, "test", "test", 1)
	query := "UPDATE addons SET name = $1, description = $2, price = $3 WHERE id = $4 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(addon.Name, addon.Description, addon.Price, addon.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), addon)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *addonRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	addon := &domain.Addon{ID: 1, Name: "test", Description: "test", Price: 1}
	data := suite.mock.
		NewRows([]string{"id", "name", "description", "price"}).
		AddRow(1, nil, nil, nil)
	query := "UPDATE addons SET name = $1, description = $2, price = $3 WHERE id = $4 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(addon.Name, addon.Description, addon.Price, addon.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), addon)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *addonRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM addons WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	data := &domain.Addon{ID: 1}
	err := suite.repo.Delete(context.TODO(), data)
	require.Nil(suite.T(), err)
}

func TestAddonRepository(t *testing.T) {
	suite.Run(t, new(addonRepositoryTestSuite))
}
