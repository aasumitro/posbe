package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
)

type attributeRepository struct {
	db model.IPgxPool
}

func (repository attributeRepository) GetAllUnit(ctx context.Context) ([]*model.Unit, error) {
	q := `
		SELECT u.id, u.magnitude, u.name, u.symbol,
		       (SELECT COUNT(*) FROM product_variants AS pv WHERE pv.unit_id = u.id) AS usage
		FROM units AS u
	`
	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []*model.Unit
	var unit model.Unit
	if _, err = pgx.ForEachRow(rows, []any{
		&unit.ID, &unit.Magnitude, &unit.Name,
		&unit.Symbol, &unit.Usage,
	}, func() error {
		u := unit
		units = append(units, &u)
		return nil
	}); err != nil {
		return nil, err
	}

	return units, nil
}

func (repository attributeRepository) CreateUnit(ctx context.Context, form *UnitForm) error {
	q := "INSERT INTO units (magnitude, name, symbol) VALUES ($1, $2, $3)"
	_, err := repository.db.Exec(ctx, q, form.Magnitude, form.Name, form.Symbol)
	return err
}

func (repository attributeRepository) UpdateUnit(ctx context.Context, form *UnitForm) error {
	if form.ID == 0 {
		return errors.New("missing ID for update")
	}

	var setClauses []string
	var args []any
	argPos := 1

	if form.Magnitude != "" {
		setClauses = append(setClauses, "magnitude = $"+strconv.Itoa(argPos))
		args = append(args, form.Magnitude)
		argPos++
	}
	if form.Name != "" {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, form.Name)
		argPos++
	}
	if form.Symbol != "" {
		setClauses = append(setClauses, "symbol = $"+strconv.Itoa(argPos))
		args = append(args, form.Symbol)
		argPos++
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, form.ID)

	q := fmt.Sprintf(`UPDATE units SET %s WHERE id = $%d`,
		strings.Join(setClauses, ", "), argPos)
	_, err := repository.db.Exec(ctx, q, args...)
	return err
}

func (repository attributeRepository) DeleteUnit(ctx context.Context, id int64) error {
	q := "DELETE FROM units WHERE id = $1"
	_, err := repository.db.Exec(ctx, q, id)
	return err
}

func (repository attributeRepository) GetAllCategory(ctx context.Context) ([]*model.Category, error) {
	q := `
		SELECT 
		  c.id, 
		  c.name,
		  (SELECT COUNT(*) FROM products AS p WHERE p.category_id = c.id) AS usage,
		  COALESCE(
		    json_agg(
		      json_build_object(
		        'id', sc.id,
		        'category_id', sc.category_id,
		        'name', sc.name,
		        'usage', COALESCE(
				  (SELECT COUNT(*) FROM products AS p WHERE p.subcategory_id = sc.id),
				  0
				)
		      )
		    ) FILTER (WHERE sc.id IS NOT NULL), 
		    '[]'
		  ) AS subcategories
		FROM categories c
		LEFT JOIN subcategories sc ON c.id = sc.category_id
		GROUP BY c.id, c.name
		ORDER BY c.id;
	`
	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type categoryRow struct {
		ID            int64           `json:"id"`
		Name          string          `json:"name"`
		Usage         int64           `json:"usage"`
		Subcategories json.RawMessage `json:"subcategories"`
	}
	var categories []*model.Category

	for rows.Next() {
		var r categoryRow
		if err := rows.Scan(&r.ID, &r.Name, &r.Usage, &r.Subcategories); err != nil {
			return nil, err
		}
		var subs []*model.Subcategory
		if err := json.Unmarshal(r.Subcategories, &subs); err != nil {
			return nil, err
		}
		categories = append(categories, &model.Category{
			ID: r.ID, Name: r.Name, Usage: r.Usage, Subcategories: subs,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (repository attributeRepository) CreateCategory(ctx context.Context, form *NewCategoryForm) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// Insert category
	var categoryID int64
	if err = tx.QueryRow(ctx,
		`INSERT INTO categories (name) VALUES ($1) RETURNING id`,
		form.Name,
	).Scan(&categoryID); err != nil {
		return err
	}

	// Insert subcategories if provided
	if len(form.Subcategories) > 0 {
		rows := make([]string, 0, len(form.Subcategories))
		args := make([]interface{}, 0, len(form.Subcategories)*2)
		for i, sub := range form.Subcategories {
			rows = append(rows, fmt.Sprintf("($1, $%d)", i+2))
			args = append(args, sub)
		}
		args = append([]interface{}{categoryID}, args...)
		query := fmt.Sprintf("INSERT INTO subcategories (category_id, name) VALUES %s", strings.Join(rows, ","))
		if _, err = tx.Exec(ctx, query, args...); err != nil {
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

func (repository attributeRepository) UpdateCategory(ctx context.Context, form *EditCategoryForm) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	if form.Name != "" {
		q := "UPDATE categories SET name = $1 WHERE id = $2"
		if _, err := tx.Exec(ctx, q, form.Name, form.ID); err != nil {
			return err
		}
	}

	if len(form.EditSubcategories) > 0 {
		values := make([]string, 0, len(form.EditSubcategories))
		args := []interface{}{form.ID}

		for i, sub := range form.EditSubcategories {
			if sub.Name == "" {
				continue
			}
			values = append(values, fmt.Sprintf("($%d::BIGINT, $%d::TEXT)", i*2+2, i*2+3))
			args = append(args, sub.ID, sub.Name)
		}

		if len(values) > 0 {
			query := fmt.Sprintf(`
		        UPDATE subcategories AS sc
		        SET name = v.name
		        FROM (VALUES %s) AS v(id, name)
		        WHERE sc.id = v.id
		          AND sc.category_id = $1
		    `, strings.Join(values, ","))
			if _, err := tx.Exec(ctx, query, args...); err != nil {
				return err
			}
		}
	}

	// Insert new subcategories if provided
	if len(form.NewSubcategories) > 0 {
		rows := make([]string, 0, len(form.NewSubcategories))
		args := make([]interface{}, 0, len(form.NewSubcategories)*2)
		for i, sub := range form.NewSubcategories {
			rows = append(rows, fmt.Sprintf("($1, $%d)", i+2))
			args = append(args, sub)
		}
		args = append([]interface{}{form.ID}, args...)
		query := fmt.Sprintf("INSERT INTO subcategories (category_id, name) VALUES %s", strings.Join(rows, ","))
		if _, err = tx.Exec(ctx, query, args...); err != nil {
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

func (repository attributeRepository) DeleteCategory(ctx context.Context, id int64) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// First delete tables under this floor
	if _, err := tx.Exec(ctx, "DELETE FROM subcategories WHERE category_id = $1", id); err != nil {
		return err
	}

	// Then delete the floor itself
	if _, err := tx.Exec(ctx, "DELETE FROM categories WHERE id = $1", id); err != nil {
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

func (repository attributeRepository) DeleteSubcategory(ctx context.Context, cid, sid int64) error {
	q := "DELETE FROM subcategories WHERE category_id = $1 AND id = $2"
	_, err := repository.db.Exec(ctx, q, cid, sid)
	return err
}

func NewAttributeRepository(db model.IPgxPool) IAttributeRepository {
	return &attributeRepository{db: db}
}
