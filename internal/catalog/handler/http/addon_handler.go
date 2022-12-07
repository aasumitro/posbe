package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type addonHandler struct {
	svc domain.ICatalogCommonService
}

// addons godoc
// @Schemes
// @Summary 	 Addons List
// @Description  Get Addons List.
// @Tags 		 Addons
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]domain.Addon} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/addons [GET]
func (handler addonHandler) fetch(ctx *gin.Context) {
	data, err := handler.svc.AddonList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, data)
}

// addons godoc
// @Schemes
// @Summary 	 Store addon Data
// @Description  Create new addon.
// @Tags 		 Addons
// @Accept       mpfd
// @Produce      json
// @Param name 			formData string true "name"
// @Param description 	formData string true "description"
// @Param price 		formData string true "price"
// @Success 201 {object} utils.SuccessRespond{data=domain.Addon} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/addons [POST]
func (handler addonHandler) store(ctx *gin.Context) {
	var form domain.Addon
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := handler.svc.AddAddon(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, data)
}

// addons godoc
// @Schemes
// @Summary 	 Update addon Data
// @Description  Update addon Data by ID.
// @Tags 		 Addons
// @Accept       mpfd
// @Produce      json
// @Param id   			path     int  	true "addon id"
// @Param name 			formData string true "name"
// @Param description 	formData string true "description"
// @Param price 		formData string true "price"
// @Success 200 {object} utils.SuccessRespond{data=domain.Addon} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/addons/{id} [PUT]
func (handler addonHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Addon
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	data, err := handler.svc.EditAddon(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, data)
}

// addons godoc
// @Schemes
// @Summary 	 Delete addon Data
// @Description  Delete addon Data by ID.
// @Tags 		 Addons
// @Accept       json
// @Produce      json
// @Param id   			path     int  	true "category id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/addons/{id} [DELETE]
func (handler addonHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := domain.Addon{ID: id}

	err := handler.svc.DeleteAddon(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewAddonHandler(svc domain.ICatalogCommonService, router gin.IRoutes) {
	handler := addonHandler{svc: svc}
	router.GET("/addons", handler.fetch)
	router.POST("/addons", handler.store)
	router.PUT("/addons/:id", handler.update)
	router.DELETE("/addons/:id", handler.destroy)
}
