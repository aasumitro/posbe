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

func NewAttributeCategoryHandler(service IAttributeService, router gin.IRoutes) {
	handler := attributeCategoryHandler{service: service}
	router.GET("/categories", handler.fetch)
	authz := router.Use(utils.AuthZ([]string{"admin"}))
	authz.POST("/categories", handler.add)
	authz.PATCH("/categories/:id", handler.edit)
	authz.DELETE("/categories/:id", handler.destroy)
	authz.DELETE("/categories/:id/subcategories/:sid", handler.destroySub)
}
