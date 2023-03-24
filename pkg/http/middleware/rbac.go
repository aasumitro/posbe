package middleware

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AcceptedRoles expected tobe needed role
func AcceptedRoles(accepted []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(accepted) > 0 && accepted[0] != "*" {
			payload := context.MustGet("payload")
			role := payload.(map[string]interface{})["role"]
			name := role.(map[string]interface{})["name"]
			if !utils.InArray(name.(string), accepted) {
				context.AbortWithStatusJSON(http.StatusUnauthorized, "USER_NOT_AUTHORIZED")
				return
			}
		}

		context.Next()
	}
}
