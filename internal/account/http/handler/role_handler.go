package handler

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	svc    domain.IAccountService
	router *gin.Engine
}

func (handler roleHandler) fetch() {
	//TODO implement me
	panic("implement me")
}

func (handler roleHandler) show() {
	//TODO implement me
	panic("implement me")
}

func (handler roleHandler) store() {
	//TODO implement me
	panic("implement me")
}

func (handler roleHandler) update() {
	//TODO implement me
	panic("implement me")
}

func (handler roleHandler) destroy() {
	//TODO implement me
	panic("implement me")
}

func NewRoleHandler(accountService domain.IAccountService, router *gin.Engine) {
	handler := roleHandler{svc: accountService, router: router}
	handler.show()
}
