package account

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	svc IAccountService
}

// roles godoc
// @Schemes
// @Summary Role List
// @Description Get Role List.
// @Tags Users Roles
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessRespond{data=[]model.Role} "OK RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/roles [GET]
func (handler roleHandler) fetch(ctx *gin.Context) {
	roles, err := handler.svc.Roles(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, roles)
}

func NewRoleHandler(accountService IAccountService, router gin.IRoutes) {
	handler := roleHandler{svc: accountService}
	router.GET("/roles", handler.fetch)
}
