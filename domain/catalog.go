package domain

import "github.com/aasumitro/posbe/pkg/utils"

type (
	Category struct {
		ID          int            `json:"id"`
		Name        string         `json:"name" form:"name" binding:"required"`
		Subcategory []*Subcategory `json:"subcategories,omitempty" binding:"-"`
	}

	Subcategory struct {
		ID         int    `json:"id"`
		CategoryId int    `json:"category_id" form:"category_id" binding:"required"`
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
		// TODO
	}
)
