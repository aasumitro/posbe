package model

import "database/sql"

const (
	UsersCacheKey = "users"
	RolesCacheKey = "roles"
)

type (
	User struct {
		ID        int           `json:"id"`
		Name      string        `json:"name"`
		Username  string        `json:"username"`
		Email     string        `json:"email"`
		RoleID    int           `json:"-"`
		Password  string        `json:"-"`
		CreatedAt sql.NullInt64 `json:"-"`
		DeletedAt sql.NullInt64 `json:"-"`

		// embed items
		Role Role `json:"role"`

		// only for auth (will be stored/validate via cookie)
		AccessToken  string `json:"-"`
		RefreshToken string `json:"-"`
	}

	Role struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Usage       int    `json:"usage,omitempty"`
	}
)
