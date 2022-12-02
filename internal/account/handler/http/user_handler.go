package http

import (
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userHandler struct {
	svc domain.IAccountService
}

func (handler userHandler) fetch(ctx *gin.Context) {
	users, err := handler.svc.UserList()
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, users)
}

func (handler userHandler) show(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}

	user, err := handler.svc.ShowUser(id)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, user)
}

func (handler userHandler) store(ctx *gin.Context) {
	var form domain.User
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := handler.svc.AddUser(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusCreated, user)
}

func (handler userHandler) update(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}

	var form domain.User
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHttpRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	form.ID = id
	user, err := handler.svc.EditUser(&form)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusOK, user)
}

func (handler userHandler) destroy(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, errParse := strconv.Atoi(idParams)
	if errParse != nil {
		utils.NewHttpRespond(ctx,
			http.StatusInternalServerError,
			errParse.Error())
		return
	}
	data := domain.User{ID: id}

	err := handler.svc.DeleteUser(&data)
	if err != nil {
		utils.NewHttpRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHttpRespond(ctx, http.StatusNoContent, nil)
}

func NewUserHandler(accountService domain.IAccountService, router *gin.RouterGroup) {
	handler := userHandler{svc: accountService}
	router.GET("/users", handler.fetch)
	router.GET("/users/:id", handler.show)
	router.POST("/users", handler.store)
	router.PUT("/users/:id", handler.update)
	router.DELETE("/users/:id", handler.destroy)
}
