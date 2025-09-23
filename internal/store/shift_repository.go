package store

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

type shiftRepository struct {
	db model.IPgxPool
}

var ErrShiftHasOrders = errors.New("cannot delete shift, it has orders")

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
			s.Active = &sa
		}
		shifts = append(shifts, &s)
		return nil
	}); err != nil {
		return nil, err
	}

	return shifts, nil
}

func (repository shiftRepository) FindByID(ctx context.Context, id int64) (*model.Shift, error) {
	batch := &pgx.Batch{}
	batch.Queue(`SELECT s.id, s.name, s.start_time, s.end_time, s.created_at, s.updated_at, 
		(SELECT COUNT(*) FROM orders o WHERE o.shift_id = s.id AND o.status <> 'cancel') AS total_transaction
		FROM shifts s WHERE s.id = $1`, id)
	batch.Queue(`SELECT * FROM active_shifts WHERE shift_id = $1 ORDER BY created_at DESC`, id)

	br := repository.db.SendBatch(ctx, batch)
	defer func() { _ = br.Close() }()

	// Fetch shift
	var shift model.Shift
	if err := br.QueryRow().Scan(
		&shift.ID, &shift.Name,
		&shift.StartTime, &shift.EndTime,
		&shift.CreatedAt, &shift.UpdatedAt,
		&shift.TotalTransaction,
	); err != nil {
		return nil, err
	}

	// Fetch active shifts
	rows, err := br.Query()
	if err != nil {
		return nil, err
	}
	var histories []*model.ActiveShift
	for rows.Next() {
		var a model.ActiveShift
		if err := rows.Scan(
			&a.ID, &a.ShiftID,
			&a.OpenAt, &a.OpenBy, &a.OpenCash,
			&a.CloseAt, &a.CloseBy, &a.CloseCash,
			&a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			rows.Close()
			return nil, err
		}
		histories = append(histories, &a)
	}
	rows.Close()
	shift.Histories = histories

	return &shift, nil
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

func (repository shiftRepository) Update(ctx context.Context, form *ShiftForm) error {
	if form.ID == 0 {
		return errors.New("missing shift ID for update")
	}

	var setClauses []string
	var args []any
	argPos := 1

	// Build SET clause dynamically only for fields that are non-zero/non-empty
	if form.Name != "" {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, form.Name)
		argPos++
	}
	if form.StartTime > 0 {
		setClauses = append(setClauses, "start_time = $"+strconv.Itoa(argPos))
		args = append(args, form.StartTime)
		argPos++
	}
	if form.EndTime > 0 {
		setClauses = append(setClauses, "end_time = $"+strconv.Itoa(argPos))
		args = append(args, form.EndTime)
		argPos++
	}
	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}
	setClauses = append(setClauses, "updated_at = $"+strconv.Itoa(argPos))
	args = append(args, time.Now().Unix())
	argPos++
	args = append(args, form.ID) // Add WHERE clause

	q := fmt.Sprintf(`UPDATE shifts SET %s WHERE id = $%d`,
		strings.Join(setClauses, ", "), argPos)
	_, err := repository.db.Exec(ctx, q, args...)
	return err
}

func (repository shiftRepository) Delete(ctx context.Context, id int64) error {
	var count int
	qo := `SELECT COUNT(*) FROM orders WHERE shift_id = $1`
	if err := repository.db.QueryRow(ctx, qo, id).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("%w: %d", ErrShiftHasOrders, count)
	}

	q := "DELETE FROM shifts WHERE id = $1"
	_, err := repository.db.Exec(ctx, q, id)
	return err
}

func (repository shiftRepository) Open(ctx context.Context, form *ActiveShiftForm) error {
	//TODO implement me
	// Check if theres active shift or not if yes throw error to close the prev first

	//	// TODO: before open validate if theres open shift or not
	//	// qss := "SELECT * FROM store_shifts WHERE close_at = null"
	//
	//	qssi := "INSERT INTO store_shifts "
	//	qssi += "(shift_id, open_at, open_by, open_cash, created_at) "
	//	qssi += " VALUES ($1, $2, $3, $4, $5) RETURNING id"
	//	return repo.Db.QueryRowContext(ctx, qssi,
	//		form.ShiftID, time.Now().Unix(), form.UserID,
	//		form.Cash, time.Now().Unix()).Err()

	panic("implement me")
}

func (repository shiftRepository) Close(ctx context.Context, form *ActiveShiftForm) error {
	//TODO implement me
	// check if theres active orders, if yes complete the order first before close

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

	panic("implement me")
}

func NewShiftRepository(db model.IPgxPool) IStoreShiftRepository {
	return &shiftRepository{db: db}
}
