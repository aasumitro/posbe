package handler

import (
	"fmt"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler httpHandler) home(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusOK, map[string]interface{}{
		"01_title":       "POSBE",
		"02_description": " Point of Sales Backend",
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
