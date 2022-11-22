package _default

import (
	"github.com/aasumitro/posbe/internal/_default/http/handler"
	"github.com/gin-gonic/gin"
)

func InitDefaultModule(router *gin.Engine) {
	handler.NewHttpRouter(router)
}
