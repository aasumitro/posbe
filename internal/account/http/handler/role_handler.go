package handler

import "github.com/aasumitro/posbe/domain"

type roleHandler struct {
	accountService domain.AccountService
}

func (h roleHandler) fetch() {
	//TODO implement me
	panic("implement me")
}

func (h roleHandler) show() {
	//TODO implement me
	panic("implement me")
}

func (h roleHandler) store() {
	//TODO implement me
	panic("implement me")
}

func (h roleHandler) update() {
	//TODO implement me
	panic("implement me")
}

func (h roleHandler) destroy() {
	//TODO implement me
	panic("implement me")
}

func NewRoleHandler(accountService domain.AccountService) {
	handler := roleHandler{accountService: accountService}
	handler.show()
}
