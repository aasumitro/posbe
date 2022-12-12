package catalog

import (
	"context"
	"github.com/aasumitro/posbe/internal/catalog/handler/http"
	repository "github.com/aasumitro/posbe/internal/catalog/repository/sql"
	"github.com/aasumitro/posbe/internal/catalog/service"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitCatalogModule(ctx context.Context, router *gin.Engine) {
	unitRepository := repository.NewUnitSQLRepository(config.Db)
	categoryRepository := repository.NewCategorySQLRepository(config.Db)
	subcategoryRepository := repository.NewSubcategorySQLRepository(config.Db)
	addonRepository := repository.NewAddonSQLRepository(config.Db)
	catalogCommonService := service.NewCatalogCommonService(ctx, unitRepository,
		categoryRepository, subcategoryRepository, addonRepository)
	routerGroup := router.Group("v1")
	protectedRouter := routerGroup.
		Use(middleware.Auth()).
		Use(middleware.AcceptedRoles([]string{"*"}))
	http.NewUnitHandler(catalogCommonService, protectedRouter)
	http.NewCategoryHandler(catalogCommonService, protectedRouter)
	http.NewSubcategoryHandler(catalogCommonService, protectedRouter)
	http.NewAddonHandler(catalogCommonService, protectedRouter)
}
