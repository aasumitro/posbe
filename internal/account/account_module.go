package account

import (
	"context"
	"github.com/aasumitro/posbe/domain"
	repository "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/account/service"
	"github.com/aasumitro/posbe/pkg/config"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func InitAccountModule(ctx context.Context, config *config.Config, router *gin.Engine) {
	userRepository := repository.NewUserSQlRepository(config.GetDbConn())
	roleRepository := repository.NewRoleSQlRepository(config.GetDbConn())
	accountService := service.NewAccountService(ctx, roleRepository, userRepository)

	//handler.NewAuthHandler(ctx, accountService, router)
	//handler.NewRoleHandler(ctx, accountService, router)
	//handler.NewUserHandler(ctx, accountService, router)

	router.GET("/users", func(c *gin.Context) {
		users, err := userRepository.All(ctx)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)
		users, err := userRepository.Find(ctx, id)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	router.POST("/users", func(c *gin.Context) {
		var form domain.User
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		user, err := userRepository.Create(ctx, &form)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}
		user.Password = ""
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)

		var form domain.User
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		form.ID = id

		user, err := userRepository.Update(ctx, &form)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}
		user.Password = ""
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)
		user := domain.User{ID: id}
		if err := userRepository.Delete(ctx, &user); err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": "SUCCESS"})
	})

	router.GET("/roles", func(ctx *gin.Context) {
		roles, err := accountService.RoleList()
		if err != nil {
			utils.NewHttpRespond(ctx, err.Code, err.Message)
			return
		}

		utils.NewHttpRespond(ctx, http.StatusOK, roles)
		return
	})

	router.GET("/roles/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)
		role, err := roleRepository.Find(ctx, id)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"data": role})
	})

	router.POST("/roles", func(c *gin.Context) {
		var form domain.Role
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		role, err := roleRepository.Create(ctx, &form)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": role})
	})

	router.PUT("/roles/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)

		var form domain.Role
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		form.ID = id

		role, err := roleRepository.Update(ctx, &form)
		if err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": role})
	})

	router.DELETE("/roles/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)
		role := domain.Role{ID: id}
		if err := roleRepository.Delete(ctx, &role); err != nil {
			log.Panicf("ERROR_REPO: %s", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": "SUCCESS"})
	})

	router.DELETE("/test/:id", func(c *gin.Context) {
		idParams := c.Param("id")
		id, _ := strconv.Atoi(idParams)
		role := domain.Role{ID: id}
		if err := accountService.DeleteRole(&role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "SUCCESS"})
	})
}
