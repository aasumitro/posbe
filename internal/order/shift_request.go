package order

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type ActiveShiftForm struct {
	UserID   int64  `json:"-" form:"-"`
	Action   string `json:"-" form:"-"`
	ShiftID  int64  `json:"-" form:"-"`
	ActiveID int64  `json:"active_id" form:"active_id"`
	Cash     int64  `json:"cash" form:"cash"`
}

func (f *ActiveShiftForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()

	if f.Action == "close" && f.ActiveID == 0 {
		return map[string][]string{"active_id": {"active shift id is required"}}
	}

	return g.ComplexValidator(galidator.Rules{
		"ActiveID": g.R("active_id").Required(),
		"Cash":     g.R("cash").Required(),
	}).Validate(ctx, f)
}
