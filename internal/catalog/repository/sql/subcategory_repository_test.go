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

type subcategoryRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.ICRUDRepository[domain.Subcategory]
}

func (suite *subcategoryRepositoryTestSuite) SetupSuite() {
	var (
		err error
	)

	config.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.repo = repoSql.NewSubcategorySQLRepository()
}

func (suite *subcategoryRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_All_ExpectReturnRows() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, 1, "test").
		AddRow(2, 1, "test 2")
	query := "SELECT * FROM subcategories"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromQuery() {
	query := "SELECT * FROM subcategories"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnError(errors.New(""))
	res, err := suite.repo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromScan() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, 1, "test").
		AddRow(nil, nil, nil)
	query := "SELECT * FROM subcategories"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Find_ExpectReturnRow() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, 1, "test")
	query := "SELECT * FROM subcategories WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Find_ExpectReturnError() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(nil, nil, nil)
	query := "SELECT * FROM subcategories WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Created_ExpectSuccess() {
	subcategory := &domain.Subcategory{ID: 1, CategoryId: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, 1, "test")
	query := "INSERT INTO subcategories (category_id, name) VALUES ($1, $2) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(subcategory.CategoryId, subcategory.Name).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), subcategory)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Created_ExpectError() {
	subcategory := &domain.Subcategory{ID: 1, CategoryId: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, nil, nil)
	query := "INSERT INTO subcategories (category_id, name) VALUES ($1, $2) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(subcategory.CategoryId, subcategory.Name).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), subcategory)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	subcategory := &domain.Subcategory{ID: 1, CategoryId: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, 1, "test")
	query := "UPDATE subcategories SET category_id = $1, name = $2 WHERE id = $3 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(subcategory.CategoryId, subcategory.Name, subcategory.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), subcategory)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	subcategory := &domain.Subcategory{ID: 1, CategoryId: 1, Name: "test"}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "name"}).
		AddRow(1, nil, nil)
	query := "UPDATE subcategories SET category_id = $1, name = $2 WHERE id = $3 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(subcategory.CategoryId, subcategory.Name, subcategory.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), subcategory)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *subcategoryRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM subcategories WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	data := &domain.Subcategory{ID: 1}
	err := suite.repo.Delete(context.TODO(), data)
	require.Nil(suite.T(), err)
}

func TestSubcategoryRepository(t *testing.T) {
	suite.Run(t, new(subcategoryRepositoryTestSuite))
}
