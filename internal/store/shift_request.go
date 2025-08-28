package store

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type ShiftForm struct {
	ID        int64  `form:"-" json:"-"`
	ActionAdd bool   `form:"-" json:"-"`
	Name      string `json:"name" form:"name"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

func (f *ShiftForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{}
	if f.ActionAdd {
		r = galidator.Rules{
			"Name":      g.R("name").Required(),
			"StartTime": g.R("start_time").Required(),
			"EndTime":   g.R("end_time").Required(),
		}
	} else {
		r = galidator.Rules{
			"Name":      g.R("name").Optional(),
			"StartTime": g.R("start_time").Optional(),
			"EndTime":   g.R("end_time").Optional(),
		}
	}

	if err := g.ComplexValidator(r).Validate(ctx, f); err != nil {
		return err
	}

	// custom validation: end_time must not be less than start_time
	if f.StartTime != 0 && f.EndTime != 0 && f.EndTime <= f.StartTime {
		return map[string][]string{
			"end_time": {"End time must be greater than start time"},
		}
	}

	return nil
}
