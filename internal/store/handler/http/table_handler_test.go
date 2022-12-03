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

type tableHandlerTestSuite struct {
	suite.Suite
	table  *domain.Table
	tables []*domain.Table
}

func (suite *tableHandlerTestSuite) SetupSuite() {
	suite.table = &domain.Table{
		ID:       1,
		FloorId:  1,
		Name:     "lorem",
		XPos:     1,
		YPos:     1,
		WSize:    1,
		HSize:    1,
		Capacity: 1,
		Type:     "square",
	}

	suite.tables = []*domain.Table{
		suite.table,
		{
			ID:       2,
			FloorId:  1,
			Name:     "lorem",
			XPos:     1,
			YPos:     1,
			WSize:    1,
			HSize:    1,
			Capacity: 1,
			Type:     "round",
		},
	}
}

func (suite *tableHandlerTestSuite) TestTableHandler_Fetch_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("TableList").
		Return(suite.tables, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	tableHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)

	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Fetch_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("TableList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	tableHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Store_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddTable", mock.Anything).
		Return(suite.table, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"floor_id": 1,
		"name":     "lorem",
		"x_pos":    1,
		"y_pos":    1,
		"w_size":   1,
		"h_size":   1,
		"capacity": 1,
	})
	tableHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Store_ShouldError_BadRequest() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddTable", mock.Anything).
		Return(suite.table, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", nil)
	tableHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)

	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Store_ShouldError_Internal() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddTable", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"floor_id": 1,
		"name":     "lorem",
		"x_pos":    1,
		"y_pos":    1,
		"w_size":   1,
		"h_size":   1,
		"capacity": 1,
	})
	tableHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)

	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Update_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditTable", mock.Anything).
		Return(suite.table, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"floor_id": 1,
		"name":     "lorem",
		"x_pos":    1,
		"y_pos":    1,
		"w_size":   1,
		"h_size":   1,
		"capacity": 1,
	})
	tableHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)

	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Update_ShouldError_BadRequest() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditTable", mock.Anything).
		Return(suite.table, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", nil)
	tableHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)

	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Update_ShouldError_Internal() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditTable", mock.Anything).
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
		"floor_id": 1,
		"name":     "lorem",
		"x_pos":    1,
		"y_pos":    1,
		"w_size":   1,
		"h_size":   1,
		"capacity": 1,
	})
	tableHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)

	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Destroy_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("DeleteTable", mock.Anything).
		Return(nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "DELETE", "application/json", nil)
	tableHandler{svc: svcMock}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}

func (suite *tableHandlerTestSuite) TestTableHandler_Destroy_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("DeleteTable", mock.Anything).
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
	tableHandler{svc: svcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func TestTableHandlerService(t *testing.T) {
	suite.Run(t, new(tableHandlerTestSuite))
}
