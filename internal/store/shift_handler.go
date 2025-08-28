package store

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	service IStoreShiftService
}

func (handler shiftHandler) fetch(ctx *gin.Context) {
	prefs, err := handler.service.ShiftList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, prefs)
}

func (handler shiftHandler) show(ctx *gin.Context) {
}

func (handler shiftHandler) add(ctx *gin.Context) {
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

func NewShiftHandler(service IStoreShiftService, router gin.IRoutes) {
	handler := shiftHandler{service: service}
	router.GET("/shifts", handler.fetch)
	router.GET("/shifts/:id", handler.show)
	router.POST("/shifts", handler.add)
}
