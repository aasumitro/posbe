package http

import (
	"net/http"
	"time"

	"github.com/aasumitro/posbe/config"

	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc model.IAccountService
	jwt utils.IJSONWebToken
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
func (handler AuthHandler) login(ctx *gin.Context) {
	var form model.LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	data, err := handler.svc.VerifyUserCredentials(
		ctx, form.Username, form.Password)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	token, claimErr := handler.jwt.ClaimJWTToken(data)
	if claimErr != nil {
		utils.NewHTTPRespond(ctx, http.StatusInternalServerError, claimErr.Error())
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:   "jwt",
		Value:  token,
		MaxAge: 0,
		Path:   "/",
		// Secure:   true,
		HttpOnly: true,
	})
	utils.NewHTTPRespond(ctx, http.StatusCreated, data)
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
func (handler AuthHandler) logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		MaxAge:  0,
		Path:    "/",
		Expires: time.Now().Add(-time.Hour),
	})
	utils.NewHTTPRespond(ctx, http.StatusOK, "LOGGED_OUT")
}

func NewAuthHandler(accountService model.IAccountService, router *gin.RouterGroup) {
	handler := AuthHandler{svc: accountService, jwt: &utils.JSONWebToken{
		Issuer:    config.Instance.AppName,
		SecretKey: []byte(config.Instance.JWTSecretKey),
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(config.Instance.JWTLifetime) * time.Hour),
	}}
	router.POST("/login", handler.login)
	protectedRouter := router.Use(middleware.Auth())
	protectedRouter.POST("/logout", handler.logout)
}
