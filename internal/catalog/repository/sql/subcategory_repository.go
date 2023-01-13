package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
)

type SubcategorySQLRepository struct {
	Db *sql.DB
}

func (repo SubcategorySQLRepository) All(ctx context.Context) (data []*domain.Subcategory, err error) {
	q := "SELECT * FROM subcategories"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var subcategories domain.Subcategory

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

func (repo SubcategorySQLRepository) Find(ctx context.Context, _ domain.FindWith, val any) (data *domain.Subcategory, err error) {
	q := "SELECT * FROM subcategories WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &domain.Subcategory{}
	if err := row.Scan(
		&data.ID,
		&data.CategoryID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Create(ctx context.Context, params *domain.Subcategory) (data *domain.Subcategory, err error) {
	q := "INSERT INTO subcategories (category_id, name) VALUES ($1, $2) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryID, params.Name)

	data = &domain.Subcategory{}
	if err := row.Scan(
		&data.ID,
		&data.CategoryID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Update(ctx context.Context, params *domain.Subcategory) (data *domain.Subcategory, err error) {
	q := "UPDATE subcategories SET category_id = $1, name = $2 WHERE id = $3 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryID, params.Name, params.ID)

	data = &domain.Subcategory{}
	if err := row.Scan(
		&data.ID,
		&data.CategoryID,
		&data.Name,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo SubcategorySQLRepository) Delete(ctx context.Context, params *domain.Subcategory) error {
	q := "DELETE FROM subcategories WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewSubcategorySQLRepository() domain.ICRUDRepository[domain.Subcategory] {
	return &SubcategorySQLRepository{Db: configs.DbPool}
}
