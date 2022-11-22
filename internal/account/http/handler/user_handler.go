package handler

import "github.com/aasumitro/posbe/domain"

type userHandler struct {
	accountService domain.AccountService
}

func (h userHandler) fetch() {
	//TODO implement me
	panic("implement me")
}

func (h userHandler) show() {
	//TODO implement me
	panic("implement me")
}

func (h userHandler) store() {
	//TODO implement me
	panic("implement me")
}

func (h userHandler) update() {
	//TODO implement me
	panic("implement me")
}

func (h userHandler) destroy() {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(accountService domain.AccountService) {
	handler := roleHandler{accountService: accountService}
	handler.show()
}
