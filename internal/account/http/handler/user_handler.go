package handler

import "github.com/aasumitro/posbe/domain"

type userHandler struct {
	accountService domain.IAccountService
}

func (handler userHandler) fetch() {
	//TODO implement me
	panic("implement me")
}

func (handler userHandler) show() {
	//TODO implement me
	panic("implement me")
}

func (handler userHandler) store() {
	//TODO implement me
	panic("implement me")
}

func (handler userHandler) update() {
	//TODO implement me
	panic("implement me")
}

func (handler userHandler) destroy() {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(accountService domain.IAccountService) {
	handler := userHandler{accountService: accountService}
	handler.show()
}
