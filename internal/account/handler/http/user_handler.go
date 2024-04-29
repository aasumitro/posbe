package http

import (
	"net/http"
	"strconv"

	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	svc model.IAccountService
}

// users godoc
// @Schemes
// @Summary User List
// @Description Get User List.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]domain.User} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users [GET]
func (handler userHandler) fetch(ctx *gin.Context) {
	users, err := handler.svc.UserList(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, users)
}

// users godoc
// @Schemes
// @Summary Show User
// @Description Get User By ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.SuccessRespond{data=domain.User} "OK RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users/{id} [GET]
func (handler userHandler) show(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	user, err := handler.svc.ShowUser(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	user.Password = ""
	utils.NewHTTPRespond(ctx, http.StatusOK, user)
}

// users godoc
// @Schemes
// @Summary Store User Data
// @Description Create new User.
// @Tags Users
// @Accept mpfd
// @Produce json
// @Param role_id 	formData string true "role id"
// @Param name 		formData string true "full name"
// @Param username 	formData string true "username"
// @Param email 	formData string false "email address"
// @Param phone 	formData string false "phone number"
// @Param password 	formData string true "password"
// @Success 201 {object} utils.SuccessRespond{data=domain.User} "CREATED RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users [POST]
func (handler userHandler) store(ctx *gin.Context) {
	var form model.User
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	user, err := handler.svc.AddUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	user.Password = ""
	utils.NewHTTPRespond(ctx, http.StatusCreated, user)
}

// users godoc
// @Schemes
// @Summary Update User Data
// @Description Update Specified User Data by ID.
// @Tags Users
// @Accept mpfd
// @Produce json
// @Param id   		path     int  	true "user id"
// @Param role_id 	formData string false "role id"
// @Param name 		formData string false "full name"
// @Param username 	formData string false "username"
// @Param email 	formData string false "email address"
// @Param phone 	formData string false "phone number"
// @Param password 	formData string false "password"
// @Success 200 {object} utils.SuccessRespond{data=domain.User} "OK RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users/{id} [PUT]
func (handler userHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	var form model.User
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	form.ID = id
	user, err := handler.svc.EditUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	user.Password = ""
	utils.NewHTTPRespond(ctx, http.StatusOK, user)
}

// users godoc
// @Schemes
// @Summary Destroy User Data
// @Description Delete User By ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 204 "NO CONTENT RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users/{id} [DELETE]
func (handler userHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHTTPRespond(ctx,
			http.StatusBadRequest,
			errParse.Error())
		return
	}
	data := model.User{ID: id}
	err := handler.svc.DeleteUser(ctx, &data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewUserHandler(accountService model.IAccountService, router gin.IRoutes) {
	handler := userHandler{svc: accountService}
	router.GET("/users", handler.fetch)
	router.GET("/users/:id", handler.show)
	protectedRoute := router.Use(middleware.
		AcceptedRoles([]string{"admin"}))
	protectedRoute.POST("/users", handler.store)
	protectedRoute.PUT("/users/:id", handler.update)
	protectedRoute.DELETE("/users/:id", handler.destroy)
}
