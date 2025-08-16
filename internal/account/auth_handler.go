package account

import (
	"net/http"
	"time"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	svc IAccountService
}

// login godoc
// @Schemes
// @Summary Logged User In
// @Description Generate Access Token (JWT).
// @Tags Account Auth
// @Accept mpfd
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 201 {object} utils.SuccessRespond{Data=domain.User} "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /api/v1/login [POST]
func (handler authHandler) login(ctx *gin.Context) {
	var form LoginForm

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
	
	// verify user credentials
	user, token, err := handler.svc.AuthenticateUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	// set cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "authn", Value: token, MaxAge: 0,
		Path: "/", HttpOnly: true, // Secure:   true,
	})

	// return data to users
	utils.NewHTTPRespond(ctx, http.StatusCreated,
		map[string]interface{}{"user": user, "token": token})
}

// logout godoc
// @Schemes
// @Summary Logged User Out
// @Description Remove JWT Cookie
// @Tags Account Auth
// @Accept mpfd
// @Produce json
// @Success 200 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Router /api/v1/logout [POST]
func (handler authHandler) logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "authn", Value: "", MaxAge: 0, Path: "/", // Secure: true,
		Expires: time.Now().Add(-time.Hour),
	})

	utils.NewHTTPRespond(ctx, http.StatusUnauthorized, "LOGGED_OUT")
}

func NewAuthHandler(accountService IAccountService, router *gin.RouterGroup) {
	handler := authHandler{svc: accountService}
	router.POST("/login", handler.login)
	authn := router.Use(utils.AuthN())
	authn.POST("/logout", handler.logout)
}
