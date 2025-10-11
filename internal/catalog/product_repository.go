package catalog

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
)

type productRepository struct {
	db model.IPgxPool
}

func (repository productRepository) GetAllAddon(ctx context.Context) ([]*model.Addon, error) {
	q := `
		SELECT a.id, a.price, a.name, a.description,
		       (SELECT COUNT(*) FROM order_product_addons AS opa WHERE opa.addon_id = a.id) AS usage
		FROM product_addons AS a
	`
	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addons []*model.Addon
	var addon model.Addon
	if _, err = pgx.ForEachRow(rows, []any{
		&addon.ID, &addon.Price, &addon.Name,
		&addon.Description, &addon.Usage,
	}, func() error {
		a := addon
		addons = append(addons, &a)
		return nil
	}); err != nil {
		return nil, err
	}

	return addons, nil
}

func (repository productRepository) CreateAddon(ctx context.Context, form *AddonForm) error {
	q := "INSERT INTO product_addons (price, name, description) VALUES ($1, $2, $3)"
	_, err := repository.db.Exec(ctx, q, form.Price, form.Name, form.Description)
	return err
}

func (repository productRepository) UpdateAddon(ctx context.Context, form *AddonForm) error {
	if form.ID == 0 {
		return errors.New("missing ID for update")
	}

	var setClauses []string
	var args []any
	argPos := 1

	if form.Price != nil {
		setClauses = append(setClauses, "price = $"+strconv.Itoa(argPos))
		args = append(args, form.Price)
		argPos++
	}
	if form.Name != "" {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, form.Name)
		argPos++
	}
	if form.Description != "" {
		setClauses = append(setClauses, "description = $"+strconv.Itoa(argPos))
		args = append(args, form.Description)
		argPos++
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, form.ID)

	q := fmt.Sprintf(`UPDATE product_addons SET %s WHERE id = $%d`,
		strings.Join(setClauses, ", "), argPos)
	_, err := repository.db.Exec(ctx, q, args...)
	return err
}

func (repository productRepository) DeleteAddon(ctx context.Context, id int64) error {
	q := "DELETE FROM product_addons WHERE id = $1"
	_, err := repository.db.Exec(ctx, q, id)
	return err
}

func (repository productRepository) GetAllProduct(ctx context.Context, query *ProductQuery) ([]*model.Product, error) {
	q := `
		SELECT 
		  p.id, p.category_id, p.subcategory_id, p.sku, 
		  p.image, p.name, p.description, p.status,
			
          -- embed category as JSON
		  json_build_object(
		    'id', c.id,
		    'name', c.name
		  ) AS category,

		  -- embed subcategory as JSON
		  json_build_object(
		    'id', sc.id,
		    'name', sc.name
		  ) AS subcategory,

		  COALESCE(
		    json_agg(
		      json_build_object(
		        'id', pv.id,
		        'product_id', pv.product_id,
		        'unit_id', pv.unit_id,
		        'unit_size', pv.unit_size,
		        'type', pv.type,
		        'name', pv.name,
		        'description', pv.description,
		        'price', pv.price,
				'unit', CASE 
				  WHEN u.id IS NULL THEN NULL
				  ELSE json_build_object(
				    'id', u.id,
				    'magnitude', u.magnitude,
				    'name', u.name,
				    'symbol', u.symbol
				  )
				END
		      )
		    ) FILTER (WHERE pv.id IS NOT NULL), '[]'
		  ) AS variants
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		LEFT JOIN subcategories sc ON p.subcategory_id = sc.id
		LEFT JOIN product_variants pv ON p.id = pv.product_id
		LEFT JOIN units u ON pv.unit_id = u.id
	`

	// --- Dynamic WHERE builder ---
	var conditions []string
	var args []interface{}
	argPos := 1

	if query.Status != "" {
		conditions = append(conditions, fmt.Sprintf("p.status = $%d", argPos))
		args = append(args, query.Status)
		argPos++
	}

	if len(conditions) > 0 {
		q += " WHERE " + strings.Join(conditions, " AND ")
	}

	// TODO: add limit and offest!
	// if query.Limit > 0 {
	//	// Ensure Limit and Offset are parameterized just like Status
	//	q += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	//	args = append(args, query.Limit, query.Offset)
	// }

	q += " GROUP BY p.id, c.id, sc.id ORDER BY p.id"

	rows, err := repository.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type productRow struct {
		ID            int64           `json:"id"`
		CategoryID    int64           `json:"category_id"`
		SubcategoryID int64           `json:"subcategory_id"`
		SKU           string          `json:"sku"`
		Image         sql.NullString  `json:"image"`
		Name          string          `json:"name"`
		Description   sql.NullString  `json:"description"`
		Status        string          `json:"status"`
		Category      json.RawMessage `json:"category"`
		Subcategory   json.RawMessage `json:"subcategory"`
		Variants      json.RawMessage `json:"variants"`
	}
	var products []*model.Product

	for rows.Next() {
		var r productRow
		if err := rows.Scan(
			&r.ID, &r.CategoryID, &r.SubcategoryID, &r.SKU,
			&r.Image, &r.Name, &r.Description, &r.Status,
			&r.Category, &r.Subcategory, &r.Variants,
		); err != nil {
			return nil, err
		}

		var category *model.Category
		if len(r.Category) > 0 {
			_ = json.Unmarshal(r.Category, &category)
		}

		var subcategory *model.Subcategory
		if len(r.Subcategory) > 0 {
			_ = json.Unmarshal(r.Subcategory, &subcategory)
		}

		var variants []*model.Variant
		if len(r.Variants) > 0 {
			_ = json.Unmarshal(r.Variants, &variants)
		}

		products = append(products, &model.Product{
			ID: r.ID, CategoryID: r.CategoryID, SubcategoryID: r.SubcategoryID,
			SKU: r.SKU, Image: r.Image.String, Name: r.Name,
			Description: r.Description.String, Status: r.Status,
			Category: category, Subcategory: subcategory, Variants: variants,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (repository productRepository) GetProductDetail(ctx context.Context, id int64) (*model.Product, error) {
	q := `
		SELECT 
		  p.id, p.category_id, p.subcategory_id, p.sku, 
		  p.image, p.name, p.description, p.status,
			
		  -- embed category as JSON
		  json_build_object(
		    'id', c.id,
		    'name', c.name
		  ) AS category,

		  -- embed subcategory as JSON
		  json_build_object(
		    'id', sc.id,
		    'name', sc.name
		  ) AS subcategory,

		  COALESCE(
		    json_agg(
		      json_build_object(
		        'id', pv.id,
		        'product_id', pv.product_id,
		        'unit_id', pv.unit_id,
		        'unit_size', pv.unit_size,
		        'type', pv.type,
		        'name', pv.name,
		        'description', pv.description,
		        'price', pv.price,
				'unit', CASE 
				  WHEN u.id IS NULL THEN NULL
				  ELSE json_build_object(
				    'id', u.id,
				    'magnitude', u.magnitude,
				    'name', u.name,
				    'symbol', u.symbol
				  )
				END
		      )
		    ) FILTER (WHERE pv.id IS NOT NULL), '[]'
		  ) AS variants
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		LEFT JOIN subcategories sc ON p.subcategory_id = sc.id
		LEFT JOIN product_variants pv ON p.id = pv.product_id
		LEFT JOIN units u ON pv.unit_id = u.id
		WHERE p.id = $1
		GROUP BY p.id, c.id, sc.id
	`
	row := repository.db.QueryRow(ctx, q, id)

	type productRow struct {
		ID            int64           `json:"id"`
		CategoryID    int64           `json:"category_id"`
		SubcategoryID int64           `json:"subcategory_id"`
		SKU           string          `json:"sku"`
		Image         sql.NullString  `json:"image"`
		Name          string          `json:"name"`
		Description   sql.NullString  `json:"description"`
		Status        string          `json:"status"`
		Category      json.RawMessage `json:"category"`
		Subcategory   json.RawMessage `json:"subcategory"`
		Variants      json.RawMessage `json:"variants"`
	}

	var r productRow
	if err := row.Scan(
		&r.ID, &r.CategoryID, &r.SubcategoryID, &r.SKU,
		&r.Image, &r.Name, &r.Description, &r.Status,
		&r.Category, &r.Subcategory, &r.Variants,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}

	var category *model.Category
	if len(r.Category) > 0 {
		_ = json.Unmarshal(r.Category, &category)
	}

	var subcategory *model.Subcategory
	if len(r.Subcategory) > 0 {
		_ = json.Unmarshal(r.Subcategory, &subcategory)
	}

	var variants []*model.Variant
	if len(r.Variants) > 0 {
		_ = json.Unmarshal(r.Variants, &variants)
	}

	return &model.Product{
		ID:            r.ID,
		CategoryID:    r.CategoryID,
		SubcategoryID: r.SubcategoryID,
		SKU:           r.SKU,
		Image:         r.Image.String,
		Name:          r.Name,
		Description:   r.Description.String,
		Status:        r.Status,
		Category:      category,
		Subcategory:   subcategory,
		Variants:      variants,
	}, nil
}

func (repository productRepository) CreateProduct(ctx context.Context, form *NewProductForm) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// build insert items
	var columns []string
	var placeholders []string
	args := make([]any, 0)
	argPos := 1

	// set the status
	columns = append(columns, "status")
	placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
	args = append(args, form.Status)
	argPos++
	// set the image
	if form.Image != "" {
		columns = append(columns, "image")
		placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
		args = append(args, form.Image)
		argPos++
	}
	// set the sku
	columns = append(columns, "sku")
	placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
	args = append(args, form.SKU)
	argPos++
	// set the name
	columns = append(columns, "name")
	placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
	args = append(args, form.Name)
	argPos++
	// set the category_id
	columns = append(columns, "category_id")
	placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
	args = append(args, form.CategoryID)
	argPos++
	// set the subcategory_id
	columns = append(columns, "subcategory_id")
	placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
	args = append(args, form.SubcategoryID)
	argPos++
	// set the description
	if form.Description != "" {
		columns = append(columns, "description")
		placeholders = append(placeholders, "$"+strconv.Itoa(argPos))
		args = append(args, form.Description)
		argPos++
	}

	var productID int64
	q := fmt.Sprintf(
		`INSERT INTO products (%s) VALUES (%s) RETURNING id`,
		strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if err = tx.QueryRow(ctx, q, args...).Scan(&productID); err != nil {
		return err
	}

	if len(form.Variants) > 0 {
		rows := make([]string, 0, len(form.Variants))
		varArgs := make([]interface{}, 0, len(form.Variants)*6) // 6 fields per variant
		varArgPos := 1
		for _, sub := range form.Variants {
			var unitID any = nil
			if sub.UnitID > 0 {
				unitID = sub.UnitID
			}

			// normalize UnitSize → nil if not set
			var unitSize any = nil
			if sub.UnitSize > 0 {
				unitSize = sub.UnitSize
			}

			// placeholders: ($1, $2, $3, $4, $5, $6, $7) per row
			rows = append(rows, fmt.Sprintf("($1, $%d, $%d, $%d, $%d, $%d, $%d)",
				varArgPos+1, varArgPos+2, varArgPos+3, varArgPos+4, varArgPos+5, varArgPos+6))
			varArgs = append(varArgs, sub.Type, sub.Name, sub.Description, sub.Price, unitID, unitSize)
			varArgPos += 6
		}
		varArgs = append([]interface{}{productID}, varArgs...)
		query := fmt.Sprintf(
			"INSERT INTO product_variants (product_id, type, name, description, price, unit_id, unit_size) VALUES %s",
			strings.Join(rows, ","))
		if _, err = tx.Exec(ctx, query, varArgs...); err != nil {
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	// set tx = nil so rollback is skipped
	tx = nil
	return nil
}

func (repository productRepository) UpdateProduct(ctx context.Context, form *ProductUpdateForm) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// update the product
	var setPrClauses []string
	var prArgs []any
	prArgPos := 1

	if form.Status != "" {
		setPrClauses = append(setPrClauses, fmt.Sprintf("status = $%d", prArgPos))
		prArgs = append(prArgs, form.Status)
		prArgPos++
	}
	if form.Image != "" {
		setPrClauses = append(setPrClauses, fmt.Sprintf("image = $%d", prArgPos))
		prArgs = append(prArgs, form.Image)
		prArgPos++
	}
	if form.SKU != "" {
		setPrClauses = append(setPrClauses, fmt.Sprintf("sku = $%d", prArgPos))
		prArgs = append(prArgs, form.SKU)
		prArgPos++
	}
	if form.Name != "" {
		setPrClauses = append(setPrClauses, fmt.Sprintf("name = $%d", prArgPos))
		prArgs = append(prArgs, form.Name)
		prArgPos++
	}
	if form.CategoryID > 0 {
		setPrClauses = append(setPrClauses, fmt.Sprintf("category_id = $%d", prArgPos))
		prArgs = append(prArgs, form.CategoryID)
		prArgPos++
	}
	if form.SubcategoryID > 0 {
		setPrClauses = append(setPrClauses, fmt.Sprintf("subcategory_id = $%d", prArgPos))
		prArgs = append(prArgs, form.SubcategoryID)
		prArgPos++
	}
	if form.Description != "" {
		setPrClauses = append(setPrClauses, fmt.Sprintf("description = $%d", prArgPos))
		prArgs = append(prArgs, form.Description)
		prArgPos++
	}
	if len(setPrClauses) > 0 {
		prArgs = append(prArgs, form.ID)
		q := fmt.Sprintf(`UPDATE products SET %s WHERE id = $%d`,
			strings.Join(setPrClauses, ", "), prArgPos)
		if _, err := tx.Exec(ctx, q, prArgs...); err != nil {
			return err
		}
	}

	// update exist variants
	// For EditVariants, we need to run an UPDATE per variant,
	// because each variant might update different fields and has its own id.
	// It’s not practical to build one big bulk UPDATE like we did with the INSERT.
	// update existing variants
	if len(form.EditVariants) > 0 {
		for _, sub := range form.EditVariants {
			var setClauses []string
			var args []any
			argPos := 1

			// always include product_id in args to ensure ownership
			args = append(args, form.ID)

			if sub.Type != "" {
				setClauses = append(setClauses, fmt.Sprintf("type = $%d", argPos+1))
				args = append(args, sub.Type)
				argPos++
			}
			if sub.Name != "" {
				setClauses = append(setClauses, fmt.Sprintf("name = $%d", argPos+1))
				args = append(args, sub.Name)
				argPos++
			}
			if sub.Description != "" {
				setClauses = append(setClauses, fmt.Sprintf("description = $%d", argPos+1))
				args = append(args, sub.Description)
				argPos++
			}
			if sub.Price > 0 {
				setClauses = append(setClauses, fmt.Sprintf("price = $%d", argPos+1))
				args = append(args, sub.Price)
				argPos++
			}

			// normalize UnitID
			var unitID any = nil
			if sub.UnitID > 0 {
				unitID = sub.UnitID
			}
			setClauses = append(setClauses, fmt.Sprintf("unit_id = $%d", argPos+1))
			args = append(args, unitID)
			argPos++

			// normalize UnitSize
			var unitSize any = nil
			if sub.UnitSize > 0 {
				unitSize = sub.UnitSize
			}
			setClauses = append(setClauses, fmt.Sprintf("unit_size = $%d", argPos+1))
			args = append(args, unitSize)
			argPos++

			// add variant id at the end
			args = append(args, sub.ID)

			if len(setClauses) == 0 {
				continue
			}

			q := fmt.Sprintf(
				"UPDATE product_variants SET %s WHERE product_id = $1 AND id = $%d",
				strings.Join(setClauses, ", "), argPos+1,
			)
			if _, err := tx.Exec(ctx, q, args...); err != nil {
				return err
			}
		}
	}

	// create new variants for selected product
	if len(form.NewVariants) > 0 {
		rows := make([]string, 0, len(form.NewVariants))
		varArgs := make([]interface{}, 0, len(form.NewVariants)*6) // 6 fields per variant
		varArgPos := 1
		for _, sub := range form.NewVariants {
			var unitID any = nil
			if sub.UnitID > 0 {
				unitID = sub.UnitID
			}

			// normalize UnitSize → nil if not set
			var unitSize any = nil
			if sub.UnitSize > 0 {
				unitSize = sub.UnitSize
			}

			// placeholders: ($1, $2, $3, $4, $5, $6, $7) per row
			rows = append(rows, fmt.Sprintf("($1, $%d, $%d, $%d, $%d, $%d, $%d)",
				varArgPos+1, varArgPos+2, varArgPos+3, varArgPos+4, varArgPos+5, varArgPos+6))
			varArgs = append(varArgs, sub.Type, sub.Name, sub.Description, sub.Price, unitID, unitSize)
			varArgPos += 6
		}
		varArgs = append([]interface{}{form.ID}, varArgs...)
		query := fmt.Sprintf(
			"INSERT INTO product_variants (product_id, type, name, description, price, unit_id, unit_size) VALUES %s",
			strings.Join(rows, ","))
		if _, err = tx.Exec(ctx, query, varArgs...); err != nil {
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	// set tx = nil so rollback is skipped
	tx = nil
	return nil
}

func (repository productRepository) DeleteProduct(ctx context.Context, id int64) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// First delete variants under this product
	if _, err := tx.Exec(ctx, "DELETE FROM product_variants WHERE product_id = $1", id); err != nil {
		return err
	}

	// Then delete the product itself
	if _, err := tx.Exec(ctx, "DELETE FROM products WHERE id = $1", id); err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	// set tx = nil so rollback is skipped
	tx = nil
	return nil
}

func (repository productRepository) DeleteProductVariant(ctx context.Context, pid, vid int64) error {
	q := "DELETE FROM product_variants WHERE product_id = $1 AND id = $2"
	_, err := repository.db.Exec(ctx, q, pid, vid)
	return err
}

func NewProductRepository(db model.IPgxPool) IProductRepository {
	return &productRepository{db: db}
}
