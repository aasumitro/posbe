package http

import (
	"net/http"
	"strconv"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	svc model.IStoreShiftService
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
	data, err := handler.svc.ShiftList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, data)
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
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data, err := handler.svc.ShiftDetail(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func NewShiftHandler(
	svc model.IStoreShiftService,
	router gin.IRoutes,
) {
	handler := &shiftHandler{svc: svc}
	router.GET("/shifts", handler.fetch)
	router.GET("/shifts/:id", handler.show)
}
