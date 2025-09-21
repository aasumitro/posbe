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

func NewAttributeRepository(db model.IPgxPool) IAttributeRepository {
	return &attributeRepository{db: db}
}
