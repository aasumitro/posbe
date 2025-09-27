package store

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

func New(router *gin.RouterGroup) {
	// store setting
	settingRepo := NewSettingRepository(config.PgxPool)
	settingSvc := NewSettingService(settingRepo)
	// store shift
	shiftRepo := NewShiftRepository(config.PgxPool)
	shiftSvc := NewShiftService(shiftRepo)
	// seating (floor and table)
	seatingRepo := NewSeatingRepository(config.PgxPool)
	seatingSvc := NewSeatingService(seatingRepo)
	// register handler
	authn := router.Group("store")
	authn.Use(utils.AuthN())
	{
		NewShiftHandler(shiftSvc, authn)
		NewSettingHandler(settingSvc, authn)
		NewSeatingHandler(seatingSvc, authn)
	}
}
