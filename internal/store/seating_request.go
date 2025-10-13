package store

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Status    string   `form:"status" json:"status"`
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
	r["Status"] = g.R("status").Optional().Choices("available",
		"occupied", "reserved", "disabled", "ordering", "billed")
	return g.ComplexValidator(r).Validate(ctx, f)
}

func (f *SeatingTableRequest) Bind(data string) (*SeatingTableRequest, error) {
	// Expected sample data:
	//      {"id": 12, "action": "update", "field": "status", "value": "available"}

	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return nil, fmt.Errorf("error unmarshaling payload: %w", err)
	}

	if len(payload) == 0 {
		return nil, errors.New("payload is empty")
	}

	action, ok := payload["action"].(string)
	if !ok || action == "" {
		return nil, errors.New("invalid or missing action field")
	}
	f.ActionAdd = strings.ToLower(action) != "update"

	pid, ok := payload["id"].(float64)
	if !ok || pid == 0 {
		return nil, errors.New("invalid or missing id field")
	}
	f.ID = int64(pid)

	field, ok := payload["field"].(string)
	if !ok || field == "" {
		return nil, errors.New("invalid or missing field field")
	}

	value, exists := payload["value"]
	if !exists {
		return nil, errors.New("missing value field")
	}

	switch strings.ToLower(field) {
	case "status":
		if strVal, ok := value.(string); ok && strVal != "" {
			f.Status = strVal
		} else {
			return nil, fmt.Errorf("invalid value for field 'status': %v", value)
		}
		// optionally support other future updates here
		// e.g:
		// case "capacity":
		//		switch v := value.(type) {
		//		case float64:
		//			f.Capacity = int64(v)
		//		case int:
		//			f.Capacity = int64(v)
		//		case string:
		//			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
		//				f.Capacity = parsed
		//			} else {
		//				return nil, fmt.Errorf("invalid numeric string for capacity: %v", v)
		//			}
		//		default:
		//			return nil, fmt.Errorf("unsupported type for capacity: %T", v)
		//		}
	}

	return f, nil
}
