package catalog

import (
	"context"
	"net/http"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type attributeService struct {
	repository IAttributeRepository
}

func (service attributeService) UnitList(ctx context.Context) ([]*model.Unit, *utils.ServiceError) {
	data, err := service.repository.GetAllUnit(ctx)
	return utils.HandleMultipleResults[model.Unit]("units", data, err)
}

func (service attributeService) CreateUnit(ctx context.Context, form *UnitForm) *utils.ServiceError {
	if err := service.repository.CreateUnit(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service attributeService) UpdateUnit(ctx context.Context, form *UnitForm) *utils.ServiceError {
	if err := service.repository.UpdateUnit(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service attributeService) DeleteUnit(ctx context.Context, id int64) *utils.ServiceError {
	// TODO: validate in use or not (has used by product variant)!

	if err := service.repository.DeleteUnit(ctx, id); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service attributeService) CategoryList(ctx context.Context, query *AttributeCategoryQuery) ([]*model.Category, *utils.ServiceError) {
	data, err := service.repository.GetAllCategory(ctx, query)
	return utils.HandleMultipleResults[model.Category]("categories", data, err)
}

func (service attributeService) CreateCategory(ctx context.Context, form *NewCategoryForm) *utils.ServiceError {
	if err := service.repository.CreateCategory(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service attributeService) UpdateCategory(ctx context.Context, form *EditCategoryForm) *utils.ServiceError {
	if err := service.repository.UpdateCategory(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service attributeService) DeleteCategory(ctx context.Context, id int64) *utils.ServiceError {
	// TODO: validate in use or not (has used by product)!

	if err := service.repository.DeleteCategory(ctx, id); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service attributeService) DeleteSubcategory(ctx context.Context, cid, sid int64) *utils.ServiceError {
	// TODO: validate in use or not (has used by product)!

	if err := service.repository.DeleteSubcategory(ctx, cid, sid); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func NewAttributeService(repository IAttributeRepository) IAttributeService {
	return &attributeService{repository: repository}
}
