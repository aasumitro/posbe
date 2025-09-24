package account

import (
	"net/http"
	"strings"
	"time"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"

	_ "github.com/aasumitro/posbe/internal/model"
)

type authHandler struct {
	svc IAccountService
}

// login godoc
// @Schemes
// @Summary Logged User In
// @Description Generate Access Token (JWT).
// @Tags Account Module - Auth
// @Accept mpfd
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 201 {object} utils.SuccessRespond{Data=model.User} "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /api/v1/auth/login [POST]
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
	user, err := handler.svc.AuthenticateUser(ctx, &form)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	// set cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "access_token", Value: user.AccessToken, MaxAge: 0,
		Path: "/", HttpOnly: true, // Secure:   true,
	})

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "refresh_token", Value: user.RefreshToken, MaxAge: 0,
		Path: "/", HttpOnly: true, // Secure:   true,
	})

	// return data to users
	utils.NewHTTPRespond(ctx, http.StatusCreated, map[string]interface{}{
		"user": user, "token": map[string]interface{}{
			"access_token":  user.AccessToken,
			"refresh_token": user.RefreshToken,
		},
	})
}

// logout godoc
// @Schemes
// @Summary Refresh User Access Token
// @Description Remove JWT Cookie
// @Tags Account Module - Auth
// @Accept mpfd
// @Produce json
// @Success 201 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED"
// @Router /api/v1/auth/refresh [POST]
func (handler authHandler) refresh(ctx *gin.Context) {
	token := func() string {
		// Check cookie first
		if tokenCookie, err := ctx.Request.Cookie("refresh_token"); err == nil {
			return tokenCookie.Value
		}
		// Check Authorization header
		authHeader := strings.TrimSpace(ctx.GetHeader("X-REFRESH-TOKEN"))
		if authHeader != "" {
			if strings.HasPrefix(authHeader, "Bearer ") {
				return strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			}
			return authHeader
		}
		return ""
	}()

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			"invalid refresh token")
	}

	// verify user credentials
	user, err := handler.svc.RefreshToken(ctx, token)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	// set cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "access_token", Value: user.AccessToken, MaxAge: 0,
		Path: "/", HttpOnly: true, // Secure:   true,
	})

	ctx.Set("user_id", user.ID)
	ctx.Set("role_id", user.Role.Name)

	// return data to users
	utils.NewHTTPRespond(ctx, http.StatusCreated, map[string]interface{}{
		"user": user, "token": map[string]interface{}{
			"access_token": user.AccessToken,
		},
	})
}

// logout godoc
// @Schemes
// @Summary Logged User Out
// @Description Remove JWT Cookie
// @Tags Account Module - Auth
// @Accept mpfd
// @Produce json
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED"
// @Router /api/v1/auth/logout [POST]
func (handler authHandler) logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "access_token", Value: "", MaxAge: 0, Path: "/", // Secure: true,
		Expires: time.Now().Add(-time.Hour),
	})

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "refresh_token", Value: "", MaxAge: 0, Path: "/", // Secure: true,
		Expires: time.Now().Add(-time.Hour),
	})

	utils.NewHTTPRespond(ctx, http.StatusUnauthorized, "LOGGED_OUT")
}

func NewAuthHandler(accountService IAccountService, router *gin.RouterGroup) {
	handler := authHandler{svc: accountService}
	router = router.Group("auth")
	router.POST("/login", handler.login)
	router.POST("/refresh", handler.refresh)
	// TODO: router.POST("/password/forgot", handler.forgot)
	// TODO: router.POST("/password/reset", handler.reset)
	authn := router.Group("")
	authn.Use(utils.AuthN())
	{
		authn.POST("/logout", handler.logout)
	}
}
