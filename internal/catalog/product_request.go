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
}

func (f *NewProductForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{}).Validate(ctx, f)
}

type ProductUpdateForm struct {
	ID int64 `form:"-" json:"-"`
}

func (f *ProductUpdateForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{}).Validate(ctx, f)
}
