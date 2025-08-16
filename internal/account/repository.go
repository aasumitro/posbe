package account

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func (repository accountRepository) GetAllRoles(ctx context.Context) (roles []*model.Role, err error) {
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"

	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var role model.Role
	if _, err = pgx.ForEachRow(rows, []any{
		&role.ID,
		&role.Name,
		&role.Description,
		&role.Usage,
	}, func() error {
		r := role
		roles = append(roles, &r)
		return nil
	}); err != nil {
		return nil, err
	}

	return roles, nil
}

func (repository accountRepository) GetAllUsers(ctx context.Context) (users []*model.User, err error) {
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, r.id as role_id, "
	q += "r.name as role_name, r.description, u.created_at FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id"

	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	if _, err = pgx.ForEachRow(rows, []any{
		&user.ID,
		&user.RoleID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Role.ID,
		&user.Role.Name,
		&user.Role.Description,
		&user.CreatedAt,
	}, func() error {
		u := user
		users = append(users, &u)
		return nil
	}); err != nil {
		return nil, err
	}

	return users, nil
}

func (repository accountRepository) FindUserBy(ctx context.Context, key model.FindWith, val any) (*model.User, error) {
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id WHERE "
	// notes: all the value is from service not from user input so it's safe!
	switch key {
	case model.FindWithID:
		q += "u.id = $1"
	case model.FindWithUsername:
		q += "u.username = $1"
	case model.FindWithEmail:
		q += "u.email = $1"
	default:
		return nil, errors.New("invalid search key")
	}
	q += " LIMIT 1"

	var user model.User
	if err := repository.db.QueryRow(ctx, q, val).Scan(
		&user.ID,
		&user.RoleID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role.ID,
		&user.Role.Name,
		&user.Role.Description,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository accountRepository) InsertUser(ctx context.Context, user model.User) (*model.User, error) {
	q := "WITH u AS (INSERT INTO users(role_id, name, username, email, password, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6) RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"

	var inserted model.User
	if err := repository.db.QueryRow(
		ctx, q,
		user.RoleID,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		time.Now().Unix(),
	).Scan(
		&inserted.ID,
		&inserted.RoleID,
		&inserted.Name,
		&inserted.Username,
		&inserted.Email,
		&inserted.Role.ID,
		&inserted.Role.Name,
		&inserted.Role.Description,
	); err != nil {
		return nil, err
	}

	return &inserted, nil
}

func (repository accountRepository) UpdateUserByID(ctx context.Context, user model.User) (*model.User, error) {
	var setClauses []string
	var args []any
	argPos := 1

	// Build SET clause dynamically only for fields that are non-zero/non-empty
	if user.RoleID != 0 {
		setClauses = append(setClauses, "role_id = $"+strconv.Itoa(argPos))
		args = append(args, user.RoleID)
		argPos++
	}
	if user.Name != "" {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, user.Name)
		argPos++
	}
	if user.Username != "" {
		setClauses = append(setClauses, "username = $"+strconv.Itoa(argPos))
		args = append(args, user.Username)
		argPos++
	}
	if user.Email != "" {
		setClauses = append(setClauses, "email = $"+strconv.Itoa(argPos))
		args = append(args, user.Email)
		argPos++
	}
	if user.Password != "" {
		setClauses = append(setClauses, "password = $"+strconv.Itoa(argPos))
		args = append(args, user.Password)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	// Add WHERE clause
	if user.ID == 0 {
		return nil, errors.New("missing user ID for update")
	}
	args = append(args, user.ID)

	q := fmt.Sprintf(`
		WITH u AS (
			UPDATE users SET %s
			WHERE id = $%d
			RETURNING *
		)
		SELECT u.id, u.role_id, u.name, u.username, u.email, u.password,
		       r.id as role_id, r.name as role_name, r.description
		FROM u
		JOIN roles as r ON r.id = u.role_id
	`, strings.Join(setClauses, ", "), argPos)

	var updated model.User
	if err := repository.db.QueryRow(ctx, q, args...).Scan(
		&updated.ID,
		&updated.RoleID,
		&updated.Name,
		&updated.Username,
		&updated.Email,
		&updated.Password,
		&updated.Role.ID,
		&updated.Role.Name,
		&updated.Role.Description,
	); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (repository accountRepository) DeleteUserByID(ctx context.Context, user model.User) error {
	q := "DELETE FROM users WHERE id = $1"
	_, err := repository.db.Exec(ctx, q, user.ID)
	return err
}

func NewAccountRepository(db *pgxpool.Pool) IAccountRepository {
	return &accountRepository{db: db}
}
