package http

import (
	"net/http"
	"strconv"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type subcategoryHandler struct {
	svc model.ICatalogCommonService
}

// subcategories godoc
// @Schemes
// @Summary Subcategories List
// @Description Get Subcategories List.
// @Tags Product Subcategories
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Subcategory} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/subcategories [GET]
func (handler subcategoryHandler) fetch(ctx *gin.Context) {
	data, err := handler.svc.SubcategoryList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// subcategories godoc
// @Schemes
// @Summary Store Subcategory Data
// @Description Create new Subcategory.
// @Tags Product Subcategories
// @Accept mpfd
// @Produce
// @Param name 			formData string true "name"
// @Param category_id 	formData string true "category_id"
// @Success 201 {object} utils.SuccessRespond{data=model.Subcategory} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/subcategories [POST]
func (handler subcategoryHandler) store(ctx *gin.Context) {
	var form model.Subcategory
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := handler.svc.AddSubcategory(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, data)
}

// subcategories godoc
// @Schemes
// @Summary Update Subcategory Data
// @Description Update Subcategory Data by ID.
// @Tags Product Subcategories
// @Accept mpfd
// @Produce json
// @Param id   			path     int  	true "subcategory id"
// @Param category_id 	formData string true "category_id"
// @Param name 			formData string true "name"
// @Success 200 {object} utils.SuccessRespond{data=model.Subcategory} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/subcategories/{id} [PUT]
func (handler subcategoryHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form model.Subcategory
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	data, err := handler.svc.EditSubcategory(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// subcategories godoc
// @Schemes
// @Summary Delete Subcategory Data
// @Description Delete Subcategory Data by ID.
// @Tags Product Subcategories
// @Accept json
// @Produce json
// @Param id path int true "subcategory id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/subcategories/{id} [DELETE]
func (handler subcategoryHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := model.Subcategory{ID: id}

	err := handler.svc.DeleteSubcategory(ctx, &data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewSubcategoryHandler(svc model.ICatalogCommonService, router gin.IRoutes) {
	handler := subcategoryHandler{svc: svc}
	router.GET("/subcategories", handler.fetch)
	router.POST("/subcategories", handler.store)
	router.PUT("/subcategories/:id", handler.update)
	router.DELETE("/subcategories/:id", handler.destroy)
}
