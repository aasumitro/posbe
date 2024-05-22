package sql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type StoreShiftSQLRepository struct {
	Db *sql.DB
}

func (repo StoreShiftSQLRepository) All(
	ctx context.Context,
) (data []*model.Shift, err error) {
	// TODO: join with store_shift and validate
	// if today and current time is open/close or not
	q := "SELECT * FROM shifts"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var s model.Shift
		if err := rows.Scan(
			&s.ID, &s.Name, &s.StartTime,
			&s.EndTime, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		data = append(data, &s)
	}
	return data, nil
}

// Find will not impl
func (repo StoreShiftSQLRepository) Find(
	_ context.Context,
	_ model.FindWith, _ any,
) (data *model.Shift, err error) {
	// TODO implement me
	panic("implement me")
}

func (repo StoreShiftSQLRepository) Create(
	ctx context.Context,
	params *model.Shift,
) (data *model.Shift, err error) {
	q := "INSERT INTO shifts "
	q += "(name, start_time, end_time, created_at) "
	q += " VALUES ($1, $2, $3, $4) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.Name, params.StartTime,
		params.EndTime, time.Now().Unix())
	data = &model.Shift{}
	if err := row.Scan(&data.ID, &data.Name, &data.StartTime,
		&data.EndTime, &data.CreatedAt, &data.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo StoreShiftSQLRepository) Update(
	ctx context.Context,
	params *model.Shift,
) (data *model.Shift, err error) {
	q := "UPDATE shifts SET "
	q += "name = $1, start_time = $2, "
	q += "end_time = $3, updated_at = $4 "
	q += " WHERE id = $5 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.Name, params.StartTime,
		params.EndTime, time.Now().Unix(), params.ID)
	data = &model.Shift{}
	if err := row.Scan(&data.ID, &data.Name, &data.StartTime,
		&data.EndTime, &data.CreatedAt, &data.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo StoreShiftSQLRepository) Delete(
	ctx context.Context,
	params *model.Shift,
) error {
	// TODO: Check store_shift & transaction
	// if shift being used by this 2 collection
	// then reject the deletion command,
	// instead user just can update this item
	qsst := "SELECT store_shifts.id as id, COUNT(orders) as order_count "
	qsst += "FROM store_shifts WHERE shift_id = $1 "
	qsst += "LEFT OUTER JOIN orders ON orders.shift_id = id"
	row := repo.Db.QueryRowContext(ctx, qsst, params.ID)
	shiftTR := &model.StoreShiftTransaction{}
	if err := row.Scan(
		&shiftTR.ID,
		&shiftTR.OrderCount,
	); err != nil {
		return err
	}
	if shiftTR.OrderCount > 0 {
		return fmt.Errorf(
			"ERROR_RELATION: store shift used by %d transaction",
			shiftTR.OrderCount)
	}
	tr, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	qs := "DELETE FROM shifts WHERE id = $1"
	if _, err := tr.ExecContext(ctx, qs, params.ID); err != nil {
		_ = tr.Rollback()
		return err
	}
	qss := "DELETE FROM store_shifts WHERE shift_id = $1"
	if _, err := tr.ExecContext(ctx, qss, params.ID); err != nil {
		_ = tr.Rollback()
		return err
	}
	return tr.Commit()
}

func (repo StoreShiftSQLRepository) OpenShift(
	ctx context.Context,
	form *model.StoreShiftForm,
) error {
	// TODO: before open validate if theres open shift or not
	// qss := "SELECT * FROM store_shifts WHERE close_at = null"

	qssi := "INSERT INTO store_shifts "
	qssi += "(shift_id, open_at, open_by, open_cash, created_at) "
	qssi += " VALUES ($1, $2, $3, $4, $5) RETURNING id"
	return repo.Db.QueryRowContext(ctx, qssi,
		form.ShiftID, time.Now().Unix(), form.UserID,
		form.Cash, time.Now().Unix()).Err()
}

func (repo StoreShiftSQLRepository) CloseShift(
	ctx context.Context,
	form *model.StoreShiftForm,
) error {
	// TODO: before close validate theres open transaction or not
	// qo := "SELECT count(*) FROM orders WHERE shift_id = $1 "
	// qo += "AND status NOT IN ('paid', 'cancel') AND time_close = null"
	q := "UPDATE store_shifts SET "
	q += "close_at = $1, close_by = $2, "
	q += "close_cash = $3, updated_at = $4 "
	q += " WHERE id = $5 AND shift_id = $6 RETURNING id"
	return repo.Db.QueryRowContext(ctx, q,
		time.Now().Unix(), form.UserID,
		form.Cash, time.Now().Unix(),
		form.ID, form.ShiftID).Err()
}

func NewStoreShiftSQLRepository() model.IStoreShiftRepository {
	return &StoreShiftSQLRepository{Db: config.PostgresPool}
}
