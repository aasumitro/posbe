package order

import (
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service ITransactionService
}

func (handler transactionHandler) fetch(ctx *gin.Context) {
	// TODO:
}

func (handler transactionHandler) show(ctx *gin.Context) {
	// TODO:
}

func (handler transactionHandler) add(ctx *gin.Context) {
	// TODO:
}

func NewTransactionHandler(service ITransactionService, router *gin.RouterGroup) {
	handler := transactionHandler{service: service}
	router.GET(utils.EmptyPath, handler.fetch)
	router.GET("/:id", handler.show)
	router.POST(utils.EmptyPath, handler.add)
}
