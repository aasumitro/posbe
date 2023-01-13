package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type tableHandler struct {
	svc domain.IStoreService
}

// tables godoc
// @Schemes
// @Summary 	 Table List
// @Description  Get Table List.
// @Tags 		 Tables
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]domain.Table} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/tables [GET]
func (handler tableHandler) fetch(ctx *gin.Context) {
	tables, err := handler.svc.TableList()
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, tables)
}

// tables godoc
// @Schemes
// @Summary 	 Store Table Data
// @Description  Create new Table.
// @Tags 		 Tables
// @Accept       mpfd
// @Produce      json
// @Param floor_id 	formData string true "floor id"
// @Param name 		formData string true "name"
// @Param x_pos 	formData string true "x position"
// @Param y_pos 	formData string true "y position"
// @Param w_size 	formData string true "weight"
// @Param h_size 	formData string true "height"
// @Param capacity 	formData string true "capacity"
// @Param type 		formData string true "type"
// @Success 201 {object} utils.SuccessRespond{data=domain.Table} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/tables [POST]
func (handler tableHandler) store(ctx *gin.Context) {
	var form domain.Table
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	table, err := handler.svc.AddTable(&form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, table)
}

// tables godoc
// @Schemes
// @Summary 	 Update Table Data
// @Description  Update Table Data by ID.
// @Tags 		 Tables
// @Accept       mpfd
// @Produce      json
// @Param id   		path     int  	true "table id"
// @Param floor_id 	formData string true "floor id"
// @Param name 		formData string true "name"
// @Param x_pos 	formData string true "x position"
// @Param y_pos 	formData string true "y position"
// @Param w_size 	formData string true "weight"
// @Param h_size 	formData string true "height"
// @Param capacity 	formData string true "capacity"
// @Param type 		formData string true "type"
// @Success 200 {object} utils.SuccessRespond{data=domain.Table} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/tables/{id} [PUT]
func (handler tableHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Table
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	form.ID = id
	table, err := handler.svc.EditTable(&form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, table)
}

// tables godoc
// @Schemes
// @Summary 	 Delete Table Data
// @Description  Delete Table Data by ID.
// @Tags 		 Tables
// @Accept       json
// @Produce      json
// @Param id   	path     int  	true "table id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/tables/{id} [DELETE]
func (handler tableHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := domain.Table{ID: id}

	err := handler.svc.DeleteTable(&data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewTableHandler(svc domain.IStoreService, router gin.IRoutes) {
	handler := tableHandler{svc: svc}
	router.GET("/tables", handler.fetch)
	router.POST("/tables", handler.store)
	router.PUT("/tables/:id", handler.update)
	router.DELETE("/tables/:id", handler.destroy)
}
