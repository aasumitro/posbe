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

func (service productService) ProductList(ctx context.Context) ([]*model.Product, *utils.ServiceError) {
	data, err := service.repository.GetAllProduct(ctx)
	return utils.HandleMultipleResults[model.Product]("products", data, err)
}

func (service productService) ProductDetail(ctx context.Context, id int64) (*model.Product, *utils.ServiceError) {
	data, err := service.repository.GetProductDetail(ctx, id)
	return utils.HandleSingleResult[model.Product]("product", data, err)
}

func (service productService) CreateProduct(ctx context.Context, form *NewProductForm) *utils.ServiceError {
	if err := service.repository.CreateProduct(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service productService) UpdateProduct(ctx context.Context, form *ProductUpdateForm) *utils.ServiceError {
	if err := service.repository.UpdateProduct(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service productService) DeleteProduct(ctx context.Context, id int64) *utils.ServiceError {
	// TODO: validate in use or not (has used by product)!

	if err := service.repository.DeleteProduct(ctx, id); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service productService) DeleteProductVariant(ctx context.Context, pid, vid int64) *utils.ServiceError {
	// TODO: validate in use or not (has used by order_product)!

	if err := service.repository.DeleteProductVariant(ctx, pid, vid); err != nil {
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
