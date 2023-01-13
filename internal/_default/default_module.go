package _default

import (
	"github.com/aasumitro/posbe/internal/_default/handler/http"
	"github.com/gin-gonic/gin"
)

func InitDefaultModule(router *gin.Engine) {
	http.NewHTTPRouter(router)
}
