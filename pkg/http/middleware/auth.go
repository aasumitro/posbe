package middleware

import (
	"github.com/aasumitro/posbe/configs"
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

		token, err := jwt.ParseWithClaims(
			tokenCookie.Value,
			&utils.JWTClaim{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(configs.Cfg.JWTSecretKey), nil
			})
		if err != nil && !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		claims, ok := token.Claims.(*utils.JWTClaim)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.Set("payload", claims.Payload)
		context.Next()
	}
}
