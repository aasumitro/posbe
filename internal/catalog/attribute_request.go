package catalog

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type UnitForm struct {
	ActionAdd bool   `form:"-" json:"-"`
	ID        int64  `form:"-" json:"-"`
	Magnitude string `form:"magnitude" json:"magnitude"`
	Name      string `form:"name" json:"name"`
	Symbol    string `form:"symbol" json:"symbol"`
}

func (f *UnitForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{}
	if f.ActionAdd {
		r = galidator.Rules{
			"Magnitude": g.R("magnitude").Required(),
			"Name":      g.R("name").Required(),
			"Symbol":    g.R("symbol").Required(),
		}
	} else {
		r = galidator.Rules{
			"Magnitude": g.R("magnitude").Optional(),
			"Name":      g.R("name").Optional(),
			"Symbol":    g.R("symbol").Optional(),
		}
	}
	return g.ComplexValidator(r).Validate(ctx, f)
}

type NewCategoryForm struct {
	Name          string   `form:"name" json:"name"`
	Subcategories []string `form:"subcategories" json:"subcategories"`
}

func (f *NewCategoryForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()

	if err := g.ComplexValidator(galidator.Rules{
		"Name": g.R("name").Required().Min(3).Max(20),
	}).Validate(ctx, f); err != nil {
		return err
	}

	if f.Subcategories != nil {
		errs := map[string][]string{}
		for _, sub := range f.Subcategories {
			if len(sub) < 3 {
				errs[sub] = append(errs[sub], "subcategory must be at least 3 characters long")
			}
		}
		if len(errs) > 0 {
			return map[string]interface{}{"subcategories": errs}
		}
	}

	return nil
}

type EditCategoryForm struct {
	ID                int64                  `form:"-" json:"-"`
	Name              string                 `form:"name" json:"name"`
	EditSubcategories []*EditSubcategoryForm `form:"subcategories" json:"subcategories"`
	NewSubcategories  []string               `form:"new_subcategories" json:"new_subcategories"`
}

type EditSubcategoryForm struct {
	ID   int64  `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

func (f *EditCategoryForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()

	if err := g.ComplexValidator(galidator.Rules{
		"Name": g.R("name").Optional().Min(3).Max(20),
	}).Validate(ctx, f); err != nil {
		return err
	}

	if f.EditSubcategories != nil {
		rules := galidator.Rules{
			"ID":   g.R("id").Required(),
			"Name": g.R("name").Required().Min(3).Max(20),
		}
		validator := g.ComplexValidator(rules)

		for _, sub := range f.EditSubcategories {
			if err := validator.Validate(ctx, sub); err != nil {
				return err
			}
		}
	}

	if f.NewSubcategories != nil {
		errs := map[string][]string{}
		for _, sub := range f.NewSubcategories {
			if len(sub) < 3 {
				errs[sub] = append(errs[sub], "subcategory must be at least 3 characters long")
			}
		}
		if len(errs) > 0 {
			return map[string]interface{}{"subcategories": errs}
		}
	}

	return nil
}
