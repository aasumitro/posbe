package order

import (
	"net/http"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	service IShiftService
}

func (handler shiftHandler) show(ctx *gin.Context) {
	data, err := handler.service.CurrentActiveShift(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}
	utils.NewHTTPRespond(ctx, http.StatusOK, data)
}

func (handler shiftHandler) action(ctx *gin.Context) {
	var form ActiveShiftForm

	uid, ok := ctx.Get("user_id")
	if !ok {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, "invalid user id")
		return
	}
	form.UserID = int64(uid.(float64))

	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if val := form.Validate(ctx); val != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, val)
		return
	}

	if err := handler.service.ActiveShiftAction(ctx, &form); err != nil {
		utils.NewHTTPRespond(ctx, err.Code, err.Message)
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, nil)
}

func NewShiftHandler(service IShiftService, router *gin.RouterGroup) {
	handler := shiftHandler{service: service}
	shift := router.Group("shifts")
	shift.GET(utils.EmptyPath, handler.show)
	shift.Use(utils.AuthZ([]string{"admin", "cashier"}))
	{
		shift.POST(utils.EmptyPath, handler.action)
	}
}
