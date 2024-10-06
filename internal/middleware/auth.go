package middleware

import (
	"log"
	"net/http"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
			func(_ *jwt.Token) (interface{}, error) {
				return []byte(config.Instance.JWTSecretKey), nil
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
		// extract payload
		payload := claims.Payload
		context.Set("payload", payload)
		user, ok := payload.(map[string]interface{})
		if !ok {
			log.Println("failed to extract user data from payload")
			context.Next()
			return
		}
		role, ok := user["role"].(map[string]interface{})
		if !ok {
			log.Println("failed extract user role from payload")
			context.Next()
			return
		}
		context.Set("user_id", user["id"])
		context.Set("role_id", role["id"])
		context.Set("role_name", role["name"])
		context.Next()
	}
}
