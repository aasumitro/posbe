package sql

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type ProductVariantSQLRepository struct {
	Db *sql.DB
}

func (repo ProductVariantSQLRepository) All(_ context.Context) (data []*model.ProductVariant, err error) {
	panic("not implement")
}

func (repo ProductVariantSQLRepository) Find(ctx context.Context, _ model.FindWith, val any) (data *model.ProductVariant, err error) {
	q := "SELECT * FROM product_variants WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &model.ProductVariant{}
	if err := row.Scan(
		&data.ID, &data.ProductID, &data.UnitID,
		&data.UnitSize, &data.Type, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductVariantSQLRepository) Create(ctx context.Context, params *model.ProductVariant) (data *model.ProductVariant, err error) {
	q := "INSERT INTO product_variants (product_id, unit_id, unit_size, type, name, description, price) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.ProductID, params.UnitID, params.UnitSize, params.Type, params.Name, params.Description, params.Price)

	data = &model.ProductVariant{}
	if err := row.Scan(
		&data.ID, &data.ProductID, &data.UnitID,
		&data.UnitSize, &data.Type, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductVariantSQLRepository) Update(ctx context.Context, params *model.ProductVariant) (data *model.ProductVariant, err error) {
	q := "UPDATE product_variants SET product_id = $1, unit_id = $2, unit_size = $3, type = $4, name = $5, description = $6, price = $7 WHERE id = $8 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.ProductID, params.UnitID, params.UnitSize, params.Type, params.Name, params.Description, params.Price, params.ID)

	data = &model.ProductVariant{}
	if err := row.Scan(
		&data.ID, &data.ProductID, &data.UnitID,
		&data.UnitSize, &data.Type, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductVariantSQLRepository) Delete(ctx context.Context, params *model.ProductVariant) error {
	q := "DELETE FROM product_variants WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewProductVariantSQLRepository() model.ICRUDRepository[model.ProductVariant] {
	return &ProductVariantSQLRepository{Db: config.PostgresPool}
}
