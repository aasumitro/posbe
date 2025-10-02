package order

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

const (
	ShiftActionOpen  = "open"
	ShiftActionClose = "close"
)

type ActiveShiftForm struct {
	UserID        int64   `json:"-" form:"-"`
	Action        string  `json:"action" form:"action"`
	ShiftID       int64   `json:"shift_id" form:"shift_id"`
	ActiveShiftID int64   `json:"active_shift_id" form:"active_shift_id"`
	Cash          float64 `json:"cash" form:"cash"`
}

func (f *ActiveShiftForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()

	f.Action = strings.ToLower(strings.TrimSpace(f.Action))

	if f.Action == ShiftActionClose && f.ActiveShiftID == 0 {
		return map[string][]string{"active_shift_id": {"active shift id is required"}}
	}

	return g.ComplexValidator(galidator.Rules{
		"Action": g.R("action").Required().
			Choices(ShiftActionOpen, ShiftActionClose),
		"ShiftID":       g.R("shift_id").Required(),
		"ActiveShiftID": g.R("active_shift_id").Optional(),
		"Cash":          g.R("cash").Required(),
	}).Validate(ctx, f)
}
