package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
)

type catalogCommonService struct {
	unitRepo        model.ICRUDRepository[model.Unit]
	categoryRepo    model.ICRUDRepository[model.Category]
	subcategoryRepo model.ICRUDRepository[model.Subcategory]
	addonRepo       model.ICRUDRepository[model.Addon]
}

func (service catalogCommonService) UnitList(
	ctx context.Context,
) (
	units []*model.Unit,
	errData *utils.ServiceError,
) {
	data, err := service.unitRepo.All(ctx)
	return utils.ValidateDataRows[model.Unit](data, err)
}

func (service catalogCommonService) AddUnit(
	ctx context.Context,
	item *model.Unit,
) (
	units *model.Unit,
	errData *utils.ServiceError,
) {
	data, err := service.unitRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Unit](data, err)
}

func (service catalogCommonService) EditUnit(
	ctx context.Context,
	item *model.Unit,
) (
	units *model.Unit,
	errData *utils.ServiceError,
) {
	data, err := service.unitRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Unit](data, err)
}

func (service catalogCommonService) DeleteUnit(
	ctx context.Context,
	item *model.Unit,
) *utils.ServiceError {
	data, err := service.unitRepo.Find(ctx, model.FindWithID, item.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if err := service.unitRepo.Delete(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service catalogCommonService) CategoryList(
	ctx context.Context,
) (
	units []*model.Category,
	errData *utils.ServiceError,
) {
	data, err := service.categoryRepo.All(ctx)
	return utils.ValidateDataRows[model.Category](data, err)
}

func (service catalogCommonService) AddCategory(
	ctx context.Context,
	item *model.Category,
) (
	units *model.Category,
	errData *utils.ServiceError,
) {
	data, err := service.categoryRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Category](data, err)
}

func (service catalogCommonService) EditCategory(
	ctx context.Context,
	item *model.Category,
) (
	units *model.Category,
	errData *utils.ServiceError,
) {
	data, err := service.categoryRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Category](data, err)
}

func (service catalogCommonService) DeleteCategory(
	ctx context.Context,
	item *model.Category,
) *utils.ServiceError {
	data, err := service.categoryRepo.Find(ctx, model.FindWithID, item.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if err := service.categoryRepo.Delete(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service catalogCommonService) SubcategoryList(
	ctx context.Context,
) (units []*model.Subcategory, errData *utils.ServiceError) {
	data, err := service.subcategoryRepo.All(ctx)
	return utils.ValidateDataRows[model.Subcategory](data, err)
}

func (service catalogCommonService) AddSubcategory(
	ctx context.Context,
	item *model.Subcategory,
) (units *model.Subcategory, errData *utils.ServiceError) {
	data, err := service.subcategoryRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Subcategory](data, err)
}

func (service catalogCommonService) EditSubcategory(
	ctx context.Context,
	item *model.Subcategory,
) (units *model.Subcategory, errData *utils.ServiceError) {
	data, err := service.subcategoryRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Subcategory](data, err)
}

func (service catalogCommonService) DeleteSubcategory(
	ctx context.Context,
	item *model.Subcategory,
) *utils.ServiceError {
	data, err := service.subcategoryRepo.Find(
		ctx, model.FindWithID, item.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if err := service.subcategoryRepo.Delete(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service catalogCommonService) AddonList(
	ctx context.Context,
) (units []*model.Addon, errData *utils.ServiceError) {
	data, err := service.addonRepo.All(ctx)
	return utils.ValidateDataRows[model.Addon](data, err)
}

func (service catalogCommonService) AddAddon(
	ctx context.Context,
	item *model.Addon,
) (units *model.Addon, errData *utils.ServiceError) {
	data, err := service.addonRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Addon](data, err)
}

func (service catalogCommonService) EditAddon(
	ctx context.Context,
	item *model.Addon,
) (units *model.Addon, errData *utils.ServiceError) {
	data, err := service.addonRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Addon](data, err)
}

func (service catalogCommonService) DeleteAddon(
	ctx context.Context,
	item *model.Addon,
) *utils.ServiceError {
	data, err := service.addonRepo.Find(ctx, model.FindWithID, item.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := service.addonRepo.Delete(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func NewCatalogCommonService(
	unitRepo model.ICRUDRepository[model.Unit],
	categoryRepo model.ICRUDRepository[model.Category],
	subcategoryRepo model.ICRUDRepository[model.Subcategory],
	addonRepo model.ICRUDRepository[model.Addon],
) model.ICatalogCommonService {
	return &catalogCommonService{
		unitRepo:        unitRepo,
		categoryRepo:    categoryRepo,
		subcategoryRepo: subcategoryRepo,
		addonRepo:       addonRepo,
	}
}
