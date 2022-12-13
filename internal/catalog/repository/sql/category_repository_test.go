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

type categoryRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.ICRUDRepository[domain.Category]
}

func (suite *categoryRepositoryTestSuite) SetupSuite() {
	var (
		err error
	)

	config.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.repo = repoSql.NewCategorySQLRepository()
}

func (suite *categoryRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *categoryRepositoryTestSuite) TestRepository_All_ExpectReturnRows() {
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test").
		AddRow(2, "test 2")
	query := "SELECT * FROM categories"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *categoryRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromQuery() {
	query := "SELECT * FROM categories"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnError(errors.New(""))
	res, err := suite.repo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *categoryRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromScan() {
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test").
		AddRow(nil, nil)
	query := "SELECT * FROM categories"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Find_ExpectReturnRow() {
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test")
	query := "SELECT * FROM categories WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Find_ExpectReturnError() {
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(nil, nil)
	query := "SELECT * FROM categories WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Created_ExpectSuccess() {
	category := &domain.Category{ID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test")
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(category.Name).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), category)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Created_ExpectError() {
	category := &domain.Category{ID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, nil)
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(category.Name).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), category)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	category := &domain.Category{ID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, "test")
	query := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(category.Name, category.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), category)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	category := &domain.Category{ID: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "name"}).
		AddRow(1, nil)
	query := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(category.Name, category.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), category)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *categoryRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM categories WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	data := &domain.Category{ID: 1}
	err := suite.repo.Delete(context.TODO(), data)
	require.Nil(suite.T(), err)
}

func TestCategoryRepository(t *testing.T) {
	suite.Run(t, new(categoryRepositoryTestSuite))
}
