package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type roomHandler struct {
	svc domain.IStoreService
}

// rooms godoc
// @Schemes
// @Summary 	 Room List
// @Description  Get Room List.
// @Tags 		 Rooms
// @Accept       json
// @Produce      json
// @Success 200 {array} domain.Room "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms [GET]
func (handler roomHandler) fetch(ctx *gin.Context) {
	rooms, err := handler.svc.RoomList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, rooms)
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
// @Success 201 {object} domain.Room "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms [POST]
func (handler roomHandler) store(ctx *gin.Context) {
	var form domain.Room
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	room, err := handler.svc.AddRoom(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, room)
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
// @Success 200 {object} domain.Room "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/rooms/{id} [PUT]
func (handler roomHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Room
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	form.ID = id
	room, err := handler.svc.EditRoom(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, room)
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
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
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
	handler := roomHandler{svc: svc}
	router.GET("/rooms", handler.fetch)
	router.POST("/rooms", handler.store)
	router.PUT("/rooms/:id", handler.update)
	router.DELETE("/rooms/:id", handler.destroy)
}
