package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type roomHandler struct {
	svc    domain.IStoreService
	router *gin.RouterGroup
}

func (handler roomHandler) fetch(ctx *gin.Context) {
	rooms, err := handler.svc.RoomList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, rooms)
}

func (handler roomHandler) store(ctx *gin.Context) {
	var form domain.Room
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	room, err := handler.svc.AddRoom(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, room)
}

func (handler roomHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}

	var form domain.Room
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ID = id
	table, err := handler.svc.EditRoom(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, table)
}

func (handler roomHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}
	data := domain.Room{ID: id}

	err := handler.svc.DeleteRoom(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewRoomHandler(svc domain.IStoreService, router *gin.RouterGroup) {
	handler := roomHandler{svc: svc, router: router}
	router.GET("/rooms", handler.fetch)
	router.POST("/rooms", handler.store)
	router.PUT("/rooms/:id", handler.update)
	router.DELETE("/rooms/:id", handler.destroy)
}
