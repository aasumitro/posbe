package http

import (
	"encoding/json"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"net/http"
	"strconv"
	"time"
)

type roleHandler struct {
	svc   domain.IAccountService
	cache *redis.Client
}

var roleCacheKey = "roles"

// roles godoc
// @Schemes
// @Summary 	 Role List
// @Description  Get Role List.
// @Tags 		 Roles
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]domain.Role} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/roles [GET]
func (handler roleHandler) fetch(ctx *gin.Context) {
	helper := utils.RedisCache{Ctx: ctx, RdpConn: handler.cache}
	roles, err := helper.CacheFirstData(&utils.CacheDataSupplied{
		Key: roleCacheKey,
		Ttl: time.Hour * 1,
		CbF: func() (data any, err *utils.ServiceError) {
			return handler.svc.RoleList()
		},
	})

	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	if data, ok := roles.(string); ok {
		var r []domain.Role
		_ = json.Unmarshal([]byte(data), &r)
		roles = r
	}

	utils.NewHttpRespond(ctx, http.StatusOK, roles)
}

// roles godoc
// @Schemes
// @Summary 	 Store Role Data
// @Description  Create new Role.
// @Tags 		 Roles
// @Accept       mpfd
// @Produce      json
// @Param name 			formData string true "name"
// @Param description 	formData string true "description"
// @Success 201 {object} utils.SuccessRespond{data=domain.Role} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/roles [POST]
func (handler roleHandler) store(ctx *gin.Context) {
	var form domain.Role
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	role, err := handler.svc.AddRole(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	handler.cache.Del(ctx, roleCacheKey)

	utils.NewHttpRespond(ctx, http.StatusCreated, role)
}

// roles godoc
// @Schemes
// @Summary 	 Update Role Data
// @Description  Update Role Data by ID.
// @Tags 		 Roles
// @Accept       mpfd
// @Produce      json
// @Param id   			path     int  	true "role id"
// @Param name 			formData string true "name"
// @Param description 	formData string true "description"
// @Success 200 {object} utils.SuccessRespond{data=domain.Role} "CREATED RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/roles/{id} [PUT]
func (handler roleHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}

	var form domain.Role
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = id
	role, err := handler.svc.EditRole(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	handler.cache.Del(ctx, roleCacheKey)

	utils.NewHttpRespond(ctx, http.StatusOK, role)
}

// roles godoc
// @Schemes
// @Summary 	 Delete Role Data
// @Description  Delete Role Data by ID.
// @Tags 		 Roles
// @Accept       json
// @Produce      json
// @Param id   			path     int  	true "role id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /v1/roles/{id} [DELETE]
func (handler roleHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := domain.Role{ID: id}

	err := handler.svc.DeleteRole(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	handler.cache.Del(ctx, roleCacheKey)

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewRoleHandler(accountService domain.IAccountService, router gin.IRoutes, cache *redis.Client) {
	handler := roleHandler{svc: accountService, cache: cache}
	router.GET("/roles", handler.fetch)
	protectedRoute := router.Use(middleware.
		AcceptedRoles([]string{"admin"}))
	protectedRoute.POST("/roles", handler.store)
	protectedRoute.PUT("/roles/:id", handler.update)
	protectedRoute.DELETE("/roles/:id", handler.destroy)
}
