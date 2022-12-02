package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type roleHandler struct {
	svc domain.IAccountService
}

func (handler roleHandler) fetch(ctx *gin.Context) {
	roles, err := handler.svc.RoleList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, roles)
}

func (handler roleHandler) store(ctx *gin.Context) {
	var form domain.Role
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	role, err := handler.svc.AddRole(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, role)
}

func (handler roleHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}

	var form domain.Role
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ID = id
	role, err := handler.svc.EditRole(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, role)
}

func (handler roleHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}
	data := domain.Role{ID: id}

	err := handler.svc.DeleteRole(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewRoleHandler(accountService domain.IAccountService, router *gin.RouterGroup) {
	handler := roleHandler{svc: accountService}
	router.GET("/roles", handler.fetch)
	router.POST("/roles", handler.store)
	router.PUT("/roles/:id", handler.update)
	router.DELETE("/roles/:id", handler.destroy)
}
