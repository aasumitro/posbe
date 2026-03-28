package order

import (
	"context"
	"fmt"

	"github.com/aasumitro/posbe/internal/model"
)

type shiftRepository struct {
	db model.IPgxPool
}

func (repository shiftRepository) ActiveShift(ctx context.Context) (*model.ActiveShift, error) {
	////TODO implement me
	//panic("implement me")
	return nil, nil
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

	fmt.Println(form.UserID, form.ShiftID, form.ActiveShiftID, form.Cash)
	return nil
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

	fmt.Println(form.UserID, form.ShiftID, form.ActiveShiftID, form.Cash)
	return nil
}

func NewShiftRepository(db model.IPgxPool) IShiftRepository {
	return &shiftRepository{db: db}
}
