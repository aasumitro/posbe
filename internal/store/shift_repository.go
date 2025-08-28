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
	//TODO implement me
	panic("implement me")
}

func (repository shiftRepository) Create(ctx context.Context, form *ShiftForm) error {
	q := `
    INSERT INTO shifts (name, start_time, end_time, created_at)
    VALUES ($1, $2, $3, $4)
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
	//TODO implement me
	panic("implement me")
}

func (repository shiftRepository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewShiftRepository(db model.IPgxPool) IStoreShiftRepository {
	return &shiftRepository{db: db}
}
