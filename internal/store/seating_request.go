package store

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type SeatingFloorRequest struct {
	ID   int64  `form:"-" json:"-"`
	Name string `form:"name" json:"name"`
}

func (f *SeatingFloorRequest) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		"Name": g.R("name").Required(),
	}).Validate(ctx, f)
}

type SeatingTableRequest struct {
	ActionAdd bool     `form:"-" json:"-"`
	ID        int64    `form:"-" json:"-"`
	FloorID   int64    `form:"floor_id" json:"floor_id"`
	Name      string   `form:"name" json:"name"`
	XPos      *float64 `form:"x_pos" json:"x_pos"`
	YPos      *float64 `form:"y_pos" json:"y_pos"`
	WSize     *float64 `form:"w_size" json:"w_size"`
	HSize     *float64 `form:"h_size" json:"h_size"`
	DSize     *float64 `form:"d_size" json:"d_size"`
	Capacity  int64    `form:"capacity" json:"capacity"`
	Type      string   `form:"type" json:"type"`
}

func (f *SeatingTableRequest) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{}
	if f.ActionAdd {
		r = galidator.Rules{
			"FloorID":  g.R("floor_id").Required(),
			"Name":     g.R("name").Required(),
			"XPos":     g.R("x_pos").Required(),
			"YPos":     g.R("y_pos").Required(),
			"Capacity": g.R("capacity").Required(),
			"Type": g.R("type").Required().Choices(
				"rectangle", "circle"),
		}
		switch strings.ToLower(f.Type) {
		case "rectangle":
			r["WSize"] = g.R("w_size").Required()
			r["HSize"] = g.R("h_size").Required()
		case "circle":
			r["DSize"] = g.R("d_size").Required()
		}
	} else {
		r = galidator.Rules{
			"FloorID":  g.R("floor_id").Optional(),
			"Name":     g.R("name").Optional(),
			"XPos":     g.R("x_pos").Optional(),
			"YPos":     g.R("y_pos").Optional(),
			"Capacity": g.R("capacity").Optional(),
			"Type":     g.R("type").Optional(),
			"WSize":    g.R("w_size").Optional(),
			"HSize":    g.R("h_size").Optional(),
			"DSize":    g.R("d_size").Optional(),
		}
	}
	return g.ComplexValidator(r).Validate(ctx, f)
}
