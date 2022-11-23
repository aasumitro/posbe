package catalog

import (
	"context"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/gin-gonic/gin"
)

func InitCatalogModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	// Product
	// --- variant
	// --- unit
	// --- add on
	// Category
	// Subcategory
}
