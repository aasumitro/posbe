package sql

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type AddonSQLRepository struct {
	Db *sql.DB
}

func (repo AddonSQLRepository) All(ctx context.Context) (data []*model.Addon, err error) {
	q := "SELECT * FROM addons"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var addon model.Addon

		if err := rows.Scan(
			&addon.ID, &addon.Name,
			&addon.Description, &addon.Price,
		); err != nil {
			return nil, err
		}

		data = append(data, &addon)
	}

	return data, nil
}

func (repo AddonSQLRepository) Find(ctx context.Context, _ model.FindWith, val any) (data *model.Addon, err error) {
	q := "SELECT * FROM addons WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &model.Addon{}
	if err := row.Scan(
		&data.ID, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo AddonSQLRepository) Create(ctx context.Context, params *model.Addon) (data *model.Addon, err error) {
	q := "INSERT INTO addons (name, description, price) VALUES ($1, $2, $3) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, params.Description, params.Price)

	data = &model.Addon{}
	if err := row.Scan(
		&data.ID, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo AddonSQLRepository) Update(ctx context.Context, params *model.Addon) (data *model.Addon, err error) {
	q := "UPDATE addons SET name = $1, description = $2, price = $3 WHERE id = $4 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, params.Description, params.Price, params.ID)

	data = &model.Addon{}
	if err := row.Scan(
		&data.ID, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo AddonSQLRepository) Delete(ctx context.Context, params *model.Addon) error {
	q := "DELETE FROM addons WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewAddonSQLRepository() model.ICRUDRepository[model.Addon] {
	return &AddonSQLRepository{Db: config.PostgresPool}
}
