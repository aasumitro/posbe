package sql_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/config"
	repoSql "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type productVariantsRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo model.ICRUDRepository[model.ProductVariant]
}

func (suite *productVariantsRepositoryTestSuite) SetupSuite() {
	var err error

	config.PostgresPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.repo = repoSql.NewProductVariantSQLRepository()
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Find_ExpectReturnRow() {
	data := suite.mock.
		NewRows([]string{"id", "product_id", "unit_id", "unit_size", "type", "name", "description", "price"}).
		AddRow(1, 1, 1, 12, "color", "test", "test", 12)
	query := "SELECT * FROM product_variants WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), model.FindWithID, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Find_ExpectReturnError() {
	data := suite.mock.
		NewRows([]string{"id", "product_id", "unit_id", "unit_size", "type", "name", "description", "price"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil)
	query := "SELECT * FROM product_variants WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), model.FindWithID, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Created_ExpectSuccess() {
	variant := &model.ProductVariant{ID: 1, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test", Description: sql.NullString{String: "test"}, Price: 12}
	data := suite.mock.
		NewRows([]string{"id", "product_id", "unit_id", "unit_size", "type", "name", "description", "price"}).
		AddRow(1, 1, 1, 12, "color", "test", "test", 12)
	query := "INSERT INTO product_variants (product_id, unit_id, unit_size, type, name, description, price) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(variant.ProductID, variant.UnitID, variant.UnitSize, variant.Type, variant.Name, variant.Description, variant.Price).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), variant)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Created_ExpectError() {
	variant := &model.ProductVariant{ID: 1, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test", Description: sql.NullString{String: "test"}, Price: 12}
	data := suite.mock.
		NewRows([]string{"id", "product_id", "unit_id", "unit_size", "type", "name", "description", "price"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil)
	query := "INSERT INTO product_variants (product_id, unit_id, unit_size, type, name, description, price) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(variant.ProductID, variant.UnitID, variant.UnitSize, variant.Type, variant.Name, variant.Description, variant.Price).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), variant)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	variant := &model.ProductVariant{ID: 1, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test", Description: sql.NullString{String: "test"}, Price: 12}
	data := suite.mock.
		NewRows([]string{"id", "product_id", "unit_id", "unit_size", "type", "name", "description", "price"}).
		AddRow(1, 1, 1, 12, "color", "test", "test", 12)
	query := "UPDATE product_variants SET product_id = $1, unit_id = $2, unit_size = $3, type = $4, name = $5, description = $6, price = $7 WHERE id = $8 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(variant.ProductID, variant.UnitID, variant.UnitSize, variant.Type, variant.Name, variant.Description, variant.Price, variant.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), variant)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	variant := &model.ProductVariant{ID: 1, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test", Description: sql.NullString{String: "test"}, Price: 12}
	data := suite.mock.
		NewRows([]string{"id", "product_id", "unit_id", "unit_size", "type", "name", "description", "price"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil)
	query := "UPDATE product_variants SET product_id = $1, unit_id = $2, unit_size = $3, type = $4, name = $5, description = $6, price = $7 WHERE id = $8 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(variant.ProductID, variant.UnitID, variant.UnitSize, variant.Type, variant.Name, variant.Description, variant.Price, variant.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), variant)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productVariantsRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM product_variants WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	data := &model.ProductVariant{ID: 1}
	err := suite.repo.Delete(context.TODO(), data)
	require.Nil(suite.T(), err)
}

func TestProductVariantsRepository(t *testing.T) {
	suite.Run(t, new(productVariantsRepositoryTestSuite))
}
