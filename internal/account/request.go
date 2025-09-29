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

type UpdatePasswordForm struct {
	ID          int64  `json:"-" form:"-"`
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

func (f *UpdatePasswordForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		"Password":    g.R("password").Required(),
		"NewPassword": g.R("new_password").Required().Password(),
	}).Validate(ctx, f)
}

type NewUserForm struct {
	RoleID   int    `json:"role_id" form:"role_id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (f *NewUserForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		"RoleID":   g.R("role_id").Required(),
		"Name":     g.R("name").Required(),
		"Username": g.R("username").Required(),
		"Email":    g.R("email").Required().Email(),
		"Password": g.R("password").Required().Password(),
	}).Validate(ctx, f)
}

type UpdateUserForm struct {
	ID       int64  `json:"-" form:"-"`
	RoleID   int    `json:"role_id" form:"role_id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (f *UpdateUserForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	return g.ComplexValidator(galidator.Rules{
		"RoleID":   g.R("role_id").Optional(),
		"Name":     g.R("name").Optional(),
		"Username": g.R("username").Optional(),
		"Email":    g.R("email").Optional().Email(),
		"Password": g.R("password").Optional().Password(),
	}).Validate(ctx, f)
}
