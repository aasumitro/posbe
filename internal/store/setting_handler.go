package store

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type settingHandler struct {
	service IStoreSettingService
}

func (handler settingHandler) fetch(ctx *gin.Context) {
	prefs, err := handler.service.AllSetting(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, prefs)
}

func (handler settingHandler) edit(ctx *gin.Context) {
	var form SettingForm

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

	if err := handler.service.UpdateSetting(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func NewSettingHandler(service IStoreSettingService, router gin.IRoutes) {
	handler := &settingHandler{service: service}
	router.GET("/settings", handler.fetch)
	router.PATCH("/settings", handler.edit)
}
