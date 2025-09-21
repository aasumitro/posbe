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
