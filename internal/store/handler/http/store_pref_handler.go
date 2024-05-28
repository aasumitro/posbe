package http

import (
	"net/http"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type storePrefHandler struct {
	svc model.IStorePrefService
}

// store godoc
// @Schemes
// @Summary Store Settings
// @Description Get Store Settings List.
// @Tags Store
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.StoreSetting} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/store/prefs [GET]
func (handler storePrefHandler) fetch(ctx *gin.Context) {
	prefs, err := handler.svc.AllPrefs(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, prefs)
}

// store godoc
// @Schemes
// @Summary Update Floor Data
// @Description Update Floor Data by ID.
// @Tags Store
// @Accept mpfd
// @Produce json
// @Param key formData string true "key"
// @Param value formData string true "value"
// @Success 200 {object} utils.SuccessRespond{data=model.StoreSetting} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router  /api/v1/store/prefs [PUT]
func (handler storePrefHandler) update(ctx *gin.Context) {
	var form model.StorePref
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}
	pref, err := handler.svc.UpdatePrefs(ctx, form.Key, form.Value)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, pref)
}

func NewStorePrefHandler(svc model.IStorePrefService, router gin.IRoutes) {
	handler := storePrefHandler{svc: svc}
	router.GET("/store/prefs", handler.fetch)
	router.PUT("/store/prefs", handler.update)
}
