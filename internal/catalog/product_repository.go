package catalog

import (
	"context"
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

func NewProductRepository(db model.IPgxPool) IProductRepository {
	return &productRepository{db: db}
}
