package sql

import (
	"context"
	"database/sql"
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
	tr, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := "DELETE FROM shifts WHERE id = $1"
	if _, err := tr.ExecContext(ctx, q, params.ID); err != nil {
		_ = tr.Rollback()
		return err
	}
	q2 := "DELETE FROM store_shifts WHERE shift_id = $1"
	if _, err := tr.ExecContext(ctx, q2, params.ID); err != nil {
		_ = tr.Rollback()
		return err
	}
	return tr.Commit()
}

func (repo StoreShiftSQLRepository) OpenShift(
	ctx context.Context,
	form *model.StoreShiftForm,
) error {
	q := "INSERT INTO store_shifts "
	q += "(shift_id, open_at, open_by, open_cash, created_at) "
	q += " VALUES ($1, $2, $3, $4, $5) RETURNING id"
	return repo.Db.QueryRowContext(ctx, q,
		form.ShiftID, time.Now().Unix(), form.UserID,
		form.Cash, time.Now().Unix()).Err()
}

func (repo StoreShiftSQLRepository) CloseShift(
	ctx context.Context,
	form *model.StoreShiftForm,
) error {
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
