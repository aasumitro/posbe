package middleware

import (
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// Auth expected tobe logged in
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenCookie, err := context.Request.Cookie("jwt")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		if token, err := jwt.ParseWithClaims(
			tokenCookie.Value,
			&utils.JWTClaim{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(config.Cfg.JWTSecretKey), nil
			},
		); err != nil && !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		} else {
			claims, ok := token.Claims.(*utils.JWTClaim)
			if !ok {
				context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
				return
			}
			context.Set("payload", claims.Payload)
		}

		context.Next()
	}
}
