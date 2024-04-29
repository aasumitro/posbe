package service_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/aasumitro/posbe/internal/catalog/service"
	mocks2 "github.com/aasumitro/posbe/mocks"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type catalogProductService struct {
	suite.Suite
	variant  *model.ProductVariant
	variants []*model.ProductVariant
	svcErr   *utils.ServiceError
}

func (suite *catalogProductService) SetupSuite() {
	suite.variant = &model.ProductVariant{ID: 1, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test", Description: sql.NullString{String: "test"}, Price: 12}
	suite.variants = []*model.ProductVariant{
		suite.variant, {
			ID: 2, ProductID: 1, UnitID: 1, UnitSize: 12, Type: "color", Name: "test 2", Description: sql.NullString{String: "test 2"}, Price: 12,
		},
	}
	suite.svcErr = &utils.ServiceError{Code: 500, Message: "UNEXPECTED"}
}

func (suite *catalogProductService) TestService_AddVariant_ShouldSuccess() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.variant, nil).Once()
	data, err := svc.AddProductVariant(context.TODO(), suite.variant)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.variant)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogProductService) TestService_AddVariant_ShouldError() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddProductVariant(context.TODO(), suite.variant)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogProductService) TestService_EditVariant_ShouldSuccess() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.variant, nil).Once()
	data, err := svc.EditProductVariant(context.TODO(), suite.variant)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.variant)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogProductService) TestService_EditVariant_ShouldError() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditProductVariant(context.TODO(), suite.variant)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogProductService) TestService_DeleteVariant_ShouldSuccess() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.variant, nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteProductVariant(context.TODO(), suite.variant)
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogProductService) TestService_DeleteVariant_ShouldErrorWhenFind() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteProductVariant(context.TODO(), suite.variant)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogProductService) TestService_DeleteVariant_ShouldErrorWhenFindNotFound() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteProductVariant(context.TODO(), suite.variant)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogProductService) TestService_DeleteVariant_ShouldErrorWhenDelete() {
	repoMock := new(mocks2.ICRUDRepository[model.ProductVariant])
	svc := service.NewCatalogProductService(
		new(mocks2.ICRUDWithSearchRepository[model.Product]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.variant, nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteProductVariant(context.TODO(), suite.variant)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func TestCatalogProductService(t *testing.T) {
	suite.Run(t, new(catalogProductService))
}
