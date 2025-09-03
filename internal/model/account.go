package model

import (
	"database/sql"
	"errors"
	"runtime"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

const (
	UsersCacheKey = "users"
	RolesCacheKey = "roles"

	accessTokenDurationSecond  = 3600  // 1 hrs
	refreshTokenDurationSecond = 28800 // 8 hrs
)

type (
	User struct {
		ID        int64         `json:"id"`
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

var ErrInvalidCredential = errors.New("invalid username or password")

func (user *User) ComparePassword(password string) error {
	valid, err := utils.ComparePassword(runtime.NumCPU(), user.Password, password)
	if err != nil {
		return err
	}
	if !valid {
		return ErrInvalidCredential
	}
	return nil
}

func (user *User) GenerateToken(secret string, includeRefresh bool) error {
	claim := jwt.MapClaims{"id": user.ID, "role_id": user.Role.ID, "role_name": user.Role.Name}

	// generate access token
	accessToken, err := utils.NewJWT(claim, secret, accessTokenDurationSecond)
	if err != nil {
		return err
	}
	user.AccessToken = accessToken

	if includeRefresh {
		// generate refresh token
		refreshToken, err := utils.NewJWT(claim, secret, refreshTokenDurationSecond)
		if err != nil {
			return err
		}
		user.RefreshToken = refreshToken
	}

	return nil
}
