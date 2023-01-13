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

type storePrefHandlerTestSuite struct {
	suite.Suite
}

func (suite *storePrefHandlerTestSuite) TestStorePref_Fetch_ShouldSuccess() {
	svcMock := new(mocks.IStorePrefService)
	svcMock.
		On("AllPrefs").
		Return(&domain.StoreSetting{
			"lorem": "lorem",
			"ipsum": "ipsum",
		}, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	storePrefHandler{svc: svcMock}.fetch(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *storePrefHandlerTestSuite) TestStorePref_Fetch_ShouldError() {
	svcMock := new(mocks.IStorePrefService)
	svcMock.
		On("AllPrefs").
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	storePrefHandler{svc: svcMock}.fetch(ctx)
	var got utils.ErrorRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
	assert.Equal(suite.T(), "UNEXPECTED_ERROR", got.Data)
}

func (suite *storePrefHandlerTestSuite) TestStorePref_Update_ShouldSuccess() {
	svcMock := new(mocks.IStorePrefService)
	svcMock.
		On("UpdatePrefs", mock.Anything, mock.Anything).
		Return(&domain.StoreSetting{
			"lorem": "lorem",
		}, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"key":   "lorem",
		"value": "lorem",
	})
	storePrefHandler{svc: svcMock}.update(ctx)
	var got utils.SuccessRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusOK, writer.Code)
	assert.Equal(suite.T(), http.StatusOK, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusOK), got.Status)
}

func (suite *storePrefHandlerTestSuite) TestStorePref_Update_ShouldBadRequest() {
	svcMock := new(mocks.IStorePrefService)
	svcMock.
		On("UpdatePrefs", mock.Anything, mock.Anything).
		Return(&domain.StoreSetting{
			"lorem": "lorem",
		}, nil).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "PUT", "application/json", nil)
	storePrefHandler{svc: svcMock}.update(ctx)
	var got utils.ErrorRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, writer.Code)
	assert.Equal(suite.T(), http.StatusUnprocessableEntity, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusUnprocessableEntity), got.Status)
}

func (suite *storePrefHandlerTestSuite) TestStorePref_Update_ShouldInternalError() {
	svcMock := new(mocks.IStorePrefService)
	svcMock.
		On("UpdatePrefs", mock.Anything, mock.Anything).
		Return(nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "UNEXPECTED_ERROR",
		}).
		Once()
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	ctx.Request = &http.Request{Header: make(http.Header)}
	utils.MockJSONRequest(ctx, "PUT", "application/json", map[string]interface{}{
		"key":   "lorem",
		"value": "lorem",
	})
	storePrefHandler{svc: svcMock}.update(ctx)
	var got utils.ErrorRespond
	_ = json.Unmarshal(writer.Body.Bytes(), &got)
	assert.Equal(suite.T(), http.StatusInternalServerError, writer.Code)
	assert.Equal(suite.T(), http.StatusInternalServerError, got.Code)
	assert.Equal(suite.T(), http.StatusText(http.StatusInternalServerError), got.Status)
}

func TestStorePrefHandler(t *testing.T) {
	suite.Run(t, new(storePrefHandlerTestSuite))
}
