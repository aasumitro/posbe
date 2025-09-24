package catalog

import (
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service IProductService
}

func (handler productHandler) fetch(ctx *gin.Context) {}

func (handler productHandler) show(ctx *gin.Context) {}

func (handler productHandler) add(ctx *gin.Context) {}

func (handler productHandler) edit(ctx *gin.Context) {}

func (handler productHandler) destroy(ctx *gin.Context) {}

func (handler productHandler) destroyVariant(ctx *gin.Context) {}

func NewProductHandler(service IProductService, router *gin.RouterGroup) {
	handler := productHandler{service: service}
	router.GET("/products", handler.fetch)
	router.GET("/products/:id", handler.show)
	authz := router.Group("/products")
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		authz.POST(utils.EmptyPath, handler.add)
		authz.PATCH("/:id", handler.edit)
		authz.DELETE("/:id", handler.destroy)
		authz.DELETE("/:id/variants/:pid", handler.destroyVariant)
	}
}
