package catalog

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type attributeUnitHandler struct {
	service IAttributeService
}

// units godoc
// @Schemes
// @Summary Units List
// @Description Get Units List.
// @Tags Attribute Units
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Unit} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units [GET]
func (handler attributeUnitHandler) fetch(ctx *gin.Context) {
	data, err := handler.service.UnitList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// units godoc
// @Schemes
// @Summary Store Unit Data
// @Description Create new Unit.
// @Tags Attribute Units
// @Accept json
// @Produce json
// @Param body formData UnitForm true "body"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units [POST]
func (handler attributeUnitHandler) add(ctx *gin.Context) {
	var form UnitForm

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

	if err := handler.service.CreateUnit(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

// units godoc
// @Schemes
// @Summary Update Unit Data
// @Description Update Unit Data by ID.
// @Tags Attribute Units
// @Accept json
// @Produce json
// @Param id   			path     int  	true "unit id"
// @Param body formData UnitForm true "body"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units/{id} [PATCH]
func (handler attributeUnitHandler) edit(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	var form UnitForm

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
	if err := handler.service.UpdateUnit(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

// units godoc
// @Schemes
// @Summary Delete Unit Data
// @Description Delete Unit Data by ID.
// @Tags Attribute Units
// @Accept json
// @Produce json
// @Param id path int true "unit id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units/{id} [DELETE]
func (handler attributeUnitHandler) destroy(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := handler.service.DeleteUnit(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewAttributeUnitHandler(service IAttributeService, router *gin.RouterGroup) {
	handler := attributeUnitHandler{service: service}
	router.GET("/units", handler.fetch)
	authz := router.Group("/units")
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		authz.POST(utils.EmptyPath, handler.add)
		authz.PATCH("/:id", handler.edit)
		authz.DELETE("/:id", handler.destroy)
	}
}
