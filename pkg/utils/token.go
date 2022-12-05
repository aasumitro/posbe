package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type IJSONWebToken interface {
	ClaimJWTToken() (string, error)
}

type JWTClaim struct {
	jwt.RegisteredClaims
	Payload interface{} `json:"payload"`
}

type JSONWebToken struct {
	Issuer    string
	SecretKey []byte
	Payload   interface{}
	IssuedAt  time.Time
	ExpiredAt time.Time
}

// ClaimJWTToken
// args app name, expiration time, secret key, payload
func (j *JSONWebToken) ClaimJWTToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			IssuedAt:  &jwt.NumericDate{Time: j.IssuedAt},
			ExpiresAt: &jwt.NumericDate{Time: j.ExpiredAt},
		},
		Payload: j.Payload,
	})

	return token.SignedString(j.SecretKey)
}
