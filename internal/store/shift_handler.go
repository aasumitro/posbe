package store

import (
	"net/http"
	"slices"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	service IStoreShiftService
}

// shift godoc
// @Schemes
// @Summary Shift List
// @Description Get Shift List.
// @Tags Shifts
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Table} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/shifts [GET]
func (handler shiftHandler) fetch(ctx *gin.Context) {
	prefs, err := handler.service.ShiftList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, prefs)
}

// shift godoc
// @Schemes
// @Summary Shift Detail
// @Description Get Shift Detail.
// @Tags Shifts
// @Accept mpfd
// @Produce json
// @Param id   		path     int  	true "table id"
// @Success 200 {object} utils.SuccessRespond{data=[]model.Table} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/shifts/{id} [GET]
func (handler shiftHandler) show(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	data, err := handler.service.ShiftDetail(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func (handler shiftHandler) store(ctx *gin.Context) {
	var form ShiftForm

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

	if err := handler.service.CreateShift(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

func (handler shiftHandler) update(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	var form ShiftForm

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
	if err := handler.service.UpdateShift(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func (handler shiftHandler) destroy(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := handler.service.DeleteShift(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func (handler shiftHandler) active(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	actionParams := ctx.Param("action")
	if !slices.Contains([]string{"open", "close"}, actionParams) {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, "invalid action parameter")
		return
	}

	uid, ok := ctx.Get("user_id")
	if !ok {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	var form ActiveShiftForm
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	form.ShiftID = id
	form.Action = actionParams
	form.UserID = int64(uid.(float64))

	if err := handler.service.ActiveShiftAction(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func NewShiftHandler(service IStoreShiftService, router gin.IRoutes) {
	handler := shiftHandler{service: service}
	router.GET("/shifts", handler.fetch)
	router.GET("/shifts/:id", handler.show)
	authz := router.Use(utils.AuthZ([]string{"admin"}))
	authz.POST("/shifts", handler.store)
	authz.PATCH("/shifts/:id", handler.update)
	authz.DELETE("/shifts/:id", handler.destroy)
	authz.POST("/shifts/:id/:action", handler.active)
}
