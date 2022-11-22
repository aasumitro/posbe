package service

import "github.com/aasumitro/posbe/domain"

type AccountService struct {
	RoleRepo domain.RoleRepository
	UserRepo domain.UserRepository
}
