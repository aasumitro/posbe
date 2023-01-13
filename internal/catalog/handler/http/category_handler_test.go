package http

import (
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

type categoryHandlerTestSuite struct {
	suite.Suite
	row  *domain.Category
	rows []*domain.Category
}

func (suite *categoryHandlerTestSuite) SetupSuite() {
	suite.row = &domain.Category{ID: 1, Name: "test"}
	suite.rows = []*domain.Category{suite.row, {ID: 1, Name: "test 2"}}
}

func (suite *categoryHandlerTestSuite) TestHandler_Fetch_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("CategoryList").
		Return(suite.rows, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	categoryHandler{svc: svc}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Fetch_ShouldError() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("CategoryList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	categoryHandler{svc: svc}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *categoryHandlerTestSuite) TestHandler_Store_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("AddCategory", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	categoryHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Store_ShouldError_UnprocessableEntity() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("AddCategory", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", nil)
	categoryHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Store_ShouldError_Internal() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("AddCategory", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	categoryHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *categoryHandlerTestSuite) TestHandler_Update_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("EditCategory", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	categoryHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Update_ShouldError_BadRequest() {
	svc := new(mocks.ICatalogCommonService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd123"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	categoryHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Update_ShouldError_UnprocessableEntity() {
	svc := new(mocks.ICatalogCommonService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	categoryHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Update_ShouldError_Internal() {
	svc := new(mocks.ICatalogCommonService)
	svc.On("EditCategory", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	categoryHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *categoryHandlerTestSuite) TestHandler_Destroy_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("DeleteCategory", mock.Anything).
		Return(nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	categoryHandler{svc: svc}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}
func (suite *categoryHandlerTestSuite) TestHandler_Destroy_ShouldErrorInternal() {
	svc := new(mocks.ICatalogCommonService)
	svc.On("DeleteCategory", mock.Anything).
		Return(&utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	categoryHandler{svc: svc}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}
func (suite *categoryHandlerTestSuite) TestHandler_Destroy_ShouldErrorBadRequest() {
	svc := new(mocks.ICatalogCommonService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	categoryHandler{svc: svc}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func TestCategoryHandlerService(t *testing.T) {
	suite.Run(t, new(categoryHandlerTestSuite))
}
