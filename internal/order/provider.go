package order

import (
	"encoding/json"
	"net/http"

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
		eventRouterGroup.POST(utils.EmptyPath, func(ctx *gin.Context) {
			var form struct {
				Type string `json:"type" form:"type"`
			}

			// bind user input
			if err := ctx.ShouldBind(&form); err != nil {
				utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
				return
			}

			if form.Type == "" {
				utils.NewHTTPRespond(ctx, http.StatusBadRequest, "Type is required")
				return
			}

			var payload interface{}
			switch form.Type {
			case "T1":
				payload = map[string]interface{}{
					"id": 1, "type": "table", "field": []string{"status", "customers"},
					"status": "occupied", "customers": 3,
				}
			case "T3":
				payload = map[string]interface{}{
					"id": 3, "type": "table", "field": []string{"status", "customers"},
					"status": "reserved", "customers": 1,
				}
			case "RELOAD":
				payload = map[string]interface{}{"type": "reload"}
			default:
				utils.NewHTTPRespond(ctx, http.StatusBadRequest, "Unknown type")
				return
			}

			dataBytes, err := json.Marshal(payload)
			if err != nil {
				utils.NewHTTPRespond(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			data := string(dataBytes)

			if err := utils.PublishEvent(ctx, SSEChannelKey, data); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.JSON(http.StatusCreated, "Notify sent . . .")
		})

		eventRouterGroup.GET(utils.EmptyPath, func(ctx *gin.Context) {
			ctx.Writer.Header().Set("Content-Type", "text/event-stream")
			ctx.Writer.Header().Set("Cache-Control", "no-cache")
			ctx.Writer.Header().Set("Connection", "keep-alive")
			ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
			utils.SubscribeEvent(ctx.Request.Context(), SSEChannelKey, func(message *redis.Message) {
				// status: "available" | "occupied" | "reserved" | "needs-cleaning"
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

func initTableStatus() {

}
