package handler

import (
	"context"
	"github.com/aasumitro/posbe/domain"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	svc    domain.IAccountService
	router *gin.Engine
}

func (handler authHandler) check() {
	// TODO
}

func (handler authHandler) verify() {
	// TODO
}

func (handler authHandler) logout() {
	// TODO
}

func NewAuthHandler(ctx context.Context, accountService domain.IAccountService, router *gin.Engine) {
	handler := authHandler{svc: accountService, router: router}
	handler.logout()
}
