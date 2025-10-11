package catalog

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

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

func (service productService) ProductList(ctx context.Context, query *ProductQuery) ([]*model.Product, *utils.ServiceError) {
	data, err := service.repository.GetAllProduct(ctx, query)
	return utils.HandleMultipleResults[model.Product]("products", data, err)
}

func (service productService) ProductDetail(ctx context.Context, id int64) (*model.Product, *utils.ServiceError) {
	data, err := service.repository.GetProductDetail(ctx, id)
	return utils.HandleSingleResult[model.Product]("product", data, err)
}

func (service productService) CreateProduct(ctx context.Context, form *NewProductForm) *utils.ServiceError {
	if len(form.Variants) == 1 && form.Variants[0].Price <= 0 {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "invalid product price: when there is only 1 variant, price must be greater than 0",
		}
	}

	if form.Status == "publish" {
		form.Status = "active"
	}

	if form.Image != "" {
		nn := strconv.FormatInt(time.Now().UnixMicro(), 10) // unique name
		filePath, err := utils.UploadBase64Asset(form.Image, "products", nn)
		if err != nil {
			return &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		form.Image = filePath
	}

	if err := service.repository.CreateProduct(ctx, form); err != nil {
		// rollback image if DB update fails
		if form.Image != "" {
			parts := strings.Split(form.Image, "/")
			if len(parts) == 2 {
				folder, name := parts[0], parts[1]
				_ = utils.DeleteAsset(folder, name)
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service productService) UpdateProduct(ctx context.Context, form *ProductUpdateForm) *utils.ServiceError {
	product, errSvc := service.ProductDetail(ctx, form.ID)
	if errSvc != nil {
		return errSvc
	}

	if len(product.Variants) == 1 && len(form.Variants) == 1 && form.Variants[0].Price <= 0 {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "invalid product price: when there is only 1 variant, price must be greater than 0",
		}
	}

	if form.Status == "publish" {
		form.Status = "active"
	}

	if form.Image != "" {
		// upload new image
		nn := strconv.FormatInt(time.Now().UnixMicro(), 10) // unique name
		filePath, err := utils.UploadBase64Asset(form.Image, "products", nn)
		if err != nil {
			return &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		form.Image = filePath
	}

	if err := service.repository.UpdateProduct(ctx, form); err != nil {
		// rollback image if DB update fails
		if form.Image != "" {
			parts := strings.Split(form.Image, "/")
			if len(parts) == 2 {
				folder, name := parts[0], parts[1]
				_ = utils.DeleteAsset(folder, name)
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// remove the old image
	if form.Image != "" && product.Image != "" {
		parts := strings.Split(product.Image, "/")
		if len(parts) == 2 {
			folder, name := parts[0], parts[1]
			_ = utils.DeleteAsset(folder, name)
		}
	}

	return nil
}

func (service productService) DeleteProduct(ctx context.Context, id int64) *utils.ServiceError {
	product, err := service.ProductDetail(ctx, id)
	if err != nil {
		return err
	}

	if err := service.repository.DeleteProduct(ctx, product.ID); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if product.Image != "" {
		parts := strings.Split(product.Image, "/")
		if len(parts) == 2 {
			folder, name := parts[0], parts[1]
			_ = utils.DeleteAsset(folder, name)
		}
	}

	return nil
}

func (service productService) DeleteProductVariant(ctx context.Context, pid, vid int64) *utils.ServiceError {
	product, err := service.ProductDetail(ctx, pid)
	if err != nil {
		return err
	}

	// Verify variant exists
	var found bool
	for _, v := range product.Variants {
		if v.ID == vid {
			found = true
			break
		}
	}

	if !found {
		return &utils.ServiceError{
			Code:    http.StatusNotFound,
			Message: "variant not found",
		}
	}

	// Prevent deleting last variant
	if len(product.Variants) == 1 {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "a product must always have at least one variant",
		}
	}

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
