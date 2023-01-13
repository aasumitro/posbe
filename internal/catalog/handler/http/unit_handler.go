package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type unitHandler struct {
	svc domain.ICatalogCommonService
}

// units godoc
// @Schemes
// @Summary 	 Units List
// @Description  Get Units List.
// @Tags 		 Units
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]domain.Unit} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/units [GET]
func (handler unitHandler) fetch(ctx *gin.Context) {
	data, err := handler.svc.UnitList()
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// units godoc
// @Schemes
// @Summary 	 Store Unit Data
// @Description  Create new Unit.
// @Tags 		 Units
// @Accept       mpfd
// @Produce      json
// @Param magnitude 	formData string true "magnitude"
// @Param name 			formData string true "name"
// @Param symbol 		formData string true "symbol"
// @Success 201 {object} utils.SuccessRespond{data=domain.Unit} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/units [POST]
func (handler unitHandler) store(ctx *gin.Context) {
	var form domain.Unit
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := handler.svc.AddUnit(&form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, data)
}

// units godoc
// @Schemes
// @Summary 	 Update Unit Data
// @Description  Update Unit Data by ID.
// @Tags 		 Units
// @Accept       mpfd
// @Produce      json
// @Param id   			path     int  	true "unit id"
// @Param magnitude 	formData string true "magnitude"
// @Param name 			formData string true "name"
// @Param symbol 		formData string true "symbol"
// @Success 200 {object} utils.SuccessRespond{data=domain.Unit} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/units/{id} [PUT]
func (handler unitHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Unit
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	data, err := handler.svc.EditUnit(&form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// units godoc
// @Schemes
// @Summary 	 Delete Unit Data
// @Description  Delete Unit Data by ID.
// @Tags 		 Units
// @Accept       json
// @Produce      json
// @Param id   			path     int  	true "unit id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/units/{id} [DELETE]
func (handler unitHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := domain.Unit{ID: id}

	err := handler.svc.DeleteUnit(&data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewUnitHandler(svc domain.ICatalogCommonService, router gin.IRoutes) {
	handler := unitHandler{svc: svc}
	router.GET("/units", handler.fetch)
	router.POST("/units", handler.store)
	router.PUT("/units/:id", handler.update)
	router.DELETE("/units/:id", handler.destroy)
}
