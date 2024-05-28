package http

import (
	"net/http"
	"strconv"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type unitHandler struct {
	svc model.ICatalogCommonService
}

// units godoc
// @Schemes
// @Summary Units List
// @Description Get Units List.
// @Tags Product Units
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Unit} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units [GET]
func (handler unitHandler) fetch(ctx *gin.Context) {
	data, err := handler.svc.UnitList(ctx)
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
// @Tags Product Units
// @Accept mpfd
// @Produce json
// @Param magnitude formData string true "magnitude"
// @Param name formData string true "name"
// @Param symbol formData string true "symbol"
// @Success 201 {object} utils.SuccessRespond{data=model.Unit} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units [POST]
func (handler unitHandler) store(ctx *gin.Context) {
	var form model.Unit
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	data, err := handler.svc.AddUnit(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusCreated, data)
}

// units godoc
// @Schemes
// @Summary Update Unit Data
// @Description Update Unit Data by ID.
// @Tags Product Units
// @Accept mpfd
// @Produce json
// @Param id   			path     int  	true "unit id"
// @Param magnitude 	formData string true "magnitude"
// @Param name 			formData string true "name"
// @Param symbol 		formData string true "symbol"
// @Success 200 {object} utils.SuccessRespond{data=model.Unit} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units/{id} [PUT]
func (handler unitHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form model.Unit
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	data, err := handler.svc.EditUnit(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// units godoc
// @Schemes
// @Summary Delete Unit Data
// @Description Delete Unit Data by ID.
// @Tags Product Units
// @Accept json
// @Produce json
// @Param id path int true "unit id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/units/{id} [DELETE]
func (handler unitHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := model.Unit{ID: id}

	err := handler.svc.DeleteUnit(ctx, &data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewUnitHandler(svc model.ICatalogCommonService, router gin.IRoutes) {
	handler := unitHandler{svc: svc}
	router.GET("/units", handler.fetch)
	router.POST("/units", handler.store)
	router.PUT("/units/:id", handler.update)
	router.DELETE("/units/:id", handler.destroy)
}
