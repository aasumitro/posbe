package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
)

type CategorySQLRepository struct {
	Db *sql.DB
}

func (repo CategorySQLRepository) All(ctx context.Context) (data []*domain.Category, err error) {
	q := "SELECT * FROM categories"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var category domain.Category

		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		data = append(data, &category)
	}

	return data, nil
}

func (repo CategorySQLRepository) Find(ctx context.Context, _ domain.FindWith, val any) (data *domain.Category, err error) {
	q := "SELECT * FROM categories WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &domain.Category{}
	if err := row.Scan(
		&data.ID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo CategorySQLRepository) Create(ctx context.Context, params *domain.Category) (data *domain.Category, err error) {
	q := "INSERT INTO categories (name) VALUES ($1) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name)

	data = &domain.Category{}
	if err := row.Scan(
		&data.ID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo CategorySQLRepository) Update(ctx context.Context, params *domain.Category) (data *domain.Category, err error) {
	q := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, params.ID)

	data = &domain.Category{}
	if err := row.Scan(
		&data.ID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo CategorySQLRepository) Delete(ctx context.Context, params *domain.Category) error {
	q := "DELETE FROM categories WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewCategorySQLRepository() domain.ICRUDRepository[domain.Category] {
	return &CategorySQLRepository{Db: configs.DbPool}
}
