package catalog

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type attributeCategoryHandler struct {
	service IAttributeService
}

func (handler attributeCategoryHandler) fetch(ctx *gin.Context) {
	data, err := handler.service.CategoryList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

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

func NewAttributeCategoryHandler(service IAttributeService, router gin.IRoutes) {
	handler := attributeCategoryHandler{service: service}
	router.GET("/categories", handler.fetch)
	authz := router.Use(utils.AuthZ([]string{"admin"}))
	authz.POST("/categories", handler.add)
	// authz.PATCH("/categories/:id", handler.edit)
	authz.DELETE("/categories/:id", handler.destroy)
	// authz.POST("/categories/:id/subcategories", handler.addSub)
	// authz.PATCH("/categories/:category_id/subcategories/:subcategory_id", handler.editSub)
	// authz.DELETE("/categories/:category_id/subcategories/:subcategory_id", handler.destroySub)
}
