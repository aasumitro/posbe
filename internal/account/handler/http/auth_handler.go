package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/http/middleware"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthHandler struct {
	svc    domain.IAccountService
	config *config.Config
}

// login godoc
// @Schemes
// @Summary Logged User In
// @Description Generate Access Token (JWT).
// @Tags Auth
// @Accept mpfd
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 201 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/login [POST]
func (handler AuthHandler) login(ctx *gin.Context) {
	var form domain.LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := handler.svc.VerifyUserCredentials(form.Username, form.Password)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	if token, claimErr := utils.ClaimJWTToken(
		handler.config.AppName,
		handler.config.JWTLifetime,
		data,
		[]byte(handler.config.JWTSecretKey),
	); claimErr != nil {
		utils.NewHttpRespond(ctx, http.StatusInternalServerError, nil)
		return
	} else {
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:   "jwt",
			Value:  token,
			MaxAge: 0,
			Path:   "/",
		})
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, data)
}

// logout godoc
// @Schemes
// @Summary Logged User Out
// @Description Remove JWT Cookie
// @Tags Auth
// @Accept mpfd
// @Produce json
// @Success 200 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Router /v1/logout [POST]
func (handler AuthHandler) logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		MaxAge:  0,
		Path:    "/",
		Expires: time.Now().Add(-time.Hour),
	})

	utils.NewHttpRespond(ctx, http.StatusOK, "LOGGED_OUT")
}

func NewAuthHandler(accountService domain.IAccountService, config *config.Config, router *gin.RouterGroup) {
	handler := AuthHandler{svc: accountService, config: config}
	router.POST("/login", handler.login)
	router.POST("/logout", handler.logout).
		Use(middleware.Auth(config.JWTSecretKey))
}
