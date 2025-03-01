package http

import (
	"github.com/aasumitro/posbe/internal/model"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	svc model.ICatalogProductService
}

func (handler productHandler) fetch(_ *gin.Context) {}

func (handler productHandler) show(_ *gin.Context) {}

func (handler productHandler) store(_ *gin.Context) {}

func (handler productHandler) update(_ *gin.Context) {}

func (handler productHandler) destroy(_ *gin.Context) {}

func NewProductHandler(svc model.ICatalogProductService, router gin.IRouter) {
	handler := productHandler{svc: svc}
	router.GET("/products", handler.fetch)
	router.POST("/products", handler.store)
	router.GET("/products/:id", handler.show)
	router.PUT("/products/:id", handler.update)
	router.DELETE("/products/:id", handler.destroy)
}
