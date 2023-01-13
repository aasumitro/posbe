package domain

import (
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
		UnitList() (units []*Unit, errData *utils.ServiceError)
		AddUnit(data *Unit) (units *Unit, errData *utils.ServiceError)
		EditUnit(data *Unit) (units *Unit, errData *utils.ServiceError)
		DeleteUnit(data *Unit) *utils.ServiceError

		CategoryList() (units []*Category, errData *utils.ServiceError)
		AddCategory(data *Category) (units *Category, errData *utils.ServiceError)
		EditCategory(data *Category) (units *Category, errData *utils.ServiceError)
		DeleteCategory(data *Category) *utils.ServiceError

		SubcategoryList() (units []*Subcategory, errData *utils.ServiceError)
		AddSubcategory(data *Subcategory) (units *Subcategory, errData *utils.ServiceError)
		EditSubcategory(data *Subcategory) (units *Subcategory, errData *utils.ServiceError)
		DeleteSubcategory(data *Subcategory) *utils.ServiceError

		AddonList() (units []*Addon, errData *utils.ServiceError)
		AddAddon(data *Addon) (units *Addon, errData *utils.ServiceError)
		EditAddon(data *Addon) (units *Addon, errData *utils.ServiceError)
		DeleteAddon(data *Addon) *utils.ServiceError
	}

	ICatalogProductService interface {
		AddProductVariant(data *ProductVariant) (variant *ProductVariant, errData *utils.ServiceError)
		EditProductVariant(data *ProductVariant) (variant *ProductVariant, errData *utils.ServiceError)
		DeleteProductVariant(data *ProductVariant) *utils.ServiceError

		ProductSearch(keys []FindWith, values []any) (products []*Product, errData *utils.ServiceError)
		ProductList() (products []*Product, errData *utils.ServiceError)
		ProductDetail(id int) (product *Product, errData *utils.ServiceError)
		AddProduct(data *Product) (product *Product, errData *utils.ServiceError)
		EditProduct(data *Product) (product *Product, errData *utils.ServiceError)
		DeleteProduct(data *Product) *utils.ServiceError
	}
)
