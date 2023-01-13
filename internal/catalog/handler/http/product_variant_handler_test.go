package http

import (
	"database/sql"
	"encoding/json"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/mocks"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type productVariantHandlerTestSuite struct {
	suite.Suite
	row  *domain.ProductVariant
	rows []*domain.ProductVariant
}

func (suite *productVariantHandlerTestSuite) SetupSuite() {
	suite.row = &domain.ProductVariant{ID: 1, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test", Description: sql.NullString{String: "test"}, Price: 12}
	suite.rows = []*domain.ProductVariant{
		suite.row, {
			ID: 2, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test 2", Description: sql.NullString{String: "test 2"}, Price: 12,
		},
	}
}

func (suite *productVariantHandlerTestSuite) TestHandler_Store_ShouldSuccess() {
	svc := new(mocks.ICatalogProductService)
	svc.
		On("AddProductVariant", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"product_id": 1,
		"unit_id":    1,
		"unit_size":  1,
		"type":       "none",
		"name":       "lorem",
		"price":      1,
	})
	variantHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Store_ShouldError_UnprocessableEntity() {
	svc := new(mocks.ICatalogProductService)
	svc.
		On("AddProductVariant", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", nil)
	variantHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Store_ShouldError_Internal() {
	svc := new(mocks.ICatalogProductService)
	svc.
		On("AddProductVariant", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"product_id": 1,
		"unit_id":    1,
		"unit_size":  1,
		"type":       "none",
		"name":       "lorem",
		"price":      1,
	})
	variantHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *productVariantHandlerTestSuite) TestHandler_Update_ShouldSuccess() {
	svc := new(mocks.ICatalogProductService)
	svc.
		On("EditProductVariant", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"product_id": 1,
		"unit_id":    1,
		"unit_size":  1,
		"type":       "none",
		"name":       "lorem",
		"price":      1,
	})
	variantHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Update_ShouldError_BadRequest() {
	svc := new(mocks.ICatalogProductService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd123"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	variantHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Update_ShouldError_UnprocessableEntity() {
	svc := new(mocks.ICatalogProductService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	variantHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Update_ShouldError_Internal() {
	svc := new(mocks.ICatalogProductService)
	svc.On("EditProductVariant", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"product_id": 1,
		"unit_id":    1,
		"unit_size":  1,
		"type":       "none",
		"name":       "lorem",
		"price":      1,
	})
	variantHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *productVariantHandlerTestSuite) TestHandler_Destroy_ShouldSuccess() {
	svc := new(mocks.ICatalogProductService)
	svc.
		On("DeleteProductVariant", mock.Anything).
		Return(nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	variantHandler{svc: svc}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Destroy_ShouldErrorInternal() {
	svc := new(mocks.ICatalogProductService)
	svc.On("DeleteProductVariant", mock.Anything).
		Return(&utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	variantHandler{svc: svc}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}
func (suite *productVariantHandlerTestSuite) TestHandler_Destroy_ShouldErrorBadRequest() {
	svc := new(mocks.ICatalogProductService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	variantHandler{svc: svc}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func TestProductVariantHandlerService(t *testing.T) {
	suite.Run(t, new(productVariantHandlerTestSuite))
}
