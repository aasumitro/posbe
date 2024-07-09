package http

import (
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	svc model.ICatalogProductService
}

func (handler productHandler) fetch(ctx *gin.Context) {}

func (handler productHandler) show(ctx *gin.Context) {}

func (handler productHandler) store(ctx *gin.Context) {}

func (handler productHandler) update(ctx *gin.Context) {}

func (handler productHandler) destroy(ctx *gin.Context) {}

func NewProductHandler(svc model.ICatalogProductService, router gin.IRouter) {
	handler := productHandler{svc: svc}
	router.GET("/products", handler.fetch)
	router.POST("/products", handler.store)
	router.GET("/products/:id", handler.show)
	router.PUT("/products/:id", handler.update)
	router.DELETE("/products/:id", handler.destroy)
}
