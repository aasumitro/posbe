package catalog

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

func New(router *gin.RouterGroup) {
	authn := router.Use(utils.AuthN())
	attRepo := NewAttributeRepository(config.PgxPool)
	attSvc := NewAttributeService(attRepo)
	NewAttributeUnitHandler(attSvc, authn)
	NewAttributeCategoryHandler(attSvc, authn)
	prdRepo := NewProductRepository(config.PgxPool)
	prdSvc := NewProductService(prdRepo)
	NewProductAddonHandler(prdSvc, authn)
}
