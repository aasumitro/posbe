package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type floorHandler struct {
	svc    domain.IStoreService
	router *gin.RouterGroup
}

func (handler floorHandler) floorsWithTables(ctx *gin.Context) {
	floors, err := handler.svc.FloorsWith(domain.Table{})
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floors)
}

func (handler floorHandler) floorsWithRooms(ctx *gin.Context) {
	floors, err := handler.svc.FloorsWith(domain.Room{})
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floors)
}

func (handler floorHandler) fetch(ctx *gin.Context) {
	floors, err := handler.svc.FloorList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floors)
}

func (handler floorHandler) store(ctx *gin.Context) {
	var form domain.Floor
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	floor, err := handler.svc.AddFloor(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, floor)
}

func (handler floorHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}

	var form domain.Floor
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ID = id
	floor, err := handler.svc.EditFloor(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floor)
}

func (handler floorHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}
	data := domain.Floor{ID: id}

	err := handler.svc.DeleteFloor(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewFloorHandler(svc domain.IStoreService, router *gin.RouterGroup) {
	handler := floorHandler{svc: svc, router: router}
	router.GET("/floors/tables", handler.floorsWithTables)
	router.GET("/floors/rooms", handler.floorsWithRooms)
	router.GET("/floors", handler.fetch)
	router.POST("/floors", handler.store)
	router.PUT("/floors/:id", handler.update)
	router.DELETE("/floors/:id", handler.destroy)
}
