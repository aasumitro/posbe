package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IJSONWebToken interface {
	ClaimJWTToken(payload interface{}) (string, error)
}

type JWTClaim struct {
	jwt.RegisteredClaims
	Payload interface{} `json:"payload"`
}

type JSONWebToken struct {
	Issuer    string
	SecretKey []byte
	IssuedAt  time.Time
	ExpiredAt time.Time
}

// ClaimJWTToken
// args app name, expiration time, secret key, payload
func (j *JSONWebToken) ClaimJWTToken(payload interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			IssuedAt:  &jwt.NumericDate{Time: j.IssuedAt},
			ExpiresAt: &jwt.NumericDate{Time: j.ExpiredAt},
		},
		Payload: payload,
	})

	return token.SignedString(j.SecretKey)
}
