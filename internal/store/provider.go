package store

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

func NewStoreModuleProvider(router *gin.RouterGroup) {
	authn := router.Group("store").Use(utils.AuthN())
	settingRepo := NewSettingRepository(config.PgxPool)
	settingSvc := NewSettingService(settingRepo)
	NewSettingHandler(settingSvc, authn)
	shiftRepo := NewShiftRepository(config.PgxPool)
	shiftSvc := NewShiftService(shiftRepo)
	NewShiftHandler(shiftSvc, authn)
}
