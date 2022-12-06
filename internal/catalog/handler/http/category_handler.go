package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type categoryHandler struct {
	svc domain.ICatalogCommonService
}

// categories godoc
// @Schemes
// @Summary 	 Categories List
// @Description  Get Categories List.
// @Tags 		 Categories
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]domain.Category} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/categories [GET]
func (handler categoryHandler) fetch(ctx *gin.Context) {
	data, err := handler.svc.CategoryList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, data)
}

// categories godoc
// @Schemes
// @Summary 	 Store Category Data
// @Description  Create new Category.
// @Tags 		 Categories
// @Accept       mpfd
// @Produce      json
// @Param name 			formData string true "name"
// @Success 201 {object} utils.SuccessRespond{data=domain.Category} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/categories [POST]
func (handler categoryHandler) store(ctx *gin.Context) {
	var form domain.Category
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := handler.svc.AddCategory(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, data)
}

// categories godoc
// @Schemes
// @Summary 	 Update Category Data
// @Description  Update Category Data by ID.
// @Tags 		 Categories
// @Accept       mpfd
// @Produce      json
// @Param id   			path     int  	true "category id"
// @Param name 			formData string true "name"
// @Success 200 {object} utils.SuccessRespond{data=domain.Category} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/categories/{id} [PUT]
func (handler categoryHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Category
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	data, err := handler.svc.EditCategory(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, data)
}

// categories godoc
// @Schemes
// @Summary 	 Delete Category Data
// @Description  Delete Category Data by ID.
// @Tags 		 Categories
// @Accept       json
// @Produce      json
// @Param id   			path     int  	true "category id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/categories/{id} [DELETE]
func (handler categoryHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := domain.Category{ID: id}

	err := handler.svc.DeleteCategory(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewCategoryHandler(svc domain.ICatalogCommonService, router gin.IRoutes) {
	handler := categoryHandler{svc: svc}
	router.GET("/categories", handler.fetch)
	router.POST("/categories", handler.store)
	router.PUT("/categories/:id", handler.update)
	router.DELETE("/categories/:id", handler.destroy)
}
