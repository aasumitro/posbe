package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type ProductSQLRepository struct {
	Db *sql.DB
}

func (repo ProductSQLRepository) Search(ctx context.Context, keys []model.FindWith, values []any) (data []*model.Product, err error) {
	q := "SELECT * FROM products "
	var whereClause string
	if len(keys) > 0 {
		whereClause += "WHERE "

		for i, key := range keys {
			switch key {
			case model.FindWithSKU:
				data := values[i].(string)
				whereClause += fmt.Sprintf("sku = '%s' ", data)
			case model.FindWithCategoryID:
				data := values[i].(int)
				whereClause += fmt.Sprintf("category_id = %d ", data)
			case model.FindWithSubcategoryID:
				data := values[i].(int)
				whereClause += fmt.Sprintf("subcategory_id = %d ", data)
			case model.FindWithPriceInRange:
				data := values[i].([]float32)
				whereClause += fmt.Sprintf("price BETWEEN %f AND %f ", data[0], data[1])
			default:
				panic("unhandled default case")
			}

			if i != (len(keys) - 1) {
				whereClause += "AND "
			}
		}

		q += whereClause
	}

	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var product model.Product

		if err := rows.Scan(
			&product.ID, &product.CategoryID, &product.SubcategoryID,
			&product.Sku, &product.Image, &product.Gallery, &product.Name,
			&product.Description, &product.Price,
		); err != nil {
			return nil, err
		}

		data = append(data, &product)
	}

	return data, nil
}

func (repo ProductSQLRepository) All(ctx context.Context) (data []*model.Product, err error) {
	q := "SELECT * FROM products"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var product model.Product

		if err := rows.Scan(
			&product.ID, &product.CategoryID, &product.SubcategoryID,
			&product.Sku, &product.Image, &product.Gallery, &product.Name,
			&product.Description, &product.Price,
		); err != nil {
			return nil, err
		}

		data = append(data, &product)
	}

	return data, nil
}

func (repo ProductSQLRepository) Find(ctx context.Context, _ model.FindWith, val any) (data *model.Product, err error) {
	q := "SELECT * FROM products WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &model.Product{}
	if err := row.Scan(
		&data.ID, &data.CategoryID, &data.SubcategoryID,
		&data.Sku, &data.Image, &data.Gallery, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductSQLRepository) Create(ctx context.Context, params *model.Product) (data *model.Product, err error) {
	q := "INSERT INTO products "
	q += "(category_id, subcategory_id, sku, image, gallery, name, description, price) "
	q += "VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryID, params.SubcategoryID,
		params.Sku, params.Image, params.Gallery, params.Name,
		params.Description, params.Price)

	data = &model.Product{}
	if err := row.Scan(
		&data.ID, &data.CategoryID, &data.SubcategoryID,
		&data.Sku, &data.Image, &data.Gallery, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductSQLRepository) Update(ctx context.Context, params *model.Product) (data *model.Product, err error) {
	q := "UPDATE products SET category_id = $1, subcategory_id = $2, sku = $3, image = $4, "
	q += "gallery = $5, name = $6, description = $7, price = $8 WHERE id = $9 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryID, params.SubcategoryID,
		params.Sku, params.Image, params.Gallery, params.Name,
		params.Description, params.Price, params.ID)

	data = &model.Product{}
	if err := row.Scan(
		&data.ID, &data.CategoryID, &data.SubcategoryID,
		&data.Sku, &data.Image, &data.Gallery, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductSQLRepository) Delete(ctx context.Context, params *model.Product) error {
	q := "DELETE FROM products WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewProductSQLRepository() model.ICRUDWithSearchRepository[model.Product] {
	return &ProductSQLRepository{Db: config.PostgresPool}
}
