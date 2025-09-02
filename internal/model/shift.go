package model

import "database/sql"

const (
	StoreShiftCacheKey = "store_shifts"
)

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
		Active    *ActiveShift   `json:"active,omitempty"`
		Last      *ActiveShift   `json:"last,omitempty"` // so if active is empty (find last)
		Histories []*ActiveShift `json:"histories,omitempty"`
		Orders    []*Order       `json:"orders,omitempty"`
		UsersID   []*int64       `json:"users_id,omitempty"`

		// Count data
		TotalUsage       int64 `json:"total_usage,omitempty"`
		TotalTransaction int64 `json:"total_transaction,omitempty"`
		TotalSurplus     int64 `json:"total_surplus,omitempty"`
		TotalDeficit     int64 `json:"total_deficit,omitempty"`
		Profit           int64 `json:"profit,omitempty"`
		Loss             int64 `json:"loss,omitempty"`
		Net              int64 `json:"net,omitempty"`
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
