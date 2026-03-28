package customer

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service ICustomerService
}

func (h handler) fetch(ctx *gin.Context) {
	var q RequestQuery

	if err := ctx.ShouldBindQuery(&q); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := q.Validate(ctx); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	data, err := h.service.List(ctx, &q)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func (h handler) add(ctx *gin.Context) {
	var form RequestForm

	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ActionAdd = true
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	if err := h.service.Create(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusCreated, nil)
}

func (h handler) edit(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}
	var form RequestForm

	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ActionAdd = false
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	form.ID = id

	if err := h.service.Update(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func (h handler) destroy(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.service.Destroy(ctx, id); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewHandler(service ICustomerService, router *gin.RouterGroup) {
	h := &handler{service: service}
	authz := router.Group("/customers")
	authz.GET(utils.EmptyPath, h.fetch)
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		authz.POST(utils.EmptyPath, h.add)
		authz.PATCH("/:id", h.edit)
		authz.DELETE("/:id", h.destroy)
	}
}
