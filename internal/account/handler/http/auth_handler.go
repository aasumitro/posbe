package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	svc    domain.IAccountService
	router *gin.RouterGroup
}

func (handler AuthHandler) login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "login route"})
}

func (handler AuthHandler) logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "logout route"})
}

func NewAuthHandler(accountService domain.IAccountService, router *gin.RouterGroup) {
	handler := AuthHandler{svc: accountService, router: router}
	router.POST("/login", handler.login)
	router.POST("/logout", handler.logout)
}
