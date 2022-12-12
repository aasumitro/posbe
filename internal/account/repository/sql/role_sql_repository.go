package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/config"
)

type RoleSQLRepository struct {
	Db *sql.DB
}

func (repo RoleSQLRepository) All(ctx context.Context) (roles []*domain.Role, err error) {
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var role domain.Role

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

func (repo RoleSQLRepository) Find(ctx context.Context, _ domain.FindWith, val any) (role *domain.Role, err error) {
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "WHERE roles.id = $1 GROUP BY roles.id LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	role = &domain.Role{}
	if err := row.Scan(
		&role.ID, &role.Name,
		&role.Description, &role.Usage,
	); err != nil {
		return nil, err
	}

	return role, nil
}

func (repo RoleSQLRepository) Create(ctx context.Context, params *domain.Role) (role *domain.Role, err error) {
	q := "INSERT INTO roles (name, description) values ($1, $2) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, params.Description)

	role = &domain.Role{}
	if err := row.Scan(&role.ID, &role.Name, &role.Description); err != nil {
		return nil, err
	}

	return role, nil
}

func (repo RoleSQLRepository) Update(ctx context.Context, params *domain.Role) (role *domain.Role, err error) {
	q := "UPDATE roles SET name = $1, description = $2 WHERE id = $3 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, params.Description, params.ID)

	role = &domain.Role{}
	if err := row.Scan(&role.ID, &role.Name, &role.Description); err != nil {
		return nil, err
	}

	return role, nil
}

func (repo RoleSQLRepository) Delete(ctx context.Context, params *domain.Role) error {
	q := "DELETE FROM roles WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewRoleSQlRepository() domain.ICRUDRepository[domain.Role] {
	return &RoleSQLRepository{Db: config.DbPool}
}
