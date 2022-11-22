package domain

import "context"

type (
	// User for response
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}

	// UserEntity for request
	UserEntity struct {
		ID       string
		Username string
		Password string
	}

	// UserRepository Contract
	UserRepository interface {
		by()
		with()
		all()
		find()
		create()
		update()
		delete()
	}

	// UserHandler Contract
	UserHandler interface {
		fetch()
		show()
		store()
		update()
		destroy()
	}
)

type (
	// Role for response & request
	Role struct {
		ID          string
		Name        string
		Description string
	}

	// RoleRepository Contract
	RoleRepository interface {
		All(ctx context.Context) (roles []Role, err error)
		Find(ctx context.Context, where []string) (role *Role, err error)
		//Create(ctx context.Context)
		//Update(ctx context.Context)
		//Delete(ctx context.Context)
	}

	// RoleHandler Contract
	RoleHandler interface {
		Fetch()
		Show()
		Store()
		Update()
		Destroy()
	}
)

type (
	// Account for response
	Account struct {
	}

	// AccountService contract
	AccountService interface {
	}
)
