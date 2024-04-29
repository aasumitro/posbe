package catalog

import (
	"github.com/aasumitro/posbe/internal/catalog/handler/http"
	repository "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/aasumitro/posbe/internal/catalog/service"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewCatalogModuleProvider(router *gin.RouterGroup) {
	unitRepository := repository.NewUnitSQLRepository()
	categoryRepository := repository.NewCategorySQLRepository()
	subcategoryRepository := repository.NewSubcategorySQLRepository()
	addonRepository := repository.NewAddonSQLRepository()
	productRepository := repository.NewProductSQLRepository()
	productVariantRepository := repository.NewProductVariantSQLRepository()
	catalogCommonService := service.NewCatalogCommonService(unitRepository,
		categoryRepository, subcategoryRepository, addonRepository)
	productCommonService := service.NewCatalogProductService(
		productRepository, productVariantRepository)
	protectedRouter := router.
		Use(middleware.Auth()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewUnitHandler(catalogCommonService, protectedRouter)
	http.NewCategoryHandler(catalogCommonService, protectedRouter)
	http.NewSubcategoryHandler(catalogCommonService, protectedRouter)
	http.NewAddonHandler(catalogCommonService, protectedRouter)
	http.NewProductVariantHandler(productCommonService, protectedRouter)
}
