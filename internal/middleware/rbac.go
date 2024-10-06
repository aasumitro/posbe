package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// AcceptedRoles expected to be a necessary role
func AcceptedRoles(accepted []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(accepted) > 0 && accepted[0] != "*" {
			name, ok := context.Get("role_name")
			if !ok {
				context.AbortWithStatusJSON(http.StatusUnauthorized,
					"USER_NOT_AUTHORIZED")
				return
			}
			if !slices.Contains(accepted, name.(string)) {
				context.AbortWithStatusJSON(http.StatusUnauthorized,
					"USER_NOT_AUTHORIZED")
				return
			}
		}
		context.Next()
	}
}
