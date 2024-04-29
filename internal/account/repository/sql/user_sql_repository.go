package sql

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type UserSQLRepository struct {
	Db *sql.DB
}

func (repo UserSQLRepository) All(ctx context.Context) (users []*model.User, err error) {
	q := `SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, 
		r.id as role_id, r.name as role_name, r.description FROM users as u 
		JOIN roles as r ON r.id = u.role_id`
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID, &user.RoleID,
			&user.Name, &user.Username,
			&user.Email, &user.Phone,
			&user.Role.ID, &user.Role.Name,
			&user.Role.Description,
		); err != nil {
			return nil, err
		}
		users = append(users, &model.User{
			ID: user.ID, Name: user.Name, Username: user.Username,
			Email: user.Email, Phone: user.Phone, Role: model.Role{
				ID: user.Role.ID, Name: user.Role.Name,
				Description: user.Role.Description,
			},
		})
	}
	return users, nil
}

func (repo UserSQLRepository) Find(ctx context.Context, key model.FindWith, val any) (user *model.User, err error) {
	q := `
		SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, 
		r.id as role_id, r.name as role_name, r.description FROM users as u 
		JOIN roles as r ON r.id = u.role_id WHERE
	`
	//goland:noinspection GoSwitchMissingCasesForIotaConsts
	switch key {
	case model.FindWithID:
		q += "u.id = $1"
	case model.FindWithName:
		q += "u.name LIKE %$1%"
	case model.FindWithUsername:
		q += "u.username = $1"
	case model.FindWithEmail:
		q += "u.email = $1"
	case model.FindWithPhone:
		q += "u.phone = $1"
	}
	q += " LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)
	return scanData(row)
}

func (repo UserSQLRepository) Create(ctx context.Context, params *model.User) (user *model.User, err error) {
	q := "WITH u AS (INSERT INTO users(role_id, name, username, email, phone, password) "
	q += "values ($1, $2, $3, $4, $5, $6) RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"
	row := repo.Db.QueryRowContext(
		ctx, q, params.RoleID, params.Name,
		params.Username, params.Email, params.Phone,
		params.Password)
	return scanData(row)
}

func (repo UserSQLRepository) Update(ctx context.Context, params *model.User) (user *model.User, err error) {
	q := "WITH u AS (UPDATE users SET role_id = $1, name = $2, username = $3, email = $4, "
	q += "phone = $5, password = $6 WHERE id = $7 RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"
	row := repo.Db.QueryRowContext(
		ctx, q, params.RoleID, params.Name, params.Username,
		params.Email, params.Phone, params.Password, params.ID)
	return scanData(row)
}

func (repo UserSQLRepository) Delete(ctx context.Context, params *model.User) error {
	q := "DELETE FROM users WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func scanData(row *sql.Row) (data *model.User, err error) {
	var user model.User
	if err := row.Scan(
		&user.ID, &user.RoleID, &user.Name,
		&user.Username, &user.Email, &user.Phone, &user.Password,
		&user.Role.ID, &user.Role.Name, &user.Role.Description,
	); err != nil {
		return nil, err
	}
	return &model.User{
		ID: user.ID, Name: user.Name, Username: user.Username, Password: user.Password,
		Email: user.Email, Phone: user.Phone, Role: model.Role{
			ID: user.Role.ID, Name: user.Role.Name,
			Description: user.Role.Description,
		},
	}, nil
}

func NewUserSQLRepository() model.ICRUDRepository[model.User] {
	return &UserSQLRepository{Db: config.PostgresPool}
}
