package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// AcceptedRoles expected tobe needed role
func AcceptedRoles(roles []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		payload := context.MustGet("payload")
		fmt.Println("=== Role Validation")
		fmt.Println(payload.(map[string]interface{})["id"])
		fmt.Println(payload.(map[string]interface{})["role"].(map[string]interface{})["id"])
	}
}
