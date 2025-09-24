package catalog

import (
	"context"
	"net/http"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type productService struct {
	repository IProductRepository
}

func (service productService) AddonList(ctx context.Context) ([]*model.Addon, *utils.ServiceError) {
	data, err := service.repository.GetAllAddon(ctx)
	return utils.HandleMultipleResults[model.Addon]("addons", data, err)
}

func (service productService) CreateAddon(ctx context.Context, form *AddonForm) *utils.ServiceError {
	if err := service.repository.CreateAddon(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service productService) UpdateAddon(ctx context.Context, form *AddonForm) *utils.ServiceError {
	if err := service.repository.UpdateAddon(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service productService) DeleteAddon(ctx context.Context, id int64) *utils.ServiceError {
	if err := service.repository.DeleteAddon(ctx, id); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func NewProductService(repository IProductRepository) IProductService {
	return &productService{repository: repository}
}
