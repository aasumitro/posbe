package catalog

import (
	"context"
	"github.com/aasumitro/posbe/internal/catalog/handler/http"
	repository "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/aasumitro/posbe/internal/catalog/service"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitCatalogModule(ctx context.Context, router *gin.Engine) {
	unitRepository := repository.NewUnitSQLRepository()
	categoryRepository := repository.NewCategorySQLRepository()
	subcategoryRepository := repository.NewSubcategorySQLRepository()
	addonRepository := repository.NewAddonSQLRepository()
	productRepository := repository.NewProductSQLRepository()
	productVariantRepository := repository.NewProductVariantSQLRepository()
	catalogCommonService := service.NewCatalogCommonService(ctx, unitRepository,
		categoryRepository, subcategoryRepository, addonRepository)
	productCommonService := service.NewCatalogProductService(ctx,
		productRepository, productVariantRepository)
	routerGroup := router.Group("v1")
	protectedRouter := routerGroup.
		Use(middleware.Auth()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewUnitHandler(catalogCommonService, protectedRouter)
	http.NewCategoryHandler(catalogCommonService, protectedRouter)
	http.NewSubcategoryHandler(catalogCommonService, protectedRouter)
	http.NewAddonHandler(catalogCommonService, protectedRouter)
	http.NewProductVariantHandler(productCommonService, protectedRouter)
}
