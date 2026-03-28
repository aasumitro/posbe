package order

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const SSEChannelKey = "orders.events"

func New(router *gin.RouterGroup) {
	shiftRepo := NewShiftRepository(config.PgxPool)
	shiftSvc := NewShiftService(shiftRepo)
	trxRepo := NewTransactionRepository(config.PgxPool)
	trxSvc := NewTransactionService(trxRepo)

	orderRouterGroup := router.Group("orders")

	eventRouterGroup := orderRouterGroup.Group("events")
	{
		eventRouterGroup.GET(utils.EmptyPath, func(ctx *gin.Context) {
			ctx.Writer.Header().Set("Content-Type", "text/event-stream")
			ctx.Writer.Header().Set("Cache-Control", "no-cache")
			ctx.Writer.Header().Set("Connection", "keep-alive")
			ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
			utils.SubscribeEvent(ctx.Request.Context(), SSEChannelKey, func(message *redis.Message) {
				ctx.SSEvent("update", message.Payload)
				ctx.Writer.Flush()
			})
		})
	}

	orderRouterGroup.Use(utils.AuthN())
	{
		NewShiftHandler(shiftSvc, orderRouterGroup)
		NewTransactionHandler(trxSvc, orderRouterGroup)
	}
}
