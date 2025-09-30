package order

import (
	"net/http"
	"slices"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	service IShiftService
}

func (handler shiftHandler) active(ctx *gin.Context) {
	id, ok := utils.GetIDParam(ctx, "id")
	if !ok {
		return
	}

	actionParams := ctx.Param("action")
	if !slices.Contains([]string{"open", "close"}, actionParams) {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, "invalid action parameter")
		return
	}

	uid, ok := ctx.Get("user_id")
	if !ok {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	var form ActiveShiftForm
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	form.ShiftID = id
	form.Action = actionParams
	form.UserID = int64(uid.(float64))

	if err := handler.service.ActiveShiftAction(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func NewShiftHandler(service IShiftService, router *gin.RouterGroup) {
	handler := shiftHandler{service: service}
	authz := router.Group("/shifts")
	authz.Use(utils.AuthZ([]string{"admin"}))
	{
		authz.POST("/:id/:action", handler.active)
	}
}
