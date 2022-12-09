package middleware

import (
	"encoding/json"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AcceptedRoles expected tobe needed role
func AcceptedRoles(accepted []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(accepted) > 0 && accepted[0] != "*" {
			var userRole domain.Role
			payload := context.MustGet("payload")
			role := payload.(map[string]interface{})["role"]
			data, _ := json.Marshal(role)
			_ = json.Unmarshal(data, &userRole)
			if !utils.InArray(userRole.Name, accepted) {
				context.AbortWithStatusJSON(http.StatusUnauthorized, "USER_NOT_AUTHORIZED")
				return
			}
		}

		context.Next()
	}
}
