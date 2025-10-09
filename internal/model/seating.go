package model

import "database/sql"

type (
	Floor struct {
		ID        int           `json:"id"`
		Name      string        `json:"name"`
		CreatedAt sql.NullInt64 `json:"-"`
		UpdatedAt sql.NullInt64 `json:"-"`

		// Embed
		Tables []*Table `json:"tables,omitempty"`

		// counting
		TotalTables int `json:"total_tables,omitempty"`
	}

	//Table Case Study Restaurant Dine in
	Table struct {
		ID        int           `json:"id"`
		FloorID   int           `json:"floor_id"`
		Name      string        `json:"name"`
		XPos      float64       `json:"x_pos"`
		YPos      float64       `json:"y_pos"`
		WSize     float64       `json:"w_size"`
		HSize     float64       `json:"h_size"`
		DSize     float64       `json:"d_size"`
		Capacity  int           `json:"capacity"`
		Type      string        `json:"type"`
		CreatedAt sql.NullInt64 `json:"-"`
		UpdatedAt sql.NullInt64 `json:"-"`
	}
)
