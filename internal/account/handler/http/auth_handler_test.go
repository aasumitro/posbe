package http

import (
	"encoding/json"
	"fmt"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/domain/mocks"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type authHandlerTestSuite struct {
	suite.Suite
	role  *domain.Role
	user  *domain.User
	login  *domain.LoginForm
}

func (suite *authHandlerTestSuite) SetupSuite() {
	suite.login = &domain.LoginForm{
		Username: "lorem",
		Password: "secret",
	}

	suite.role = &domain.Role{
		ID:          1,
		Name:        "lorem",
		Description: "lorem ipsum",
		Usage:       1,
	}

	suite.user = &domain.User{
		ID:       1,
		RoleId:   1,
		Name:     "lorem ipsum",
		Username: "lorem",
		Email:    "lorem@ipsum.id",
		Phone:    "+628227111111",
		Password: "2ad1a22d5b3c9396d16243d2fe7f067976363715e322203a456278bb80b0b4a4.7ab4dcccfcd9d36efc68f1626d2fb80804a6508f9c3a7b44f430ba082b6870d2",
		Role:     *suite.role,
	}
}

func (suite *authHandlerTestSuite) TestAuthHandler_Login_ShouldSuccess() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("VerifyUserCredentials", mock.Anything, mock.Anything).
		Return(suite.user, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"username": "lorem",
		"password": "secret",
	})
	AuthHandler{svc: accSvcMock, config: &config.Config{
		AppName: "test",
		JWTLifetime: 1,
		JWTSecretKey: "secret",
	}}.login(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusCreated, writer.Code)
	assert.Equal(suite.T(), http.StatusCreated, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusCreated), got.Status)
}

func (suite *authHandlerTestSuite) TestAuthHandler_Login_ShouldErrorEntity() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("VerifyUserCredentials", mock.Anything, mock.Anything).
		Return(suite.user, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", nil)
	AuthHandler{svc: accSvcMock, config: &config.Config{
		AppName: "test",
		JWTLifetime: 1,
		JWTSecretKey: "secret",
	}}.login(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *authHandlerTestSuite) TestAuthHandler_Login_ShouldErrorInternalWhenVerify() {
	accSvcMock := new(mocks.IAccountService)
	accSvcMock.
		On("VerifyUserCredentials", mock.Anything, mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", map[string]interface{}{
		"username": "lorem",
		"password": "secret",
	})
	AuthHandler{svc: accSvcMock, config: &config.Config{
		AppName: "test",
		JWTLifetime: 1,
		JWTSecretKey: "secret",
	}}.login(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

// TODO VALIDATE CLAIM JWT

func (suite *authHandlerTestSuite) TestAuthHandler_Logout() {
	accSvcMock := new(mocks.IAccountService)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJsonRequest(ctx, "POST", "application/json", nil)
	ctx.Request = &http.Request{Header: make(http.Header)}
	AuthHandler{svc: accSvcMock, config: &config.Config{
		AppName: "test",
		JWTLifetime: 1,
		JWTSecretKey: "secret",
	}}.logout(ctx)
	var got utils.SuccessRespond
	fmt.Println(got)
	//assert.Equal(suite.T(), http.StatusOK, writer.Code)
	//assert.Equal(suite.T(), http.StatusOK, got.Code)
	//assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func TestAuthHandlerService(t *testing.T) {
	suite.Run(t, new(authHandlerTestSuite))
}
