package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// AcceptedRoles expected to be needed role
func AcceptedRoles(accepted []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(accepted) > 0 && accepted[0] != "*" {
			payload := context.MustGet("payload")
			role := payload.(map[string]interface{})["role"]
			name := role.(map[string]interface{})["name"]
			if !slices.Contains(accepted, name.(string)) {
				context.AbortWithStatusJSON(http.StatusUnauthorized,
					"USER_NOT_AUTHORIZED")
				return
			}
		}
		context.Next()
	}
}
