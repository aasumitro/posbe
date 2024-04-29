package sql

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type SubcategorySQLRepository struct {
	Db *sql.DB
}

func (repo SubcategorySQLRepository) All(ctx context.Context) (data []*model.Subcategory, err error) {
	q := "SELECT * FROM subcategories"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var subcategories model.Subcategory

		if err := rows.Scan(
			&subcategories.ID,
			&subcategories.CategoryID,
			&subcategories.Name,
		); err != nil {
			return nil, err
		}

		data = append(data, &subcategories)
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Find(ctx context.Context, _ model.FindWith, val any) (data *model.Subcategory, err error) {
	q := "SELECT * FROM subcategories WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &model.Subcategory{}
	if err := row.Scan(
		&data.ID,
		&data.CategoryID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Create(ctx context.Context, params *model.Subcategory) (data *model.Subcategory, err error) {
	q := "INSERT INTO subcategories (category_id, name) VALUES ($1, $2) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryID, params.Name)

	data = &model.Subcategory{}
	if err := row.Scan(
		&data.ID,
		&data.CategoryID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Update(ctx context.Context, params *model.Subcategory) (data *model.Subcategory, err error) {
	q := "UPDATE subcategories SET category_id = $1, name = $2 WHERE id = $3 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryID, params.Name, params.ID)

	data = &model.Subcategory{}
	if err := row.Scan(
		&data.ID,
		&data.CategoryID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Delete(ctx context.Context, params *model.Subcategory) error {
	q := "DELETE FROM subcategories WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewSubcategorySQLRepository() model.ICRUDRepository[model.Subcategory] {
	return &SubcategorySQLRepository{Db: config.PostgresPool}
}
