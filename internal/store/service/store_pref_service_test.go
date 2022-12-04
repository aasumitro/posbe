package service_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/domain/mocks"
	"github.com/aasumitro/posbe/internal/store/service"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type storePrefTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *storePrefTestSuite) TestStorePrefService_AllPrefs_ShouldSuccess() {
	repo := new(mocks.IStorePrefRepository)
	svc := service.NewStorePrefService(context.TODO(), repo)
	repo.
		On("All", mock.Anything).
		Once().
		Return(&domain.StoreSetting{
			"lorem": "lorem",
			"ipsum": "ipsum",
		}, nil)
	data, err := svc.AllPrefs()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, &domain.StoreSetting{
		"lorem": "lorem",
		"ipsum": "ipsum",
	})
	repo.AssertExpectations(suite.T())
}

func (suite *storePrefTestSuite) TestStorePrefService_AllPrefs_ShouldError() {
	repo := new(mocks.IStorePrefRepository)
	svc := service.NewStorePrefService(context.TODO(), repo)
	repo.
		On("All", mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.AllPrefs()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	})
	repo.AssertExpectations(suite.T())
}

func (suite *storePrefTestSuite) TestStorePrefService_UpdatePrefs_ShouldSuccess() {
	repo := new(mocks.IStorePrefRepository)
	svc := service.NewStorePrefService(context.TODO(), repo)
	repo.
		On("Find", mock.Anything, mock.Anything).
		Once().
		Return(&domain.StoreSetting{"lorem": "lorem"}, nil)
	repo.
		On("Update", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(&domain.StoreSetting{"lorem": "lorem"}, nil)
	data, err := svc.UpdatePrefs("test", "test")
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, &domain.StoreSetting{
		"lorem": "lorem",
	})
	repo.AssertExpectations(suite.T())
}

func (suite *storePrefTestSuite) TestStorePrefService_UpdatePrefs_ShouldErrorWhenFind() {
	repo := new(mocks.IStorePrefRepository)
	svc := service.NewStorePrefService(context.TODO(), repo)
	repo.
		On("Find", mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.UpdatePrefs("test", "test")
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	})
	repo.AssertExpectations(suite.T())
}

func (suite *storePrefTestSuite) TestStorePrefService_UpdatePrefs_ShouldErrorWhenDelete() {
	repo := new(mocks.IStorePrefRepository)
	svc := service.NewStorePrefService(context.TODO(), repo)
	repo.
		On("Find", mock.Anything, mock.Anything).
		Once().
		Return(&domain.StoreSetting{"lorem": "lorem"}, nil)
	repo.
		On("Update", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	data, err := svc.UpdatePrefs("test", "test")
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{
		Code:    500,
		Message: "UNEXPECTED",
	})
	repo.AssertExpectations(suite.T())
}

func TestStorePrefService(t *testing.T) {
	suite.Run(t, new(storePrefTestSuite))
}
