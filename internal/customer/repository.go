package customer

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	db model.IPgxPool
}

func (r repository) GetAll(
	ctx context.Context, query *RequestQuery,
) ([]*model.Customer, error) {
	q := "SELECT id, name, phone, email, description FROM customers"

	// TODO: apply query

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	var customer model.Customer
	if _, err := pgx.ForEachRow(rows, []any{
		&customer.ID, &customer.Name,
		&customer.Phone, &customer.Email,
		&customer.Description,
	}, func() error {
		customers = append(customers, new(customer))
		return nil
	}); err != nil {
		return nil, err
	}
	return customers, nil
}

func (r repository) Insert(
	ctx context.Context, form *RequestForm,
) error {
	q := `
		INSERT INTO customers (
			name, phone, email, description, created_at
		) VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, q, form.Name, form.Phone, form.Email, form.Description, time.Now().Unix())
	return err
}

func (r repository) Update(
	ctx context.Context, form *RequestForm,
) error {
	if form.ID == 0 {
		return errors.New("missing ID for update")
	}

	var setClauses []string
	var args []any
	argPos := 1

	if form.Name != "" {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, form.Name)
		argPos++
	}
	if form.Phone != "" {
		setClauses = append(setClauses, "phone = $"+strconv.Itoa(argPos))
		args = append(args, form.Phone)
		argPos++
	}
	if form.Email != "" {
		setClauses = append(setClauses, "email = $"+strconv.Itoa(argPos))
		args = append(args, form.Email)
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

	setClauses = append(setClauses, "updated_at = $"+strconv.Itoa(argPos))
	args = append(args, time.Now().Unix())
	argPos++

	args = append(args, form.ID)

	q := fmt.Sprintf(`UPDATE customers SET %s WHERE id = $%d`,
		strings.Join(setClauses, ", "), argPos)
	_, err := r.db.Exec(ctx, q, args...)
	return err
}

func (r repository) Delete(
	ctx context.Context, id int64,
) error {
	q := "DELETE FROM customers WHERE id = $1"
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func NewRepository(db model.IPgxPool) ICustomerRepository {
	return &repository{db: db}
}
