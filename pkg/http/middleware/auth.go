package middleware

import "github.com/gin-gonic/gin"

// Auth expected tobe logged in
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// TODO
	}
}
