package http

import (
	"fmt"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type httpHandler struct{}

func (handler httpHandler) home(context *gin.Context) {
	utils.NewHTTPRespond(context, http.StatusOK, map[string]interface{}{
		"01_title":       "POSBE",
		"02_description": "Point of Sales Backend",
		"03_api_spec": fmt.Sprintf(
			"%s://%s/docs/index.html",
			"http",
			context.Request.Host,
		),
		"04_perquisites": map[string]interface{}{
			"01_language":  "https://github.com/golang/go",
			"02_framework": "https://github.com/gin-gonic/gin",
			"03_library": map[string]string{
				"01_swagger": "https://github.com/swaggo/swag",
			},
		},
	})
}

func (handler httpHandler) noMethod(context *gin.Context) {
	utils.NewHTTPRespond(context, http.StatusNotFound, "HTTP_METHOD_NOT_FOUND")
}

func (handler httpHandler) notFound(context *gin.Context) {
	utils.NewHTTPRespond(context, http.StatusNotFound, "HTTP_ROUTE_NOT_FOUND")
}

func (handler httpHandler) ping(context *gin.Context) {
	utils.NewHTTPRespond(context, http.StatusOK, "PONG")
}

func NewHTTPRouter(router *gin.Engine) {
	handler := &httpHandler{}
	router.NoMethod(handler.noMethod)
	router.NoRoute(handler.notFound)
	router.GET("/", handler.home)
	router.GET("/ping", handler.ping)
	router.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(4)))
}
