package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type RequestQuery struct{}

func (q *RequestQuery) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		// TODO:
	}).Validate(ctx, q)
}

type RequestForm struct {
	ActionAdd   bool   `form:"-" json:"-"`
	ID          int64  `form:"-" json:"-"`
	Name        string `form:"name" json:"name"`
	Phone       string `form:"phone" json:"phone"`
	Email       string `form:"email" json:"email"`
	Description string `form:"description" json:"description"`
}

func (f *RequestForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{}
	if f.ActionAdd {
		r = galidator.Rules{"Name": g.R("name").Required()}
	} else {
		r = galidator.Rules{"Name": g.R("name").Optional()}
	}
	r["Phone"] = g.R("phone").Optional().Phone()
	r["Email"] = g.R("email").Optional().Email()
	r["Description"] = g.R("description").Optional()
	return g.ComplexValidator(r).Validate(ctx, f)
}
