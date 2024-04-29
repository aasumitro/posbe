package http

import (
	"net/http"
	"strconv"

	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type variantHandler struct {
	svc model.ICatalogProductService
}

// product_variants godoc
// @Schemes
// @Summary Store variant Data
// @Description Create new variant Data.
// @Tags Product Variants
// @Accept mpfd
// @Produce json
// @Param product_id 	formData string true "product_id"
// @Param unit_id 		formData string true "unit_id"
// @Param unit_size 	formData string true "unit_size"
// @Param type	 		formData string true "type" Enums(none, size)
// @Param name 			formData string true "name"
// @Param description 	formData string false "description"
// @Param price 		formData int true "price"
// @Success 201 {object} utils.SuccessRespond{data=model.ProductVariant} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/products/variants [POST]
func (handler variantHandler) store(ctx *gin.Context) {
	var form model.ProductVariant
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := handler.svc.AddProductVariant(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, data)
}

// product_variants godoc
// @Schemes
// @Summary Update variant Data
// @Description Update variant Data by ID.
// @Tags Product Variants
// @Accept mpfd
// @Produce json
// @Param id   			path     int  	true "variant id"
// @Param product_id 	formData string true "product_id"
// @Param unit_id 		formData string true "unit_id"
// @Param unit_size 	formData string true "unit_size"
// @Param type	 		formData string true "type" Enums(none, size)
// @Param name 			formData string true "name"
// @Param description 	formData string false "description"
// @Param price 		formData int true "price"
// @Success 200 {object} utils.SuccessRespond{data=model.ProductVariant} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/products/variants/{id} [PUT]
func (handler variantHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form model.ProductVariant
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	data, err := handler.svc.EditProductVariant(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// product_variants godoc
// @Schemes
// @Summary Delete variant Data
// @Description Delete variant Data by ID.
// @Tags Product Variants
// @Accept json
// @Produce json
// @Param id path int true "variant id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/products/variants/{id} [DELETE]
func (handler variantHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := model.ProductVariant{ID: id}

	err := handler.svc.DeleteProductVariant(ctx, &data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewProductVariantHandler(svc model.ICatalogProductService, router gin.IRoutes) {
	handler := variantHandler{svc: svc}
	router.POST("/products/variants", handler.store)
	router.PUT("/products/variants/:id", handler.update)
	router.DELETE("/products/variants/:id", handler.destroy)
}
