package service_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/aasumitro/posbe/internal/catalog/service"
	"github.com/aasumitro/posbe/mocks"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type catalogCommonService struct {
	suite.Suite
	Db            *sql.DB
	unit          *model.Unit
	units         []*model.Unit
	category      *model.Category
	categories    []*model.Category
	subcategory   *model.Subcategory
	subcategories []*model.Subcategory
	addon         *model.Addon
	addons        []*model.Addon
	svcErr        *utils.ServiceError
}

func (suite *catalogCommonService) SetupSuite() {
	suite.unit = &model.Unit{ID: 1, Magnitude: "test", Name: "test", Symbol: "test"}
	suite.units = []*model.Unit{suite.unit, {ID: 2, Magnitude: "test 2", Name: "test 2", Symbol: "test 2"}}
	suite.category = &model.Category{ID: 1, Name: "test"}
	suite.categories = []*model.Category{suite.category, {ID: 1, Name: "test 2"}}
	suite.subcategory = &model.Subcategory{ID: 1, CategoryID: 1, Name: "test"}
	suite.subcategories = []*model.Subcategory{suite.subcategory, {ID: 2, CategoryID: 1, Name: "test 2"}}

	suite.addon = &model.Addon{ID: 1, Name: "test", Description: "test", Price: 1}
	suite.addons = []*model.Addon{suite.addon, {ID: 2, Name: "test 2", Description: "test 2", Price: 2}}
	suite.svcErr = &utils.ServiceError{Code: 500, Message: "UNEXPECTED"}
}

// === UNITS
func (suite *catalogCommonService) TestService_UnitList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("All", mock.Anything).
		Return(suite.units, nil).Once()
	data, err := svc.UnitList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.units)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_UnitList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.UnitList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddUnit_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.unit, nil).Once()
	data, err := svc.AddUnit(context.TODO(), suite.unit)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.unit)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddUnit_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddUnit(context.TODO(), suite.unit)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditUnit_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.unit, nil).Once()
	data, err := svc.EditUnit(context.TODO(), suite.unit)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.unit)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditUnit_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditUnit(context.TODO(), suite.unit)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteUnit_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.units[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteUnit(context.TODO(), suite.units[1])
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteUnit_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteUnit(context.TODO(), suite.unit)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteUnit_ShouldErrorWhenFindNotFound() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteUnit(context.TODO(), suite.unit)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteUnit_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[model.Unit])
	svc := service.NewCatalogCommonService(repoMock,
		new(mocks.ICRUDRepository[model.Category]), new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.units[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteUnit(context.TODO(), suite.unit)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

// === Category
func (suite *catalogCommonService) TestService_CategoryList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock, new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("All", mock.Anything).
		Return(suite.categories, nil).Once()
	data, err := svc.CategoryList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.categories)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_CategoryList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.CategoryList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddCategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.category, nil).Once()
	data, err := svc.AddCategory(context.TODO(), suite.category)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.category)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddCategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddCategory(context.TODO(), suite.category)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditCategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.category, nil).Once()
	data, err := svc.EditCategory(context.TODO(), suite.category)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.category)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditCategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditCategory(context.TODO(), suite.category)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteCategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.categories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteCategory(context.TODO(), suite.categories[1])
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteCategory_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteCategory(context.TODO(), suite.category)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteCategory_ShouldErrorWhenFindNotFound() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteCategory(context.TODO(), suite.category)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteCategory_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[model.Category])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), repoMock,
		new(mocks.ICRUDRepository[model.Subcategory]),
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.categories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteCategory(context.TODO(), suite.category)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

// === Subcategory
func (suite *catalogCommonService) TestService_SubcategoryList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("All", mock.Anything).
		Return(suite.subcategories, nil).Once()
	data, err := svc.SubcategoryList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.subcategories)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_SubcategoryList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.SubcategoryList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddSubcategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]), new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.subcategory, nil).Once()
	data, err := svc.AddSubcategory(context.TODO(), suite.subcategory)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.subcategory)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddSubcategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddSubcategory(context.TODO(), suite.subcategory)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditSubcategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.subcategory, nil).Once()
	data, err := svc.EditSubcategory(context.TODO(), suite.subcategory)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.subcategory)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditSubcategory_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditSubcategory(context.TODO(), suite.subcategory)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.subcategories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteSubcategory(context.TODO(), suite.subcategories[1])
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteSubcategory(context.TODO(), suite.subcategory)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldErrorWhenFindNotFound() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteSubcategory(context.TODO(), suite.subcategory)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteSubcategory_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[model.Subcategory])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]), repoMock,
		new(mocks.ICRUDRepository[model.Addon]))
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.subcategories[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteSubcategory(context.TODO(), suite.subcategory)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

// === Addon
func (suite *catalogCommonService) TestService_AddonList_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.On("All", mock.Anything).
		Return(suite.addons, nil).Once()
	data, err := svc.AddonList(context.TODO())
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.addons)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddonList_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.On("All", mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddonList(context.TODO())
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_AddAddon_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(suite.addon, nil).Once()
	data, err := svc.AddAddon(context.TODO(), suite.addon)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.addon)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_AddAddon_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.AddAddon(context.TODO(), suite.addon)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_EditAddon_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(suite.addon, nil).Once()
	data, err := svc.EditAddon(context.TODO(), suite.addon)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), data)
	require.Equal(suite.T(), data, suite.addon)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_EditAddon_ShouldError() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("UNEXPECTED")).Once()
	data, err := svc.EditAddon(context.TODO(), suite.addon)
	require.Nil(suite.T(), data)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func (suite *catalogCommonService) TestService_DeleteAddon_ShouldSuccess() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.addons[1], nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(nil).Once()
	err := svc.DeleteAddon(context.TODO(), suite.addon)
	require.Nil(suite.T(), err)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteAddon_ShouldErrorWhenFind() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, errors.New("UNEXPECTED"))
	err := svc.DeleteAddon(context.TODO(), suite.addon)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteAddon_ShouldErrorWhenFindNotFound() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Once().
		Return(nil, sql.ErrNoRows)
	err := svc.DeleteAddon(context.TODO(), suite.addon)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, &utils.ServiceError{Code: 404, Message: "sql: no rows in result set"})
	repoMock.AssertExpectations(suite.T())
}
func (suite *catalogCommonService) TestService_DeleteAddon_ShouldErrorWhenDelete() {
	repoMock := new(mocks.ICRUDRepository[model.Addon])
	svc := service.NewCatalogCommonService(
		new(mocks.ICRUDRepository[model.Unit]),
		new(mocks.ICRUDRepository[model.Category]),
		new(mocks.ICRUDRepository[model.Subcategory]), repoMock)
	repoMock.
		On("Find", mock.Anything, mock.Anything, mock.Anything).
		Return(suite.addon, nil).Once()
	repoMock.
		On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("UNEXPECTED")).Once()
	err := svc.DeleteAddon(context.TODO(), suite.addon)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), err, suite.svcErr)
	repoMock.AssertExpectations(suite.T())
}

func TestCatalogCommonService(t *testing.T) {
	suite.Run(t, new(catalogCommonService))
}
