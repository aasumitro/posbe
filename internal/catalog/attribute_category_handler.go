package catalog

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type attributeCategoryHandler struct {
	service IAttributeService
}

// categories godoc
// @Schemes
// @Summary Categories List
// @Description Get Categories List.
// @Tags Attribute Categories
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Category} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/categories [GET]
func (handler attributeCategoryHandler) fetch(ctx *gin.Context) {
	data, err := handler.service.CategoryList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

// categories godoc
// @Schemes
// @Summary Store Category Data
// @Description Create new Category.
// @Tags Attribute Categories
// @Accept json
// @Produce json
// @Param body formData NewCategoryForm true "body"
// @Success 201 {object} utils.SuccessRespond{data=model.Category} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/categories [POST]
func (handler attributeCategoryHandler) add(ctx *gin.Context) {
	var form NewCategoryForm

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

	if err := handler.service.CreateCategory(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

// categories godoc
// @Schemes
// @Summary Update Category Data
// @Description Update Category Data by ID.
// @Tags Attribute Categories
// @Accept json
// @Produce json
// @Param id path int true "category id"
// @Param body formData EditCategoryForm true "body"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/categories/{id} [PATCH]
func (handler attributeCategoryHandler) edit(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	var form EditCategoryForm

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
	if err := handler.service.UpdateCategory(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

// categories godoc
// @Schemes
// @Summary Delete Category Data
// @Description Delete Category Data by ID.
// @Tags Attribute Categories
// @Accept json
// @Produce json
// @Param id path int true "category id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/categories/{id} [DELETE]
func (handler attributeCategoryHandler) destroy(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := handler.service.DeleteCategory(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

// categories godoc
// @Schemes
// @Summary Delete Subcategory in Category Data
// @Description Delete Subcategory in Category Data by ID.
// @Tags Attribute Categories
// @Accept json
// @Produce json
// @Param id path int true "category id"
// @Param sid path int true "subcategory id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/categories/{id}/subcategories/{sid} [DELETE]
func (handler attributeCategoryHandler) destroySub(ctx *gin.Context) {
	cid, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	sid, ok := utils.GetIDParam(ctx, "sid")
	if !ok {
		return
	}

	if err := handler.service.DeleteSubcategory(ctx, cid, sid); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewAttributeCategoryHandler(service IAttributeService, router *gin.RouterGroup) {
	handler := attributeCategoryHandler{service: service}
	router.GET("/categories", handler.fetch)
	authz := router.Group("/categories")
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		authz.POST(utils.EmptyPath, handler.add)
		authz.PATCH("/:id", handler.edit)
		authz.DELETE("/:id", handler.destroy)
		authz.DELETE("/:id/subcategories/:sid", handler.destroySub)
	}
}
