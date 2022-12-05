package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaim struct {
	jwt.RegisteredClaims
	Payload interface{} `json:"payload"`
}

// ClaimJWTToken
// args app name, expiration time, secret key, payload
func ClaimJWTToken(args ...any) (string, error) {
	if len(args) < 4 {
		return "", errors.New("args is less than 4, needed args: app_name, exp_time, payload, secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			// ISSUER/APP_NAME args[0]
			Issuer:   args[0].(string),
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
			// EXPIRATION TIME args[1]
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(args[1].(int)) * time.Hour)},
		},
		// PAYLOAD USER args[2]
		Payload: args[2],
	})

	// SECRET/SIGNATURE args[3]
	return token.SignedString(args[3])
}

