package store

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type settingHandler struct {
	service IStoreSettingService
}

// store godoc
// @Schemes
// @Summary Store Settings
// @Description Get Store Settings List.
// @Tags Store Setting
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.StoreSetting} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/store/settings [GET]
func (handler settingHandler) fetch(ctx *gin.Context) {
	prefs, err := handler.service.AllSetting(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, prefs)
}

// store godoc
// @Schemes
// @Summary Update Store Setting
// @Description Update Store Setting by Key (one or all)
// @Tags Store Setting
// @Accept json
// @Produce json
// @Param 		form		body 	SettingForm false "form request for setting"
// @Success 200 {object} utils.SuccessRespond{data=model.StoreSetting} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router  /api/v1/store/settings [PATCH]
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
