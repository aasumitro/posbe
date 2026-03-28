package catalog

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type productAddonHandler struct {
	service IProductService
}

// addons godoc
// @Schemes
// @Summary Addons List
// @Description Get Addons List.
// @Tags Catalog Addons
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Addon} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/product-addons [GET]
func (handler productAddonHandler) fetch(ctx *gin.Context) {
	data, err := handler.service.AddonList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// addons godoc
// @Schemes
// @Summary Store addon Data
// @Description Create new addon.
// @Tags Catalog Addons
// @Accept mpfd
// @Produce json
// @Param body 			AddonForm string true "body"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/product-addons [POST]
func (handler productAddonHandler) add(ctx *gin.Context) {
	var form AddonForm

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

	if err := handler.service.CreateAddon(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

// addons godoc
// @Schemes
// @Summary Update addon Data
// @Description Update addon Data by ID.
// @Tags Catalog Addons
// @Accept mpfd
// @Produce json
// @Param id   			path     int  	true "addon id"
// @Param body 			AddonForm string true "body"
// @Success 200 {object} utils.SuccessRespond{data=model.Addon} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/product-addons/{id} [PATCH]
func (handler productAddonHandler) edit(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	var form AddonForm

	// bind user input
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate user input in advance
	form.ActionAdd = false
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	form.ID = id
	if err := handler.service.UpdateAddon(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

// addons godoc
// @Schemes
// @Summary Delete addon Data
// @Description Delete addon Data by ID.
// @Tags Catalog Addons
// @Accept json
// @Produce json
// @Param id path int true "addon id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/product-addons/{id} [DELETE]
func (handler productAddonHandler) destroy(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := handler.service.DeleteAddon(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewProductAddonHandler(service IProductService, router *gin.RouterGroup) {
	handler := productAddonHandler{service: service}
	router.GET("/product-addons", handler.fetch)
	authz := router.Group("/product-addons")
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		authz.POST(utils.EmptyPath, handler.add)
		authz.PATCH("/:id", handler.edit)
		authz.DELETE("/:id", handler.destroy)
	}
}
