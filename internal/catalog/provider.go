package catalog

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

//func NewCatalogModuleProvider(router *gin.RouterGroup) {
//	unitRepository := repository.NewUnitSQLRepository()
//	categoryRepository := repository.NewCategorySQLRepository()
//	subcategoryRepository := repository.NewSubcategorySQLRepository()
//	addonRepository := repository.NewAddonSQLRepository()
//	productRepository := repository.NewProductSQLRepository()
//	productVariantRepository := repository.NewProductVariantSQLRepository()
//	catalogCommonService := service.NewCatalogCommonService(unitRepository,
//		categoryRepository, subcategoryRepository, addonRepository)
//	productCommonService := service.NewCatalogProductService(
//		productRepository, productVariantRepository)
//	protectedRouter := router.Use(utils.AuthN())
//	http.NewUnitHandler(catalogCommonService, protectedRouter)
//	http.NewCategoryHandler(catalogCommonService, protectedRouter)
//	http.NewSubcategoryHandler(catalogCommonService, protectedRouter)
//	http.NewAddonHandler(catalogCommonService, protectedRouter)
//	http.NewProductVariantHandler(productCommonService, protectedRouter)
//}

func New(router *gin.RouterGroup) {
	authn := router.Use(utils.AuthN())
	attRepo := NewAttributeRepository(config.PgxPool)
	attSvc := NewAttributeService(attRepo)
	NewAttributeUnitHandler(attSvc, authn)
}
