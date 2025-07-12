package sql

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
)

type RoleSQLRepository struct {
	Db *sql.DB
}

func (repo RoleSQLRepository) All(ctx context.Context) (roles []*model.Role, err error) {
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(
			&role.ID, &role.Name,
			&role.Description, &role.Usage,
		); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}

func (repo RoleSQLRepository) Find(_ context.Context, _ model.FindWith, _ any) (*model.Role, error) {
	panic("implement me")
}

func (repo RoleSQLRepository) Create(_ context.Context, _ *model.Role) (*model.Role, error) {
	panic("implement me")
}

func (repo RoleSQLRepository) Update(_ context.Context, _ *model.Role) (*model.Role, error) {
	panic("implement me")
}

func (repo RoleSQLRepository) Delete(_ context.Context, _ *model.Role) error {
	panic("implement me")
}

func NewRoleSQLRepository() model.ICRUDRepository[model.Role] {
	return &RoleSQLRepository{Db: config.PostgresPool}
}
