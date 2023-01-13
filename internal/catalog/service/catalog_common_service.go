package service

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
)

type catalogCommonService struct {
	ctx             context.Context
	unitRepo        domain.ICRUDRepository[domain.Unit]
	categoryRepo    domain.ICRUDRepository[domain.Category]
	subcategoryRepo domain.ICRUDRepository[domain.Subcategory]
	addonRepo       domain.ICRUDRepository[domain.Addon]
}

func (service catalogCommonService) UnitList() (units []*domain.Unit, errData *utils.ServiceError) {
	data, err := service.unitRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Unit](data, err)
}

func (service catalogCommonService) AddUnit(data *domain.Unit) (units *domain.Unit, errData *utils.ServiceError) {
	data, err := service.unitRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Unit](data, err)
}

func (service catalogCommonService) EditUnit(data *domain.Unit) (units *domain.Unit, errData *utils.ServiceError) {
	data, err := service.unitRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Unit](data, err)
}

func (service catalogCommonService) DeleteUnit(data *domain.Unit) *utils.ServiceError {
	data, err := service.unitRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	err = service.unitRepo.Delete(service.ctx, data)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service catalogCommonService) CategoryList() (units []*domain.Category, errData *utils.ServiceError) {
	data, err := service.categoryRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Category](data, err)
}

func (service catalogCommonService) AddCategory(data *domain.Category) (units *domain.Category, errData *utils.ServiceError) {
	data, err := service.categoryRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Category](data, err)
}

func (service catalogCommonService) EditCategory(data *domain.Category) (units *domain.Category, errData *utils.ServiceError) {
	data, err := service.categoryRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Category](data, err)
}

func (service catalogCommonService) DeleteCategory(data *domain.Category) *utils.ServiceError {
	data, err := service.categoryRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	err = service.categoryRepo.Delete(service.ctx, data)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service catalogCommonService) SubcategoryList() (units []*domain.Subcategory, errData *utils.ServiceError) {
	data, err := service.subcategoryRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Subcategory](data, err)
}

func (service catalogCommonService) AddSubcategory(data *domain.Subcategory) (units *domain.Subcategory, errData *utils.ServiceError) {
	data, err := service.subcategoryRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Subcategory](data, err)
}

func (service catalogCommonService) EditSubcategory(data *domain.Subcategory) (units *domain.Subcategory, errData *utils.ServiceError) {
	data, err := service.subcategoryRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Subcategory](data, err)
}

func (service catalogCommonService) DeleteSubcategory(data *domain.Subcategory) *utils.ServiceError {
	data, err := service.subcategoryRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	err = service.subcategoryRepo.Delete(service.ctx, data)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service catalogCommonService) AddonList() (units []*domain.Addon, errData *utils.ServiceError) {
	data, err := service.addonRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Addon](data, err)
}

func (service catalogCommonService) AddAddon(data *domain.Addon) (units *domain.Addon, errData *utils.ServiceError) {
	data, err := service.addonRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Addon](data, err)
}

func (service catalogCommonService) EditAddon(data *domain.Addon) (units *domain.Addon, errData *utils.ServiceError) {
	data, err := service.addonRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Addon](data, err)
}

func (service catalogCommonService) DeleteAddon(data *domain.Addon) *utils.ServiceError {
	data, err := service.addonRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	err = service.addonRepo.Delete(service.ctx, data)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func NewCatalogCommonService(
	ctx context.Context,
	unitRepo domain.ICRUDRepository[domain.Unit],
	categoryRepo domain.ICRUDRepository[domain.Category],
	subcategoryRepo domain.ICRUDRepository[domain.Subcategory],
	addonRepo domain.ICRUDRepository[domain.Addon],
) domain.ICatalogCommonService {
	return &catalogCommonService{
		ctx:             ctx,
		unitRepo:        unitRepo,
		categoryRepo:    categoryRepo,
		subcategoryRepo: subcategoryRepo,
		addonRepo:       addonRepo,
	}
}
