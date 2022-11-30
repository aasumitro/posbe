package handler

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler httpHandler) ping(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusOK, "PONG")
}
