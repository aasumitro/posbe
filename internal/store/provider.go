package store

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

func New(router *gin.RouterGroup) {
	authn := router.Group("store").Use(utils.AuthN())
	// store setting
	settingRepo := NewSettingRepository(config.PgxPool)
	settingSvc := NewSettingService(settingRepo)
	NewSettingHandler(settingSvc, authn)
	// store shift
	shiftRepo := NewShiftRepository(config.PgxPool)
	shiftSvc := NewShiftService(shiftRepo)
	NewShiftHandler(shiftSvc, authn)
	// seating (floor and table)
	seatingRepo := NewSeatingRepository(config.PgxPool)
	seatingSvc := NewSeatingService(seatingRepo)
	NewSeatingHandler(seatingSvc, authn)
}
