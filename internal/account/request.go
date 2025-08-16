package account

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

////////////////// ======= REQUEST Query ======= //////////////////

type UserQuery struct {
	ID     int    `json:"user_id" form:"user_id"`
	RoleID int    `json:"role_id" form:"role_id"`
	Sort   string `json:"sort" form:"sort"`
}

////////////////// ======= REQUEST BODY ======= //////////////////

type LoginForm struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (f *LoginForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		"Username": g.R("username").Required(),
		"Password": g.R("password").Required(),
	}).Validate(ctx, f)
}

type UpdateProfileForm struct{}

func (f *UpdateProfileForm) Validate(ctx *gin.Context) interface{} {
	return nil
}

type UpdatePasswordForm struct {
	ID          int    `json:"-" form:"-"`
	OldPassword string `json:"old_pwd" form:"old_pwd"`
	NewPassword string `json:"new_pwd" form:"new_pwd"`
}

func (f *UpdatePasswordForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		"OldPassword": g.R("old_pwd").Required(),
		"NewPassword": g.R("new_pwd").Required().Password(),
	}).Validate(ctx, f)
}

type NewUserForm struct{}

func (f *NewUserForm) Validate(ctx *gin.Context) interface{} {
	return nil
}

type UpdateUserForm struct{}

func (f *UpdateUserForm) Validate(ctx *gin.Context) interface{} {
	return nil
}
