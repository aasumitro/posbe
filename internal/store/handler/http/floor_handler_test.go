package http

import (
	"encoding/json"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/domain/mocks"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type floorHandlerTestSuite struct {
	suite.Suite
	floor  *domain.Floor
	floors []*domain.Floor
}

func (suite *floorHandlerTestSuite) SetupSuite() {
	suite.floor = &domain.Floor{
		ID:   1,
		Name: "lorem",
	}

	suite.floors = []*domain.Floor{
		suite.floor,
		{
			ID:   2,
			Name: "lorem",
		},
	}
}

func (suite *floorHandlerTestSuite) TestFloorHandler_FetchWithTable_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("FloorsWithTables").
		Return(suite.floors, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	floorHandler{svc: svcMock}.floorsWithTables(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_FetchWithTable_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("FloorsWithTables").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	floorHandler{svc: svcMock}.floorsWithTables(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Fetch_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("FloorList").
		Return(suite.floors, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	floorHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Fetch_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("FloorList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	floorHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Store_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddFloor", mock.Anything).
		Return(suite.floor, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	floorHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Store_ShouldError_BadRequest() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddFloor", mock.Anything).
		Return(suite.floor, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", nil)
	floorHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Store_ShouldError_Internal() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddFloor", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	floorHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Update_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditFloor", mock.Anything).
		Return(suite.floor, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	floorHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Update_ShouldError_BadRequest() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditFloor", mock.Anything).
		Return(suite.floor, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", nil)
	floorHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Update_ShouldError_Internal() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditFloor", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name": "lorem",
	})
	floorHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Destroy_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("DeleteFloor", mock.Anything).
		Return(nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "DELETE", "application/json", nil)
	floorHandler{svc: svcMock}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}

func (suite *floorHandlerTestSuite) TestFloorHandler_Destroy_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("DeleteFloor", mock.Anything).
		Return(&utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "DELETE", "application/json", nil)
	floorHandler{svc: svcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func TestFloorHandlerService(t *testing.T) {
	suite.Run(t, new(floorHandlerTestSuite))
}
