package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ParseJWT verifies and parses JWT and returns its claims.
func ParseJWT(token, verificationKey string) (jwt.MapClaims, error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"}))
	parsedToken, err := parser.Parse(token, func(_ *jwt.Token) (any, error) {
		return []byte(verificationKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, errors.New("unable to parse token")
}

// NewJWT generates and returns new HS256 signed JWT.
func NewJWT(payload jwt.MapClaims, signingKey string, secondsDuration int64) (string, error) {
	seconds := time.Duration(secondsDuration) * time.Second
	claims := jwt.MapClaims{"exp": time.Now().Add(seconds).Unix()}
	for k, v := range payload {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signingKey))
}
