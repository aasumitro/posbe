package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type tableHandler struct {
	svc    domain.IStoreService
	router *gin.RouterGroup
}

func (handler tableHandler) fetch(ctx *gin.Context) {
	tables, err := handler.svc.TableList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, tables)
}

func (handler tableHandler) store(ctx *gin.Context) {
	var form domain.Table
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	table, err := handler.svc.AddTable(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, table)
}

func (handler tableHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}

	var form domain.Table
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ID = id
	table, err := handler.svc.EditTable(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, table)
}

func (handler tableHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}
	data := domain.Table{ID: id}

	err := handler.svc.DeleteTable(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewTableHandler(svc domain.IStoreService, router *gin.RouterGroup) {
	handler := tableHandler{svc: svc, router: router}
	router.GET("/tables", handler.fetch)
	router.POST("/tables", handler.store)
	router.PUT("/tables/:id", handler.update)
	router.DELETE("/tables/:id", handler.destroy)
}
