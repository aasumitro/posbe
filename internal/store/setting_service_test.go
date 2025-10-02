package store_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/store"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/aasumitro/posbe/mocks"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type settingServiceTestSuite struct {
	suite.Suite
	svcErr *utils.ServiceError
}

func (suite *settingServiceTestSuite) SetupTest() {
	config.RedisCache = redis.NewClient(&redis.Options{
		Addr: miniredis.RunT(suite.T()).Addr(),
	})
	suite.svcErr = &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	}
}

func (suite *settingServiceTestSuite) TestSettingService_AllSetting_ExpectSuccess() {
	cacheMock := new(mocks.MockICacheUtil[model.StoreSetting])
	repoMock := new(mocks.MockIStoreSettingRepository)
	data := &model.StoreSetting{"k": "v"}
	settSvc := store.NewSettingService(repoMock)
	repoMock.
		On("GetAll", mock.Anything).
		Return(data, nil).Once()
	cacheMock.On("CacheFirstData", mock.Anything).
		Return(data, nil).Once()
	data, err := settSvc.AllSetting(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, data)
	repoMock.AssertExpectations(suite.T())
}
func (suite *settingServiceTestSuite) TestSettingService_AllSetting_ExpectError() {
	repoMock := new(mocks.MockIStoreSettingRepository)
	settSvc := store.NewSettingService(repoMock)
	repoMock.On("GetAll", mock.Anything).
		Once().Return(nil, errors.New("UNEXPECTED"))
	data, err := settSvc.AllSetting(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *settingServiceTestSuite) TestSettingService_UpdateSetting_ExpectSuccess() {
	repoMock := new(mocks.MockIStoreSettingRepository)
	data := &store.SettingForm{Name: "test"}
	settSvc := store.NewSettingService(repoMock)
	repoMock.
		On("Update", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := settSvc.UpdateSetting(context.TODO(), data)
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *settingServiceTestSuite) TestSettingService_UpdateSetting_ExpectError() {
	repoMock := new(mocks.MockIStoreSettingRepository)
	data := &store.SettingForm{Name: "test"}
	settSvc := store.NewSettingService(repoMock)
	suite.T().Run("no form value", func(t *testing.T) {
		err := settSvc.UpdateSetting(context.TODO(), &store.SettingForm{})
		require.NotNil(suite.T(), err)
		require.Equal(suite.T(), err, &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "form value is required",
		})
	})
	suite.T().Run("error update", func(t *testing.T) {
		repoMock.
			On("Update", mock.Anything, mock.Anything).
			Return(errors.New("UNEXPECTED")).Once()
		err := settSvc.UpdateSetting(context.TODO(), data)
		require.NotNil(suite.T(), err)
		require.Equal(suite.T(), err, suite.svcErr)
	})
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(settingServiceTestSuite))
}
