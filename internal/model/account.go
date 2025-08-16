package model

import "database/sql"

const (
	UsersCacheKey = "users"
	RolesCacheKey = "roles"
)

type (
	User struct {
		ID        int           `json:"id"`
		RoleID    int           `json:"role_id,omitempty"`
		Name      string        `json:"name"`
		Username  string        `json:"username"`
		Email     string        `json:"email"`
		Password  string        `json:"-"`
		Role      Role          `json:"role"`
		CreatedAt sql.NullInt64 `json:"created_at"`
		DeletedAt sql.NullInt64 `json:"-"`
	}

	Role struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Usage       int    `json:"usage,omitempty"`
	}
)
