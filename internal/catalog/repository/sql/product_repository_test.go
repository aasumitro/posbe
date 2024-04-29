package sql_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/config"
	repoSql "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type productRepositoryTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo model.ICRUDWithSearchRepository[model.Product]
}

func (suite *productRepositoryTestSuite) SetupSuite() {
	var err error

	config.PostgresPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.repo = repoSql.NewProductSQLRepository()
}

func (suite *productRepositoryTestSuite) TestRepository_Search_ExpectReturnRows() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, 1, 1, "12", "test", "test", "test", "test", 12)
	keys := []model.FindWith{model.FindWithCategoryID, model.FindWithSubcategoryID, model.FindWithSKU, model.FindWithPriceInRange}
	values := []any{1, 1, "12", []float32{10, 12}}
	query := "SELECT * FROM products WHERE category_id = 1 AND subcategory_id = 1 AND sku = '12' AND price BETWEEN 10.000000 AND 12.000000"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Search(context.TODO(), keys, values)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_Search_ExpectReturnErrorFromQuery() {
	keys := []model.FindWith{model.FindWithCategoryID, model.FindWithSubcategoryID, model.FindWithSKU, model.FindWithPriceInRange}
	values := []any{1, 1, "12", []float32{10, 12}}
	query := "SELECT * FROM products WHERE category_id = 1 AND subcategory_id = 1 AND sku = '12' AND price BETWEEN 10.000000 AND 12.000000"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnError(errors.New(""))
	res, err := suite.repo.Search(context.TODO(), keys, values)
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_Search_ExpectReturnErrorFromScan() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil, nil)
	keys := []model.FindWith{model.FindWithCategoryID, model.FindWithSubcategoryID, model.FindWithSKU, model.FindWithPriceInRange}
	values := []any{1, 1, "12", []float32{10, 12}}
	query := "SELECT * FROM products WHERE category_id = 1 AND subcategory_id = 1 AND sku = '12' AND price BETWEEN 10.000000 AND 12.000000"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Search(context.TODO(), keys, values)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productRepositoryTestSuite) TestRepository_All_ExpectReturnRows() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, 1, 1, "12", "test", "test", "test", "test", 12)
	query := "SELECT * FROM products"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromQuery() {
	query := "SELECT * FROM products"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnError(errors.New(""))
	res, err := suite.repo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_All_ExpectReturnErrorFromScan() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil, nil)
	query := "SELECT * FROM products"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productRepositoryTestSuite) TestRepository_Find_ExpectReturnRow() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, 1, 1, "12", "test", "test", "test", "test", 12)
	query := "SELECT * FROM products WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), model.FindWithID, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_Find_ExpectReturnError() {
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil, nil)
	query := "SELECT * FROM products WHERE id = $1 LIMIT 1"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).WillReturnRows(data)
	res, err := suite.repo.Find(context.TODO(), model.FindWithID, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productRepositoryTestSuite) TestRepository_Created_ExpectSuccess() {
	product := &model.Product{ID: 1, CategoryID: 1, SubcategoryID: 1, Sku: "12", Image: sql.NullString{String: "test"}, Gallery: sql.NullString{String: "test"}, Name: "test", Price: 12, Description: sql.NullString{String: "test"}}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, 1, 1, "12", "test", "test", "test", "test", 12)
	q := "INSERT INTO products "
	q += "(category_id, subcategory_id, sku, image, gallery, name, description, price) "
	q += "VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	meta := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(meta).
		WithArgs(product.CategoryID, product.SubcategoryID,
			product.Sku, product.Image, product.Gallery, product.Name,
			product.Description, product.Price).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), product)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_Created_ExpectError() {
	product := &model.Product{ID: 1, CategoryID: 1, SubcategoryID: 1, Sku: "12", Image: sql.NullString{String: "test"}, Gallery: sql.NullString{String: "test"}, Name: "test", Price: 12, Description: sql.NullString{String: "test"}}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "INSERT INTO products "
	q += "(category_id, subcategory_id, sku, image, gallery, name, description, price) "
	q += "VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	meta := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(meta).
		WithArgs(product.CategoryID, product.SubcategoryID,
			product.Sku, product.Image, product.Gallery, product.Name,
			product.Description, product.Price).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Create(context.TODO(), product)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productRepositoryTestSuite) TestRepository_Updated_ExpectSuccess() {
	product := &model.Product{ID: 1, CategoryID: 1, SubcategoryID: 1, Sku: "12", Image: sql.NullString{String: "test"}, Gallery: sql.NullString{String: "test"}, Name: "test", Price: 12, Description: sql.NullString{String: "test"}}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, 1, 1, "12", "test", "test", "test", "test", 12)
	query := "UPDATE products SET category_id = $1, subcategory_id = $2, sku = $3, image = $4, gallery = $5, name = $6, description = $7, price = $8 WHERE id = $9 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(product.CategoryID, product.SubcategoryID,
			product.Sku, product.Image, product.Gallery, product.Name,
			product.Description, product.Price, product.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), product)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *productRepositoryTestSuite) TestRepository_Updated_ExpectError() {
	product := &model.Product{ID: 1, CategoryID: 1, SubcategoryID: 1, Sku: "12", Image: sql.NullString{String: "test"}, Gallery: sql.NullString{String: "test"}, Name: "test", Price: 12, Description: sql.NullString{String: "test"}}
	data := suite.mock.
		NewRows([]string{"id", "category_id", "subcategory_id", "sku", "image", "gallery", "name", "price", "description"}).
		AddRow(1, nil, nil, nil, nil, nil, nil, nil, nil)
	query := "UPDATE products SET category_id = $1, subcategory_id = $2, sku = $3, image = $4, gallery = $5, name = $6, description = $7, price = $8 WHERE id = $9 RETURNING *"
	meta := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(meta).
		WithArgs(product.CategoryID, product.SubcategoryID,
			product.Sku, product.Image, product.Gallery, product.Name,
			product.Description, product.Price, product.ID).
		WillReturnRows(data).
		WillReturnError(nil)
	res, err := suite.repo.Update(context.TODO(), product)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *productRepositoryTestSuite) TestRepository_Delete_ExpectSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM products WHERE id = $1")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	data := &model.Product{ID: 1}
	err := suite.repo.Delete(context.TODO(), data)
	require.Nil(suite.T(), err)
}

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(productRepositoryTestSuite))
}
