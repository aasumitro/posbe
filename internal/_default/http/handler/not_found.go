package handler

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler httpHandler) notFound(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusNotFound, "HTTP_ROUTE_NOT_FOUND")
}
