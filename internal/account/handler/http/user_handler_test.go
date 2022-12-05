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

type userHandlerTestSuite struct {
	suite.Suite
	role  *domain.Role
	roles []*domain.Role
	user  *domain.User
	users []*domain.User
}

func (suite *userHandlerTestSuite) SetupSuite() {
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

	suite.user = &domain.User{
		ID:       1,
		RoleId:   1,
		Name:     "lorem ipsum",
		Username: "lorem",
		Email:    "lorem@ipsum.id",
		Phone:    "+628227111111",
		Role:     *suite.role,
	}

	suite.users = []*domain.User{
		suite.user,
		{
			ID:       2,
			RoleId:   1,
			Name:     "dolor amet",
			Username: "dolor",
			Email:    "dolor@amet.id",
			Phone:    "+628227222222",
			Role:     *suite.role,
			Password: "secret",
		},
	}
}

func (suite *userHandlerTestSuite) TestUserHandler_Fetch_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("UserList").
		Return(suite.users, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	userHandler{svc: accSvcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Fetch_ShouldError() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("UserList").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	userHandler{svc: accSvcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *userHandlerTestSuite) TestUserHandler_Show_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("ShowUser", mock.Anything).
		Return(suite.user, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "GET", "application/json", nil)
	userHandler{svc: accSvcMock}.show(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Show_ShouldError() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("ShowUser", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "GET", "application/json", nil)
	userHandler{svc: accSvcMock}.show(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *userHandlerTestSuite) TestUserHandler_Show_ShouldError_BadRequest() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJsonRequest(ctx, "GET", "application/json", nil)
	userHandler{svc: accSvcMock}.show(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Store_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("AddUser", mock.Anything).
		Return(suite.user, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"role_id":  1,
		"name":     "lorem ipsum",
		"username": "lorem",
		"email":    "lorem@ipsum.id",
		"password": "secret",
		"phone":    "82271111",
	})
	userHandler{svc: accSvcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Store_ShouldError_UnprocessableEntity() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", nil)
	userHandler{svc: accSvcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Store_ShouldError_Internal() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("AddUser", mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"role_id":  1,
		"name":     "lorem ipsum",
		"username": "lorem",
		"email":    "lorem@ipsum.id",
		"password": "secret",
		"phone":    "82271111",
	})
	userHandler{svc: accSvcMock}.store(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Update_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("EditUser", mock.Anything).
		Return(suite.user, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"role_id":  1,
		"name":     "lorem ipsum",
		"username": "lorem",
		"email":    "lorem@ipsum.id",
		"password": "secret",
		"phone":    "82271111",
	})
	userHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Update_ShouldError_BadRequest() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", nil)
	userHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Update_ShouldError_UnprocessableEntity() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "PUT", "application/json", nil)
	userHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Update_ShouldError_Internal() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("EditUser", mock.Anything).
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
		"role_id":  1,
		"name":     "lorem ipsum",
		"username": "lorem",
		"email":    "lorem@ipsum.id",
		"password": "secret",
		"phone":    "82271111",
	})
	userHandler{svc: accSvcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Destroy_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("DeleteUser", mock.Anything).
		Return(nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	utils.MockJsonRequest(ctx, "DELETE", "application/json", nil)
	userHandler{svc: accSvcMock}.destroy(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, writer.Code)
}

func (suite *userHandlerTestSuite) TestUserHandler_Destroy_ShouldError_BadRequest() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	ctx.Params = []gin.Param{{Key: "id", Value: "asd1"}}
	utils.MockJsonRequest(ctx, "DELETE", "application/json", nil)
	userHandler{svc: accSvcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusBadRequest, writer.Code)
	assert.Equal(suite.T(), http.StatusBadRequest, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusBadRequest), got.Status)
}

func (suite *userHandlerTestSuite) TestUserHandler_Destroy_ShouldError() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("DeleteUser", mock.Anything).
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
	userHandler{svc: accSvcMock}.destroy(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func TestUserHandlerService(t *testing.T) {
	suite.Run(t, new(userHandlerTestSuite))
}
