package provider

import (
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	repo "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/internal/account/service"
	"sync"
)

var (
	roleRepo     *repo.RoleSQLRepository
	roleRepoOnce sync.Once

	userRepo     *repo.UserSQLRepository
	userRepoOnce sync.Once

	accountSvc     *service.AccountService
	accountSvcOnce sync.Once
)

func ProvideRoleRepository(db *sql.DB) *repo.RoleSQLRepository {
	roleRepoOnce.Do(func() {
		roleRepo = &repo.RoleSQLRepository{
			Db: db,
		}
	})

	return roleRepo
}

func ProvideUserRepository(db *sql.DB) *repo.UserSQLRepository {
	userRepoOnce.Do(func() {
		userRepo = &repo.UserSQLRepository{
			Db: db,
		}
	})

	return userRepo
}

func ProvideAccountService(role domain.RoleRepository, user domain.UserRepository) *service.AccountService {
	accountSvcOnce.Do(func() {
		accountSvc = &service.AccountService{
			RoleRepo: role,
			UserRepo: user,
		}
	})

	return accountSvc
}
