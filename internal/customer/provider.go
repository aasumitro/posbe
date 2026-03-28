package customer

import (
	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
)

func New(router *gin.RouterGroup) {
	repo := NewRepository(config.PgxPool)
	svc := NewService(repo)
	authn := router.Group(utils.EmptyPath)
	authn.Use(utils.AuthN())
	{
		NewHandler(svc, authn)
	}
}
