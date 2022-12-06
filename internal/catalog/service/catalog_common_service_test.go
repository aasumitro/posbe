package service_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/domain/mocks"
	"github.com/aasumitro/posbe/internal/catalog/service"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type catalogCommonService struct {
	suite.Suite
	Db            *sql.DB
	unit          *domain.Unit
	units         []*domain.Unit
	category      *domain.Category
	categories    []*domain.Category
	subcategory   *domain.Subcategory
	subcategories []*domain.Subcategory
	svcErr        *utils.ServiceError
}

func (suite *catalogCommonService) SetupSuite() {
	suite.unit = &domain.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	suite.units = []*domain.Unit{suite.unit, {ID: 2, Magnitude: "test 2", Name: "test 2", Symbol: "test 2"}}
	suite.category = &domain.Category{ID: 1, Name: "test"}
	suite.categories = []*domain.Category{suite.category, {ID: 1, Name: "test 2"}}
	suite.subcategory = &domain.Subcategory{ID: 1, CategoryId: 1, Name: "test"}
	suite.subcategories = []*domain.Subcategory{suite.subcategory, {ID: 2, CategoryId: 1, Name: "test 2"}}
	suite.svcErr = &utils.ServiceError{Code: 500, Message: "UNEXPECTED"}
}

// === UNITS
func (suite *catalogCommonService) TestService_UnitList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("All", mock.Anything).
		Return(suite.units, nil).Once()
	data, err := svc.UnitList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.units)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_UnitList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.UnitList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddUnit_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.unit, nil).Once()
	data, err := svc.AddUnit(suite.unit)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.unit)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddUnit_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddUnit(suite.unit)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditUnit_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.unit, nil).Once()
	data, err := svc.EditUnit(suite.unit)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.unit)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditUnit_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditUnit(suite.unit)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteUnit_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.units[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteUnit(suite.units[1])
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteUnit_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteUnit(suite.unit)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteUnit_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[domain.Unit])
	svc := service.NewCatalogCommonService(context.TODO(), repoMock,
		new(mocks.ICRUDRepository[domain.Category]), new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.units[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteUnit(suite.unit)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

// === Category
func (suite *catalogCommonService) TestService_CategoryList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("All", mock.Anything).
		Return(suite.categories, nil).Once()
	data, err := svc.CategoryList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.categories)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_CategoryList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.CategoryList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddCategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.category, nil).Once()
	data, err := svc.AddCategory(suite.category)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.category)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddCategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddCategory(suite.category)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditCategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.category, nil).Once()
	data, err := svc.EditCategory(suite.category)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.category)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditCategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditCategory(suite.category)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteCategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.categories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteCategory(suite.categories[1])
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteCategory_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteCategory(suite.category)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteCategory_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[domain.Category])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), repoMock, new(mocks.ICRUDRepository[domain.Subcategory]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.categories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteCategory(suite.category)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

// === Subcategory
func (suite *catalogCommonService) TestService_SubcategoryList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.On("All", mock.Anything).
		Return(suite.subcategories, nil).Once()
	data, err := svc.SubcategoryList()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.subcategories)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_SubcategoryList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.SubcategoryList()
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddSubcategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.subcategory, nil).Once()
	data, err := svc.AddSubcategory(suite.subcategory)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.subcategory)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddSubcategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddSubcategory(suite.subcategory)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditSubcategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.subcategory, nil).Once()
	data, err := svc.EditSubcategory(suite.subcategory)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.subcategory)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditSubcategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditSubcategory(suite.subcategory)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.subcategories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteSubcategory(suite.subcategories[1])
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteSubcategory(suite.subcategory)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[domain.Subcategory])
	svc := service.NewCatalogCommonService(context.TODO(),
		new(mocks.ICRUDRepository[domain.Unit]), new(mocks.ICRUDRepository[domain.Category]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.subcategories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteSubcategory(suite.subcategory)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func TestCatalogCommonService(t *testing.T) {
	suite.Run(t, new(catalogCommonService))
}
