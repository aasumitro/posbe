package store

import (
	"context"
	"net/http"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type seatingHandler struct {
	service IStoreSeatingService
}

// floors godoc
// @Schemes
// @Summary Floor List
// @Description Get Store Floors List
// @Tags Store Floors
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=model.Floor} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/store/floors [GET]
func (handler seatingHandler) fetchFloor(ctx *gin.Context) {
	data, err := handler.service.FloorList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)

}

func (handler seatingHandler) showFloor(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	data, err := handler.service.FloorDetail(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func (handler seatingHandler) fetchFloorTable(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	data, err := handler.service.TableList(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func (handler seatingHandler) addFloor(ctx *gin.Context) {
	var form SeatingFloorRequest

	// bind user input
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate user input in advance
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	if err := handler.service.CreateFloor(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

func (handler seatingHandler) editFloor(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	var form SeatingFloorRequest

	// bind user input
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate user input in advance
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	form.ID = id
	if err := handler.service.UpdateFloor(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func (handler seatingHandler) destroyFloor(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := handler.service.DeleteFloor(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func (handler seatingHandler) showTable(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	data, err := handler.service.TableDetail(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func (handler seatingHandler) addTable(ctx *gin.Context) {
	var form SeatingTableRequest

	// bind user input
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate user input in advance
	form.ActionAdd = true
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	if err := handler.service.CreateTable(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

func (handler seatingHandler) editTable(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	var form SeatingTableRequest

	// bind user input
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate user input in advance
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	form.ID = id
	if err := handler.service.UpdateTable(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func (handler seatingHandler) destroyTable(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := handler.service.DeleteTable(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func (handler seatingHandler) events() {
	ctx := context.Background()
	utils.SubscribeEvent(ctx, model.TableUpdateEventKey, func(message *redis.Message) {
		handler.service.ProceedEvent(ctx, message.Payload)
	})
}

func NewSeatingHandler(service IStoreSeatingService, router *gin.RouterGroup) {
	handler := seatingHandler{service: service}
	router.GET("/floors", handler.fetchFloor)
	router.GET("/floors/:id", handler.showFloor)
	router.GET("/floors/:id/tables", handler.fetchFloorTable)
	router.GET("/tables/:id", handler.showTable)
	authz := router.Group(utils.EmptyPath)
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		// store floors endpoint
		authz.POST("/floors", handler.addFloor)
		authz.PATCH("/floors/:id", handler.editFloor)
		authz.DELETE("/floors/:id", handler.destroyFloor)
		// table mgmt endpoint
		authz.POST("/tables", handler.addTable)
		authz.PATCH("/tables/:id", handler.editTable)
		authz.DELETE("/tables/:id", handler.destroyTable)
	}
	go handler.events()
}
