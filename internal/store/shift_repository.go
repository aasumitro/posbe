package store

import (
	"context"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
)

type shiftRepository struct {
	db model.IPgxPool
}

func (repository shiftRepository) GetAll(ctx context.Context) ([]*model.Shift, error) {
	q := `
	SELECT s.id, s.name, s.start_time, s.end_time,
	    s.created_at, s.updated_at,
	    (SELECT COUNT(*) FROM active_shifts AS ss2 WHERE ss2.shift_id = s.id) AS total_usages,
	    COALESCE(ss.id, 0) as last_shift_id, ss.open_at as last_shift_open,
	    ss.close_at as last_shift_close
	FROM shifts AS s
	    LEFT JOIN (
	    SELECT ss1.*
	    FROM active_shifts AS ss1
	    WHERE ss1.id = (SELECT id FROM active_shifts
	        WHERE shift_id = ss1.shift_id ORDER BY created_at DESC
	    LIMIT 1)
	) AS ss ON ss.shift_id = s.id LEFT JOIN orders AS t ON t.shift_id = ss.id
	`
	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []*model.Shift
	var shift model.Shift
	var activeShift model.ActiveShift
	if _, err = pgx.ForEachRow(rows, []any{
		&shift.ID, &shift.Name, &shift.StartTime,
		&shift.EndTime, &shift.CreatedAt, &shift.UpdatedAt,
		&shift.TotalUsage, &activeShift.ID,
		&activeShift.OpenAt, &activeShift.CloseAt,
	}, func() error {
		s := shift
		if activeShift.ID != 0 {
			sa := activeShift
			s.ActiveShift = &sa
		}
		shifts = append(shifts, &s)
		return nil
	}); err != nil {
		return nil, err
	}

	return shifts, nil
}

func (repository shiftRepository) FindByID(ctx context.Context, id int64) (*model.Shift, error) {
	//q := "SELECT * FROM shifts WHERE "
	////goland:noinspection ALL
	//if key == model.FindWithID {
	//	q += "id = $1 "
	//}
	//q += "LIMIT 1"
	//row := repo.Db.QueryRowContext(ctx, q, val)
	//shift = &model.Shift{}
	//if err := row.Scan(
	//	&shift.ID, &shift.Name, &shift.StartTime,
	//	&shift.EndTime, &shift.CreatedAt, &shift.UpdatedAt,
	//); err != nil {
	//	return nil, err
	//}
	//return shift, nil

	panic("implement me")
}

func (repository shiftRepository) Create(ctx context.Context, form *ShiftForm) error {
	q := `
	    INSERT INTO shifts (name, start_time, end_time, created_at) VALUES ($1, $2, $3, $4)
	    RETURNING id, name, start_time, end_time, created_at, updated_at
	`
	var shift model.Shift
	if err := repository.db.QueryRow(
		ctx, q, form.Name, form.StartTime,
		form.EndTime, time.Now().Unix(),
	).Scan(
		&shift.ID, &shift.Name,
		&shift.StartTime, &shift.EndTime,
		&shift.CreatedAt, &shift.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}

func (repository shiftRepository) Update(ctx context.Context, sh *ShiftForm) error {
	//q := "UPDATE shifts SET "
	//q += "name = $1, start_time = $2, "
	//q += "end_time = $3, updated_at = $4 "
	//q += " WHERE id = $5 RETURNING *"
	//row := repo.Db.QueryRowContext(ctx, q,
	//	params.Name, params.StartTime,
	//	params.EndTime, time.Now().Unix(), params.ID)
	//data = &model.Shift{}
	//if err := row.Scan(&data.ID, &data.Name, &data.StartTime,
	//	&data.EndTime, &data.CreatedAt, &data.UpdatedAt,
	//); err != nil {
	//	return nil, err
	//}
	//return data, nil

	panic("implement me")
}

func (repository shiftRepository) Delete(ctx context.Context, id int64) error {
	// TODO: Check store_shift & transaction
	// if shift being used by this 2 collection
	// then reject the deletion command,
	// instead user just can update this item
	//qsst := "SELECT store_shifts.id as id, COUNT(orders) as order_count "
	//qsst += "FROM store_shifts WHERE shift_id = $1 "
	//qsst += "LEFT OUTER JOIN orders ON orders.shift_id = id"
	//row := repo.Db.QueryRowContext(ctx, qsst, params.ID)
	//shiftTR := &model.StoreShiftTransaction{}
	//if err := row.Scan(
	//	&shiftTR.ID,
	//	&shiftTR.OrderCount,
	//); err != nil {
	//	return err
	//}
	//if shiftTR.OrderCount > 0 {
	//	return fmt.Errorf(
	//		"ERROR_RELATION: store shift used by %d transaction",
	//		shiftTR.OrderCount)
	//}
	//tr, err := repo.Db.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//qs := "DELETE FROM shifts WHERE id = $1"
	//if _, err := tr.ExecContext(ctx, qs, params.ID); err != nil {
	//	_ = tr.Rollback()
	//	return err
	//}
	//qss := "DELETE FROM store_shifts WHERE shift_id = $1"
	//if _, err := tr.ExecContext(ctx, qss, params.ID); err != nil {
	//	_ = tr.Rollback()
	//	return err
	//}
	//return tr.Commit()

	panic("implement me")
}

// func (repo ShiftSQLRepository) OpenShift(
//	ctx context.Context,
//	form *model.StoreShiftForm,
//) error {
//	// TODO: before open validate if theres open shift or not
//	// qss := "SELECT * FROM store_shifts WHERE close_at = null"
//
//	qssi := "INSERT INTO store_shifts "
//	qssi += "(shift_id, open_at, open_by, open_cash, created_at) "
//	qssi += " VALUES ($1, $2, $3, $4, $5) RETURNING id"
//	return repo.Db.QueryRowContext(ctx, qssi,
//		form.ShiftID, time.Now().Unix(), form.UserID,
//		form.Cash, time.Now().Unix()).Err()
//}
//
//func (repo ShiftSQLRepository) CloseShift(
//	ctx context.Context,
//	form *model.StoreShiftForm,
//) error {
//	// TODO: before close validate theres open transaction or not
//	// qo := "SELECT count(*) FROM orders WHERE shift_id = $1 "
//	// qo += "AND status NOT IN ('paid', 'cancel') AND time_close = null"
//	q := "UPDATE store_shifts SET "
//	q += "close_at = $1, close_by = $2, "
//	q += "close_cash = $3, updated_at = $4 "
//	q += " WHERE id = $5 AND shift_id = $6 RETURNING id"
//	return repo.Db.QueryRowContext(ctx, q,
//		time.Now().Unix(), form.UserID,
//		form.Cash, time.Now().Unix(),
//		form.ID, form.ShiftID).Err()
//}

func NewShiftRepository(db model.IPgxPool) IStoreShiftRepository {
	return &shiftRepository{db: db}
}
