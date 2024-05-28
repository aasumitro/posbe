package sql

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type CategorySQLRepository struct {
	Db *sql.DB
}

func (repo CategorySQLRepository) All(ctx context.Context) (data []*model.Category, err error) {
	q := "SELECT * FROM categories"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var category model.Category

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

func (repo CategorySQLRepository) Find(ctx context.Context, _ model.FindWith, val any) (data *model.Category, err error) {
	q := "SELECT * FROM categories WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &model.Category{}
	if err := row.Scan(
		&data.ID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo CategorySQLRepository) Create(ctx context.Context, params *model.Category) (data *model.Category, err error) {
	q := "INSERT INTO categories (name) VALUES ($1) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name)

	data = &model.Category{}
	if err := row.Scan(
		&data.ID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo CategorySQLRepository) Update(ctx context.Context, params *model.Category) (data *model.Category, err error) {
	q := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, params.ID)

	data = &model.Category{}
	if err := row.Scan(
		&data.ID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo CategorySQLRepository) Delete(ctx context.Context, params *model.Category) error {
	q := "DELETE FROM categories WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewCategorySQLRepository() model.ICRUDRepository[model.Category] {
	return &CategorySQLRepository{Db: config.PostgresPool}
}
