package model

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/pkg/utils"
)

type (
	Category struct {
		ID          int            `json:"id"`
		Name        string         `json:"name" form:"name" binding:"required"`
		Subcategory []*Subcategory `json:"subcategories,omitempty" binding:"-"`
	}

	Subcategory struct {
		ID         int    `json:"id"`
		CategoryID int    `json:"category_id" form:"category_id" binding:"required"`
		Name       string `json:"name" form:"name" binding:"required"`
	}

	Unit struct {
		ID        int    `json:"id"`
		Magnitude string `json:"magnitude" form:"magnitude" binding:"required"` // e.g: mass [length, time]
		Name      string `json:"name" form:"name" binding:"required"`           // e.g: kilogram [metre, second]
		Symbol    string `json:"symbol" form:"symbol" binding:"required"`       // e.g: kg [m, s]
	}

	Addon struct {
		ID          int     `json:"id"`
		Name        string  `json:"name" form:"name" binding:"required"`
		Description string  `json:"description" form:"description"  binding:"required"`
		Price       float32 `json:"price" form:"price" binding:"required"`
	}

	Product struct {
		ID              int               `json:"id"`
		CategoryID      int               `json:"category_id" form:"category_id" binding:"required"`
		SubcategoryID   int               `json:"subcategory_id" form:"subcategory_id" binding:"required"`
		Sku             string            `json:"sku" form:"sku" binding:"required"`
		Image           sql.NullString    `json:"image" form:"image"`
		Gallery         sql.NullString    `json:"gallery" form:"gallery"`
		Name            string            `json:"name" form:"name" binding:"required"`
		Price           float32           `json:"price" form:"price" binding:"required"`
		Description     sql.NullString    `json:"description" form:"description"`
		Category        *Category         `json:"category,omitempty" binding:"-"`
		Subcategory     *Subcategory      `json:"subcategory,omitempty" binding:"-"`
		ProductVariants []*ProductVariant `json:"variants,omitempty" form:"variants" binding:"required"`
	}

	ProductVariant struct {
		ID          int            `json:"id"`
		ProductID   int            `json:"product_id" form:"product_id" binding:"required"`
		UnitID      int            `json:"unit_id" form:"product_id" binding:"required"`
		UnitSize    float32        `json:"unit_size" form:"unit_size" binding:"required"`
		Type        string         `json:"type" form:"type" binding:"required"`
		Name        string         `json:"name" form:"name" binding:"required"`
		Description sql.NullString `json:"description" form:"description"`
		Price       float32        `json:"price" form:"price" binding:"required"`
		Unit        *Unit          `json:"unit,omitempty" binding:"-"`
	}

	ICatalogCommonService interface {
		UnitList(ctx context.Context) (units []*Unit, errData *utils.ServiceError)
		AddUnit(ctx context.Context, data *Unit) (units *Unit, errData *utils.ServiceError)
		EditUnit(ctx context.Context, data *Unit) (units *Unit, errData *utils.ServiceError)
		DeleteUnit(ctx context.Context, data *Unit) *utils.ServiceError

		CategoryList(ctx context.Context) (units []*Category, errData *utils.ServiceError)
		AddCategory(ctx context.Context, data *Category) (units *Category, errData *utils.ServiceError)
		EditCategory(ctx context.Context, data *Category) (units *Category, errData *utils.ServiceError)
		DeleteCategory(ctx context.Context, data *Category) *utils.ServiceError

		SubcategoryList(ctx context.Context) (units []*Subcategory, errData *utils.ServiceError)
		AddSubcategory(ctx context.Context, data *Subcategory) (units *Subcategory, errData *utils.ServiceError)
		EditSubcategory(ctx context.Context, data *Subcategory) (units *Subcategory, errData *utils.ServiceError)
		DeleteSubcategory(ctx context.Context, data *Subcategory) *utils.ServiceError

		AddonList(ctx context.Context) (units []*Addon, errData *utils.ServiceError)
		AddAddon(ctx context.Context, data *Addon) (units *Addon, errData *utils.ServiceError)
		EditAddon(ctx context.Context, data *Addon) (units *Addon, errData *utils.ServiceError)
		DeleteAddon(ctx context.Context, data *Addon) *utils.ServiceError
	}

	ICatalogProductService interface {
		AddProductVariant(ctx context.Context, data *ProductVariant) (variant *ProductVariant, errData *utils.ServiceError)
		EditProductVariant(ctx context.Context, data *ProductVariant) (variant *ProductVariant, errData *utils.ServiceError)
		DeleteProductVariant(ctx context.Context, data *ProductVariant) *utils.ServiceError

		ProductSearch(ctx context.Context, keys []FindWith, values []any) (products []*Product, errData *utils.ServiceError)
		ProductList(ctx context.Context) (products []*Product, errData *utils.ServiceError)
		ProductDetail(ctx context.Context, id int) (product *Product, errData *utils.ServiceError)
		AddProduct(ctx context.Context, data *Product) (product *Product, errData *utils.ServiceError)
		EditProduct(ctx context.Context, data *Product) (product *Product, errData *utils.ServiceError)
		DeleteProduct(ctx context.Context, data *Product) *utils.ServiceError
	}
)
