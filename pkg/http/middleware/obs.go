package middleware

import (
	"fmt"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/gin-gonic/gin"
)

// ActivityObserver TODO OBSERVE USER ACTIVITY
func ActivityObserver() gin.HandlerFunc {
	return func(context *gin.Context) {
		payload := context.MustGet("payload")
		user := payload.(map[string]interface{})
		_, _ = config.PostgresPool.QueryContext(context,
			"INSERT INTO activity_logs (user_id, description) values ($1, $2)",
			user["id"], fmt.Sprintf("[%s] | [%s]::%s | [%.0f]%s",
				time.Now().Format("2006-01-02 15:04:05"),
				context.Request.Method, context.Request.URL,
				user["id"], user["username"],
			))
	}
}
