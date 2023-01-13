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

type unitHandlerTestSuite struct {
	suite.Suite
	row  *domain.Unit
	rows []*domain.Unit
}

func (suite *unitHandlerTestSuite) SetupSuite() {
	suite.row = &domain.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	suite.rows = []*domain.Unit{suite.row, {ID: 2, Magnitude: "test 2", Name: "test 2", Symbol: "test 2"}}
}

func (suite *unitHandlerTestSuite) TestHandler_Fetch_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("UnitList").
		Return(suite.rows, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	unitHandler{svc: svc}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Fetch_ShouldError() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("UnitList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	unitHandler{svc: svc}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *unitHandlerTestSuite) TestHandler_Store_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("AddUnit", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name":      "lorem",
		"magnitude": "ipsum",
		"symbol":    "ipsum",
	})
	unitHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Store_ShouldError_UnprocessableEntity() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("AddUnit", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", nil)
	unitHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Store_ShouldError_Internal() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("AddUnit", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name":      "lorem",
		"magnitude": "ipsum",
		"symbol":    "ipsum",
	})
	unitHandler{svc: svc}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *unitHandlerTestSuite) TestHandler_Update_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("EditUnit", mock.Anything).
		Return(suite.row, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name":      "lorem",
		"magnitude": "ipsum",
		"symbol":    "ipsum",
	})
	unitHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Update_ShouldError_BadRequest() {
	svc := new(mocks.ICatalogCommonService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd123"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	unitHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Update_ShouldError_UnprocessableEntity() {
	svc := new(mocks.ICatalogCommonService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	unitHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Update_ShouldError_Internal() {
	svc := new(mocks.ICatalogCommonService)
	svc.On("EditUnit", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name":      "lorem",
		"magnitude": "ipsum",
		"symbol":    "ipsum",
	})
	unitHandler{svc: svc}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *unitHandlerTestSuite) TestHandler_Destroy_ShouldSuccess() {
	svc := new(mocks.ICatalogCommonService)
	svc.
		On("DeleteUnit", mock.Anything).
		Return(nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	unitHandler{svc: svc}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}
func (suite *unitHandlerTestSuite) TestHandler_Destroy_ShouldErrorInternal() {
	svc := new(mocks.ICatalogCommonService)
	svc.On("DeleteUnit", mock.Anything).
		Return(&utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	unitHandler{svc: svc}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}
func (suite *unitHandlerTestSuite) TestHandler_Destroy_ShouldErrorBadRequest() {
	svc := new(mocks.ICatalogCommonService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	unitHandler{svc: svc}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func TestUnitHandlerService(t *testing.T) {
	suite.Run(t, new(unitHandlerTestSuite))
}
