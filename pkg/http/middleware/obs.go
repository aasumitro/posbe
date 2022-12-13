package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// ActivityObserver TODO OBSERVE USER ACTIVITY
func ActivityObserver() gin.HandlerFunc {
	return func(context *gin.Context) {
		payload := context.MustGet("payload")
		fmt.Println("=== Observer")
		fmt.Println(context.Request.URL)
		fmt.Println(payload.(map[string]interface{})["id"])
		fmt.Println(payload.(map[string]interface{})["username"])
	}
}
