package utils

import (
	"net/http"
	"slices"

	"github.com/aasumitro/posbe/config"
	"github.com/gin-gonic/gin"
)

// AuthN expected tobe logged in
func AuthN() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenCookie, err := ctx.Request.Cookie("access_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		claim, err := ParseJWT(tokenCookie.Value, config.Instance.JWTSecretKey)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if userIDClaim, ok := claim["id"].(float64); ok {
			ctx.Set("user_id", userIDClaim)
		}

		if userRoleIDClaim, ok := claim["role_id"].(float64); ok {
			ctx.Set("role_id", userRoleIDClaim)
		}

		if userRoleNameClaim, ok := claim["role_name"].(string); ok {
			ctx.Set("role_name", userRoleNameClaim)
		}

		ctx.Next()
	}
}

// AuthZ expected to be a necessary role
func AuthZ(accepted []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(accepted) > 0 && accepted[0] != "*" {
			name, ok := ctx.Get("role_name")
			if !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized,
					"USER_NOT_AUTHORIZED")
				return
			}
			if !slices.Contains(accepted, name.(string)) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized,
					"USER_NOT_AUTHORIZED")
				return
			}
		}
		ctx.Next()
	}
}
