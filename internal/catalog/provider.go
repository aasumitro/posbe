package catalog

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

func New(router *gin.RouterGroup) {
	attRepo := NewAttributeRepository(config.PgxPool)
	attSvc := NewAttributeService(attRepo)
	prdRepo := NewProductRepository(config.PgxPool)
	prdSvc := NewProductService(prdRepo)
	authn := router.Group(utils.EmptyPath)
	authn.Use(utils.AuthN())
	{
		NewAttributeUnitHandler(attSvc, authn)
		NewAttributeCategoryHandler(attSvc, authn)
		NewProductAddonHandler(prdSvc, authn)
		NewProductHandler(prdSvc, authn)
	}
}
