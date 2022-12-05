package middleware

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// Auth expected tobe logged in
func Auth(jwtSignature string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenCookie, err := context.Request.Cookie("jwt")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		if token, err := jwt.ParseWithClaims(
			tokenCookie.Value,
			&utils.JWTClaim{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSignature), nil
			},
		); err != nil &&!token.Valid {
			utils.NewHttpRespond(context, http.StatusUnauthorized, "TOKEN_NOT_VALID")
			return
		} else {
			claims, ok := token.Claims.(*utils.JWTClaim)
			if !ok {
				utils.NewHttpRespond(context, http.StatusUnauthorized, "TOKEN_NOT_VALID")
				context.Abort()
				return
			}
			context.Set("payload", claims.Payload)
		}

		context.Next()
	}
}
