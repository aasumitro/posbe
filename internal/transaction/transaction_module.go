package transaction

import (
	"context"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/gin-gonic/gin"
)

func InitTransactionModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	// Order
	// OrderItem
	// OrderBill
	// Order...
}
