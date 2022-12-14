package service

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
)

type catalogProductService struct {
	ctx                context.Context
	productRepo        domain.ICRUDWithSearchRepository[domain.Product]
	productVariantRepo domain.ICRUDRepository[domain.ProductVariant]
}

func (service catalogProductService) ProductSearch(keys []domain.FindWith, values []any) (products []*domain.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.Search(service.ctx, keys, values)

	return utils.ValidateDataRows[domain.Product](data, err)
}

func (service catalogProductService) ProductList() (products []*domain.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Product](data, err)
}

func (service catalogProductService) ProductDetail(id int) (product *domain.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.Find(service.ctx, domain.FindWithId, id)

	return utils.ValidateDataRow[domain.Product](data, err)
}

func (service catalogProductService) AddProduct(data *domain.Product) (product *domain.Product, errData *utils.ServiceError) {
	variants := data.ProductVariants
	data, err := service.productRepo.Create(service.ctx, data)

	if len(variants) > 0 {
		for _, v := range variants {
			v.ProductID = data.ID
			_, _ = service.productVariantRepo.Create(service.ctx, v)
		}
	}

	return utils.ValidateDataRow[domain.Product](data, err)
}

func (service catalogProductService) EditProduct(data *domain.Product) (product *domain.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Product](data, err)
}

func (service catalogProductService) DeleteProduct(data *domain.Product) *utils.ServiceError {
	data, err := service.productRepo.Find(service.ctx, domain.FindWithId, data.ID)
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

	err = service.productRepo.Delete(service.ctx, data)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service catalogProductService) AddProductVariant(data *domain.ProductVariant) (units *domain.ProductVariant, errData *utils.ServiceError) {
	data, err := service.productVariantRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.ProductVariant](data, err)
}

func (service catalogProductService) EditProductVariant(data *domain.ProductVariant) (units *domain.ProductVariant, errData *utils.ServiceError) {
	data, err := service.productVariantRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.ProductVariant](data, err)
}

func (service catalogProductService) DeleteProductVariant(data *domain.ProductVariant) *utils.ServiceError {
	data, err := service.productVariantRepo.Find(service.ctx, domain.FindWithId, data.ID)
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

	err = service.productVariantRepo.Delete(service.ctx, data)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func NewCatalogProductService(
	ctx context.Context,
	productRepo domain.ICRUDWithSearchRepository[domain.Product],
	productVariantRepo domain.ICRUDRepository[domain.ProductVariant],
) domain.ICatalogProductService {
	return &catalogProductService{
		ctx:                ctx,
		productRepo:        productRepo,
		productVariantRepo: productVariantRepo,
	}
}
