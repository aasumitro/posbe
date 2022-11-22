package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type httpHandler struct{}

func NewHttpRouter(router *gin.Engine) {
	handler := &httpHandler{}

	fmt.Println(handler)
}
