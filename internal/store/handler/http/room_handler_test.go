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

type roomHandlerTestSuite struct {
	suite.Suite
	room  *domain.Room
	rooms []*domain.Room
}

func (suite *roomHandlerTestSuite) SetupSuite() {
	suite.room = &domain.Room{
		ID:       1,
		FloorId:  1,
		Name:     "lorem",
		XPos:     1,
		YPos:     1,
		WSize:    1,
		HSize:    1,
		Capacity: 1,
		Price:    1,
	}

	suite.rooms = []*domain.Room{
		suite.room,
		{
			ID:       2,
			FloorId:  1,
			Name:     "lorem",
			XPos:     1,
			YPos:     1,
			WSize:    1,
			HSize:    1,
			Capacity: 1,
			Price:    1,
		},
	}
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Fetch_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("RoomList").
		Return(suite.rooms, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	roomHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Fetch_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("RoomList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	roomHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Store_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddRoom", mock.Anything).
		Return(suite.room, nil).
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
		"price":    1,
	})
	roomHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Store_ShouldError_BadRequest() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddRoom", mock.Anything).
		Return(suite.room, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", nil)
	roomHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Store_ShouldError_Internal() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("AddRoom", mock.Anything).
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
		"price":    1,
	})
	roomHandler{svc: svcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Update_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditRoom", mock.Anything).
		Return(suite.room, nil).
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
		"price":    2,
	})
	roomHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Update_ShouldError_BadRequest() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditRoom", mock.Anything).
		Return(suite.room, nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", nil)
	roomHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Update_ShouldError_Internal() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("EditRoom", mock.Anything).
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
		"price":    1,
	})
	roomHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Destroy_ShouldSuccess() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("DeleteRoom", mock.Anything).
		Return(nil).
		Once()

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "DELETE", "application/json", nil)
	roomHandler{svc: svcMock}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}

func (suite *roomHandlerTestSuite) TestRoomHandler_Destroy_ShouldError() {
	svcMock := new(mocks.IStoreService)
	svcMock.
		On("DeleteRoom", mock.Anything).
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
	roomHandler{svc: svcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func TestRoomHandler(t *testing.T) {
	suite.Run(t, new(roomHandlerTestSuite))
}
