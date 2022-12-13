package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/config"
)

type ProductSQLRepository struct {
	Db *sql.DB
}

func (repo ProductSQLRepository) Search(ctx context.Context, keys []domain.FindWith, values []any) (data []*domain.Product, err error) {
	q := "SELECT * FROM products "
	var whereClause string
	if len(keys) > 0 {
		whereClause += "WHERE "

		for i, key := range keys {
			switch key {
			case domain.FindWithSKU:
				data := values[i].(string)
				whereClause += fmt.Sprintf("sku = '%s' ", data)
			case domain.FindWithCategoryID:
				data := values[i].(int)
				whereClause += fmt.Sprintf("category_id = %d ", data)
			case domain.FindWithSubcategoryID:
				data := values[i].(int)
				whereClause += fmt.Sprintf("subcategory_id = %d ", data)
			case domain.FindWithPriceInRange:
				data := values[i].([]float32)
				whereClause += fmt.Sprintf("price BETWEEN %f AND %f ", data[0], data[1])
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
		var product domain.Product

		if err := rows.Scan(
			&product.ID, &product.CategoryId, &product.SubcategoryId,
			&product.Sku, &product.Image, &product.Gallery, &product.Name,
			&product.Description, &product.Price,
		); err != nil {
			return nil, err
		}

		data = append(data, &product)
	}

	return data, nil
}

func (repo ProductSQLRepository) All(ctx context.Context) (data []*domain.Product, err error) {
	q := "SELECT * FROM products"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(
			&product.ID, &product.CategoryId, &product.SubcategoryId,
			&product.Sku, &product.Image, &product.Gallery, &product.Name,
			&product.Description, &product.Price,
		); err != nil {
			return nil, err
		}

		data = append(data, &product)
	}

	return data, nil
}

func (repo ProductSQLRepository) Find(ctx context.Context, _ domain.FindWith, val any) (data *domain.Product, err error) {
	q := "SELECT * FROM products WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &domain.Product{}
	if err := row.Scan(
		&data.ID, &data.CategoryId, &data.SubcategoryId,
		&data.Sku, &data.Image, &data.Gallery, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductSQLRepository) Create(ctx context.Context, params *domain.Product) (data *domain.Product, err error) {
	q := "INSERT INTO products "
	q += "(category_id, subcategory_id, sku, image, gallery, name, description, price) "
	q += "VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryId, params.SubcategoryId,
		params.Sku, params.Image, params.Gallery, params.Name,
		params.Description, params.Price)

	data = &domain.Product{}
	if err := row.Scan(
		&data.ID, &data.CategoryId, &data.SubcategoryId,
		&data.Sku, &data.Image, &data.Gallery, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductSQLRepository) Update(ctx context.Context, params *domain.Product) (data *domain.Product, err error) {
	q := "UPDATE products SET category_id = $1, subcategory_id = $2, sku = $3, image = $4, "
	q += "gallery = $5, name = $6, description = $7, price = $8 WHERE id = $9 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.CategoryId, params.SubcategoryId,
		params.Sku, params.Image, params.Gallery, params.Name,
		params.Description, params.Price, params.ID)

	data = &domain.Product{}
	if err := row.Scan(
		&data.ID, &data.CategoryId, &data.SubcategoryId,
		&data.Sku, &data.Image, &data.Gallery, &data.Name,
		&data.Description, &data.Price,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo ProductSQLRepository) Delete(ctx context.Context, params *domain.Product) error {
	q := "DELETE FROM products WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewProductSQLRepository() domain.ICRUDWithSearchRepository[domain.Product] {
	return &ProductSQLRepository{Db: config.DbPool}
}
