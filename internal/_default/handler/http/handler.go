package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type httpHandler struct{}

func NewHttpRouter(router *gin.Engine) {
	handler := &httpHandler{}
	router.NoMethod(handler.noMethod)
	router.NoRoute(handler.notFound)
	router.GET("/", handler.home)
	router.GET("/ping", handler.ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
