package http

import (
	"net/http"
	"strconv"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type roomHandler struct {
	svc model.IStoreService
}

// rooms godoc
// @Schemes
// @Summary 	 Room List
// @Description  Get Room List.
// @Tags 		 Rooms
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Room} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms [GET]
func (handler roomHandler) fetch(ctx *gin.Context) {
	rooms, err := handler.svc.RoomList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, rooms)
}

// rooms godoc
// @Schemes
// @Summary 	 Store Room Data
// @Description  Create new Room.
// @Tags 		 Rooms
// @Accept       mpfd
// @Produce      json
// @Param floor_id 	formData string true "floor id"
// @Param name 		formData string true "name"
// @Param x_pos 	formData string true "x position"
// @Param y_pos 	formData string true "y position"
// @Param w_size 	formData string true "weight"
// @Param h_size 	formData string true "height"
// @Param capacity 	formData string true "capacity"
// @Param price 	formData string true "price"
// @Success 201 {object} utils.SuccessRespond{data=model.Room} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms [POST]
func (handler roomHandler) store(ctx *gin.Context) {
	var form model.Room
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	room, err := handler.svc.AddRoom(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, room)
}

// rooms godoc
// @Schemes
// @Summary 	 Update Room Data
// @Description  Update Room Data by ID.
// @Tags 		 Rooms
// @Accept       mpfd
// @Produce      json
// @Param id   		path     int  	true "room id"
// @Param floor_id 	formData string true "floor id"
// @Param name 		formData string true "name"
// @Param x_pos 	formData string true "x position"
// @Param y_pos 	formData string true "y position"
// @Param w_size 	formData string true "weight"
// @Param h_size 	formData string true "height"
// @Param capacity 	formData string true "capacity"
// @Param price 	formData string true "price"
// @Success 200 {object} utils.SuccessRespond{data=model.Room} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms/{id} [PUT]
func (handler roomHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form model.Room
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	form.ID = id
	room, err := handler.svc.EditRoom(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, room)
}

// rooms godoc
// @Schemes
// @Summary 	 Delete Room Data
// @Description  Delete Room Data by ID.
// @Tags 		 Rooms
// @Accept       json
// @Produce      json
// @Param id   	path     int  	true "room id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms/{id} [DELETE]
func (handler roomHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := model.Room{ID: id}

	err := handler.svc.DeleteRoom(ctx, &data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewRoomHandler(svc model.IStoreService, router gin.IRoutes) {
	handler := roomHandler{svc: svc}
	router.GET("/rooms", handler.fetch)
	router.POST("/rooms", handler.store)
	router.PUT("/rooms/:id", handler.update)
	router.DELETE("/rooms/:id", handler.destroy)
}
