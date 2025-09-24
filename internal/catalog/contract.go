package catalog

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type IAttributeRepository interface {
	GetAllUnit(ctx context.Context) ([]*model.Unit, error)
	CreateUnit(ctx context.Context, form *UnitForm) error
	UpdateUnit(ctx context.Context, form *UnitForm) error
	DeleteUnit(ctx context.Context, id int64) error

	GetAllCategory(ctx context.Context) ([]*model.Category, error)
	CreateCategory(ctx context.Context, form *NewCategoryForm) error
	UpdateCategory(ctx context.Context, form *EditCategoryForm) error
	DeleteCategory(ctx context.Context, id int64) error
	DeleteSubcategory(ctx context.Context, cid, sid int64) error
}

type IAttributeService interface {
	UnitList(ctx context.Context) ([]*model.Unit, *utils.ServiceError)
	CreateUnit(ctx context.Context, form *UnitForm) *utils.ServiceError
	UpdateUnit(ctx context.Context, form *UnitForm) *utils.ServiceError
	DeleteUnit(ctx context.Context, id int64) *utils.ServiceError

	CategoryList(ctx context.Context) ([]*model.Category, *utils.ServiceError)
	CreateCategory(ctx context.Context, form *NewCategoryForm) *utils.ServiceError
	UpdateCategory(ctx context.Context, form *EditCategoryForm) *utils.ServiceError
	DeleteCategory(ctx context.Context, id int64) *utils.ServiceError
	DeleteSubcategory(ctx context.Context, cid, sid int64) *utils.ServiceError
}

type IProductRepository interface {
	GetAllAddon(ctx context.Context) ([]*model.Addon, error)
	CreateAddon(ctx context.Context, form *AddonForm) error
	UpdateAddon(ctx context.Context, form *AddonForm) error
	DeleteAddon(ctx context.Context, id int64) error
}

type IProductService interface {
	AddonList(ctx context.Context) ([]*model.Addon, *utils.ServiceError)
	CreateAddon(ctx context.Context, form *AddonForm) *utils.ServiceError
	UpdateAddon(ctx context.Context, form *AddonForm) *utils.ServiceError
	DeleteAddon(ctx context.Context, id int64) *utils.ServiceError
}
