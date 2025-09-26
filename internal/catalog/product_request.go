package catalog

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type AddonForm struct {
	ActionAdd   bool     `form:"-" json:"-"`
	ID          int64    `json:"-" form:"-"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Price       *float64 `json:"price" form:"price"`
}

func (f *AddonForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{}
	if f.ActionAdd {
		r = galidator.Rules{
			"Name":        g.R("name").Required(),
			"Description": g.R("description").Required(),
			"Price":       g.R("price").Optional(),
		}
	} else {
		r = galidator.Rules{
			"Name":        g.R("name").Optional(),
			"Description": g.R("description").Optional(),
			"Price":       g.R("price").Optional(),
		}
	}
	return g.ComplexValidator(r).Validate(ctx, f)
}

type NewProductForm struct {
	Status        string                   `form:"status" json:"status"`
	Image         string                   `form:"image" json:"image"`
	SKU           string                   `form:"sku" json:"sku"`
	Name          string                   `form:"name" json:"name"`
	CategoryID    int64                    `form:"category_id" json:"category_id"`
	SubcategoryID int64                    `form:"subcategory_id" json:"subcategory_id"`
	Description   string                   `form:"description" json:"description"`
	Variants      []*NewProductVariantForm `form:"variants" json:"variants"`
}

type NewProductVariantForm struct {
	Type        string  `form:"type" json:"type"`
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Price       float64 `json:"price" form:"price"`
	UnitID      int64   `form:"unit_id" json:"unit_id"`
	UnitSize    float64 `form:"unit_size" json:"unit_size"`
}

func (f *NewProductForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{
		"Status": g.R("status").Required().
			Choices("draft", "publish"),
		"Image":       g.R("image").Optional(),
		"SKU":         g.R("sku").Required(),
		"Name":        g.R("name").Required(),
		"Description": g.R("description").Optional(),
		"CategoryID": g.R("category_id").Required().Min(1).
			SpecificMessages(galidator.Messages{"min": "for published product, category_id is required"}),
		"SubcategoryID": g.R("subcategory_id").Required().Min(1).
			SpecificMessages(galidator.Messages{"min": "for published product, subcategory_id required"}),
	}

	if err := g.ComplexValidator(r).Validate(ctx, f); err != nil {
		return err
	}

	if f.Variants != nil {
		rules := galidator.Rules{
			"Type":        g.R("type").Required(),
			"Name":        g.R("name").Required(),
			"Description": g.R("description").Required(),
			"Price":       g.R("price").Optional(),
			"UnitID":      g.R("unit_id").Optional(),
			"UnitSize": g.R("unit_size").Optional().
				WhenExistOne("UnitID").SpecificMessages(galidator.Messages{
				"when_exist_one": "unit size is required",
			}),
		}
		validator := g.ComplexValidator(rules)

		for _, sub := range f.Variants {
			if err := validator.Validate(ctx, sub); err != nil {
				return err
			}
		}
	}

	return nil
}

type ProductUpdateForm struct {
	ID int64 `form:"-" json:"-"`
}

func (f *ProductUpdateForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{}).Validate(ctx, f)
}
