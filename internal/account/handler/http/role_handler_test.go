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

type roleHandlerTestSuite struct {
	suite.Suite
	role  *domain.Role
	roles []*domain.Role
}

func (suite *roleHandlerTestSuite) SetupSuite() {
	suite.role = &domain.Role{
		ID:          1,
		Name:        "lorem",
		Description: "lorem ipsum",
		Usage:       1,
	}

	suite.roles = []*domain.Role{
		suite.role,
		{
			ID:          2,
			Name:        "dolor",
			Description: "Dolor Sit Amet",
		},
	}
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Fetch_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("RoleList").
		Return(suite.roles, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	roleHandler{svc: accSvcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Fetch_ShouldError() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("RoleList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	roleHandler{svc: accSvcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Store_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("AddRole", mock.Anything).
		Return(suite.role, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name":        "lorem",
		"description": "lorem ipsum",
	})
	roleHandler{svc: accSvcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Store_ShouldError_UnprocessableEntity() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("AddRole", mock.Anything).
		Return(suite.role, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", nil)
	roleHandler{svc: accSvcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Store_ShouldError_Internal() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("AddRole", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "POST", "application/json", map[string]interface{}{
		"name":        "lorem",
		"description": "lorem ipsum",
	})
	roleHandler{svc: accSvcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Update_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("EditRole", mock.Anything).
		Return(suite.role, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name":        "lorem",
		"description": "lorem ipsum",
	})
	roleHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Update_ShouldError_BadRequest() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd123"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	roleHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Update_ShouldError_UnprocessableEntity() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	roleHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Update_ShouldError_Internal() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("EditRole", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"name":        "lorem",
		"description": "lorem ipsum",
	})
	roleHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Destroy_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("DeleteRole", mock.Anything).
		Return(nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	roleHandler{svc: accSvcMock}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Destroy_ShouldError() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("DeleteRole", mock.Anything).
		Return(&utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	roleHandler{svc: accSvcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *roleHandlerTestSuite) TestRoleHandler_Destroy_ShouldError_BadRequest() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJSONRequest(ctx, "DELETE", "application/json", nil)
	roleHandler{svc: accSvcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func TestRoleHandlerService(t *testing.T) {
	suite.Run(t, new(roleHandlerTestSuite))
}
