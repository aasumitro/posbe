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

// floors godoc
// @Schemes
// @Summary 	 Floor List With Tables
// @Description  Get Floors List With Tables.
// @Tags 		 Floors
// @Accept       json
// @Produce      json
// @Success 200 {array} domain.Floor "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/floors/tables [GET]
func (handler floorHandler) floorsWithTables(ctx *gin.Context) {
	floors, err := handler.svc.FloorsWith(domain.Table{})
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floors)
}

// floors godoc
// @Schemes
// @Summary 	 Floor List With Rooms
// @Description  Get Floors List With Rooms.
// @Tags 		 Floors
// @Accept       json
// @Produce      json
// @Success 200 {array} domain.Floor "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/floors/rooms [GET]
func (handler floorHandler) floorsWithRooms(ctx *gin.Context) {
	floors, err := handler.svc.FloorsWith(domain.Room{})
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floors)
}

// floors godoc
// @Schemes
// @Summary 	 Floor List
// @Description  Get Floors List.
// @Tags 		 Floors
// @Accept       json
// @Produce      json
// @Success 200 {array} domain.Floor "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/floors [GET]
func (handler floorHandler) fetch(ctx *gin.Context) {
	floors, err := handler.svc.FloorList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, floors)
}

// floors godoc
// @Schemes
// @Summary 	 Store Floor Data
// @Description  Create new Floor.
// @Tags 		 Floors
// @Accept       mpfd
// @Produce      json
// @Param name 	formData string true "name"
// @Success 201 {object} domain.Floor "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/floors [POST]
func (handler floorHandler) store(ctx *gin.Context) {
	var form domain.Floor
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	floor, err := handler.svc.AddFloor(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, floor)
}

// floors godoc
// @Schemes
// @Summary 	 Update Floor Data
// @Description  Update Floor Data by ID.
// @Tags 		 Floors
// @Accept       mpfd
// @Produce      json
// @Param id   			path     int  	true "floor id"
// @Param name 			formData string true "name"
// @Success 200 {object} domain.Floor "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/floors/{id} [PUT]
func (handler floorHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Floor
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
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

// floors godoc
// @Schemes
// @Summary 	 Delete Floor Data
// @Description  Delete Floor Data by ID.
// @Tags 		 Floors
// @Accept       json
// @Produce      json
// @Param id   			path     int  	true "floor id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/floors/{id} [DELETE]
func (handler floorHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
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
