package model

import "database/sql"

type Customer struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Phone       string        `json:"phone"`
	Description string        `json:"description"`
	CreatedAt   sql.NullInt64 `json:"-"`
	UpdatedAt   sql.NullInt64 `json:"-"`
}
