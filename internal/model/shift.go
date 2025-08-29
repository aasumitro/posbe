package model

import "database/sql"

type (
	// Shift is reference section data for store
	Shift struct {
		ID        int           `json:"id"`
		Name      string        `json:"name"`
		StartTime int64         `json:"start_time"`
		EndTime   int64         `json:"end_time"`
		CreatedAt sql.NullInt64 `json:"created_at"`
		UpdatedAt sql.NullInt64 `json:"updated_at,omitempty"`

		// Relation
		ActiveShift    *ActiveShift `json:"active_shift,omitempty" binding:"-"`
		ShiftHistories *ActiveShift `json:"shift_histories,omitempty" binding:"-"`

		// Count data
		TotalUsage       int64 `json:"total_usage" binding:"-"`
		TotalTransaction int64 `json:"total_transaction" binding:"-"`
	}

	ActiveShift struct {
		ID        int           `json:"id"`
		ShiftID   int           `json:"shift_id"`
		OpenAt    sql.NullInt64 `json:"open_at"`
		OpenBy    sql.NullInt64 `json:"open_by"`
		OpenCash  sql.NullInt64 `json:"open_cash"`
		CloseAt   sql.NullInt64 `json:"close_at"`
		CloseBy   sql.NullInt64 `json:"close_by"`
		CloseCash sql.NullInt64 `json:"close_cash"`
		CreatedAt sql.NullInt64 `json:"created_at"`
		UpdatedAt sql.NullInt64 `json:"updated_at,omitempty"`

		// Relation data
		Shift *Shift `json:"shift,omitempty" binding:"-"`
	}
)
