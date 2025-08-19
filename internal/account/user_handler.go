package account

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"

	_ "github.com/aasumitro/posbe/internal/model"
)

type userHandler struct {
	svc IAccountService
}

// users godoc
// @Schemes
// @Summary User List
// @Description Get User List.
// @Tags Account Module - Users
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.User} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users [GET]
func (handler userHandler) fetch(ctx *gin.Context) {
	users, err := handler.svc.Users(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, users)
}

// users godoc
// @Schemes
// @Summary Self Update User Profile
// @Description Self Update User Profile
// @Tags Account Module - Users
// @Accept mpfd
// @Produce json
// @Param name 		formData string false "full name"
// @Param username 	formData string false "username"
// @Param email 	formData string false "email address"
// @Success 200 {object} utils.SuccessRespond{data=model.User} "OK RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users [PUT]
func (handler userHandler) updateProfile(ctx *gin.Context) {
	uid, ok := ctx.Get("user_id")
	if !ok {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest,
			"invalid user id")
		return
	}

	var form model.User
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	form.ID = int(uid.(float64))
	user, err := handler.svc.UpdateUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, user)
}

// users godoc
// @Schemes
// @Summary Self Update User Password
// @Description Self Update User Password
// @Tags Account Module - Users
// @Accept mpfd
// @Produce json
// @Param password 	formData string false "password"
// @Param new_password 	formData string false "new_password"
// @Success 200 {object} utils.SuccessRespond{data=model.User} "OK RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD REQUEST RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/users [PATCH]
func (handler userHandler) updatePassword(ctx *gin.Context) {
	uid, ok := ctx.Get("user_id")
	if !ok {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest,
			"invalid user id")
		return
	}

	var form UpdatePasswordForm

	// bind user input
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// validate user input in advance
	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	// call action
	form.ID = int(uid.(float64))
	if err := handler.svc.UpdateUserPassword(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	// remove cookie so user need to re-validate the state
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "authn", Value: "", MaxAge: 0, Path: "/", // Secure: true,
		Expires: time.Now().Add(-time.Hour),
	})

	utils.NewHTTPRespond(ctx, http.StatusUnauthorized, map[string]any{
		"message":        "Password updated successfully. Please log in again.",
		"login_required": true,
	})
}

// users godoc
// @Schemes
// @Summary Show User
// @Description Get User By ID.
// @Tags Account Module - Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.SuccessRespond{data=model.User} "OK RESPOND"
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

	user, err := handler.svc.UserByID(ctx, id)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, user)
}

// users godoc
// @Schemes
// @Summary Store User Data
// @Description Create new User.
// @Tags Account Module - Users
// @Accept mpfd
// @Produce json
// @Param role_id 	formData string true "role id"
// @Param name 		formData string true "full name"
// @Param username 	formData string true "username"
// @Param email 	formData string false "email address"
// @Param phone 	formData string false "phone number"
// @Param password 	formData string true "password"
// @Success 201 {object} utils.SuccessRespond{data=model.User} "CREATED RESPOND"
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

	if form.Password == "" {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity,
			"password is required")
		return
	}

	user, err := handler.svc.CreateUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, user)
}

// users godoc
// @Schemes
// @Summary Update User Data
// @Description Update Specified User Data by ID.
// @Tags Account Module - Users
// @Accept mpfd
// @Produce json
// @Param id   		path     int  	true "user id"
// @Param role_id 	formData string false "role id"
// @Param name 		formData string false "full name"
// @Param username 	formData string false "username"
// @Param email 	formData string false "email address"
// @Success 200 {object} utils.SuccessRespond{data=model.User} "OK RESPOND"
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
	user, err := handler.svc.UpdateUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, user)
}

// users godoc
// @Schemes
// @Summary Destroy User Data
// @Description Delete User By ID.
// @Tags Account Module - Users
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
	err := handler.svc.RemoveUser(ctx, &data)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusNoContent, nil)
}

func NewUserHandler(accountService IAccountService, router gin.IRoutes) {
	handler := userHandler{svc: accountService}
	router.GET("/users", handler.fetch)
	router.PUT("/users", handler.updateProfile)
	router.PATCH("/users", handler.updatePassword)
	router.GET("/users/:id", handler.show)
	// only admin can access this route
	authz := utils.AuthZ([]string{"admin"})
	router.POST("/users", authz, handler.store)
	router.PUT("/users/:id", authz, handler.update)
	router.DELETE("/users/:id", authz, handler.destroy)
}
