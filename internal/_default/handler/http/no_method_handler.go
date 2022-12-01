package http

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler httpHandler) noMethod(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusNotFound, "HTTP_METHOD_NOT_FOUND")
}
