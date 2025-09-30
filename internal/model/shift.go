package model

import "database/sql"

const (
	StoreShiftCacheKey = "store_shifts"
)

type (
	// Shift is reference section data for store
	Shift struct {
		ID           int           `json:"id"`
		Name         string        `json:"name"`
		StartTime    int64         `json:"start_time"`
		PrevDayStart bool          `json:"prev_day_start"`
		EndTime      int64         `json:"end_time"`
		NextDayEnd   bool          `json:"next_day_end"`
		CreatedAt    sql.NullInt64 `json:"created_at"`
		UpdatedAt    sql.NullInt64 `json:"updated_at,omitempty"`

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

func (s *Shift) ApplyCountingData() {
	if s == nil || len(s.Histories) == 0 {
		return
	}

	userSet := make(map[int64]struct{})

	for _, h := range s.Histories {
		if !h.CloseAt.Valid || !h.CloseCash.Valid {
			if s.Active == nil {
				s.Active = h
			}
		} else if s.Last == nil {
			s.Last = h
		}

		s.TotalUsage++

		if h.OpenCash.Valid && h.CloseCash.Valid {
			diff := h.CloseCash.Int64 - h.OpenCash.Int64
			if diff >= 0 {
				s.TotalSurplus++
				s.Profit += diff
			} else {
				s.TotalDeficit++
				s.Loss += -diff
			}
		}

		if h.OpenBy.Valid {
			userSet[h.OpenBy.Int64] = struct{}{}
		}
		if h.CloseBy.Valid {
			userSet[h.CloseBy.Int64] = struct{}{}
		}
	}

	for uid := range userSet {
		val := uid
		s.UsersID = append(s.UsersID, &val)
	}

	s.Net = s.Profit + s.Loss

	s.Histories = nil
}
