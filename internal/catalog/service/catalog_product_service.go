package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
)

type catalogProductService struct {
	productRepo        model.ICRUDWithSearchRepository[model.Product]
	productVariantRepo model.ICRUDRepository[model.ProductVariant]
}

func (service catalogProductService) ProductSearch(
	ctx context.Context,
	keys []model.FindWith,
	values []any,
) (products []*model.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.Search(ctx, keys, values)
	return utils.ValidateDataRows[model.Product](data, err)
}

func (service catalogProductService) ProductList(
	ctx context.Context,
) (products []*model.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.All(ctx)
	return utils.ValidateDataRows[model.Product](data, err)
}

func (service catalogProductService) ProductDetail(
	ctx context.Context,
	id int,
) (product *model.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.Find(ctx, model.FindWithID, id)
	return utils.ValidateDataRow[model.Product](data, err)
}

func (service catalogProductService) AddProduct(
	ctx context.Context,
	item *model.Product,
) (product *model.Product, errData *utils.ServiceError) {
	variants := item.ProductVariants
	data, err := service.productRepo.Create(ctx, item)
	if data == nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusNotFound,
			Message: errors.New("product not found"),
		}
	}
	if len(variants) > 0 {
		for _, v := range variants {
			v.ProductID = data.ID
			_, _ = service.productVariantRepo.Create(ctx, v)
		}
	}
	return utils.ValidateDataRow[model.Product](data, err)
}

func (service catalogProductService) EditProduct(
	ctx context.Context,
	item *model.Product,
) (product *model.Product, errData *utils.ServiceError) {
	data, err := service.productRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Product](data, err)
}

func (service catalogProductService) DeleteProduct(
	ctx context.Context,
	item *model.Product,
) *utils.ServiceError {
	data, err := service.productRepo.Find(ctx, model.FindWithID, item.ID)
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
	if err := service.productRepo.Delete(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service catalogProductService) AddProductVariant(
	ctx context.Context,
	item *model.ProductVariant,
) (units *model.ProductVariant, errData *utils.ServiceError) {
	data, err := service.productVariantRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.ProductVariant](data, err)
}

func (service catalogProductService) EditProductVariant(
	ctx context.Context,
	item *model.ProductVariant,
) (units *model.ProductVariant, errData *utils.ServiceError) {
	data, err := service.productVariantRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.ProductVariant](data, err)
}

func (service catalogProductService) DeleteProductVariant(
	ctx context.Context,
	item *model.ProductVariant,
) *utils.ServiceError {
	data, err := service.productVariantRepo.Find(
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
	if err := service.productVariantRepo.Delete(ctx, data); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func NewCatalogProductService(
	productRepo model.ICRUDWithSearchRepository[model.Product],
	productVariantRepo model.ICRUDRepository[model.ProductVariant],
) model.ICatalogProductService {
	return &catalogProductService{
		productRepo:        productRepo,
		productVariantRepo: productVariantRepo,
	}
}
