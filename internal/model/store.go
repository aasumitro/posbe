package model

import (
	"database/sql"
)

const (
	StoreSettingCacheKey = "store_settings"
)

type (
	// StorePref store setting
	StorePref struct {
		Key       string `json:"key" form:"key" binding:"required"`
		Value     string `json:"value" form:"value" binding:"required"`
		CreatedAt sql.NullInt64
		UpdatedAt sql.NullInt64
	}

	StoreSetting map[string]interface{}

	//Floor struct {
	//	ID          int           `json:"id"`
	//	Name        string        `json:"name" form:"name" binding:"required"`
	//	TotalTables int           `json:"total_tables,omitempty"`
	//	TotalRooms  int           `json:"total_rooms,omitempty"`
	//	CreatedAt   sql.NullInt64 `json:"created_at"`
	//	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	//	Tables      []*Table      `json:"tables,omitempty" binding:"-"`
	//}

	// Table Case Study Restaurant Dine in
	//Table struct {
	//	ID        int           `json:"id"`
	//	FloorID   int           `json:"floor_id" form:"floor_id" binding:"required"`
	//	Name      string        `json:"name" form:"name" binding:"required"`
	//	XPos      float32       `json:"x_pos" form:"x_pos" binding:"required"`
	//	YPos      float32       `json:"y_pos" form:"y_pos" binding:"required"`
	//	WSize     float32       `json:"w_size" form:"w_size" binding:"required"`
	//	HSize     float32       `json:"h_size" form:"h_size" binding:"required"`
	//	Capacity  int           `json:"capacity" form:"capacity" binding:"required"`
	//	Type      string        `json:"type" form:"type"`
	//	CreatedAt sql.NullInt64 `json:"created_at"`
	//	UpdatedAt sql.NullInt64 `json:"updated_at,omitempty"`
	//}

	//StoreShiftTransaction struct {
	//	ID         int
	//	OrderCount int
	//}

	//StoreShiftForm struct {
	//	ID      int   `json:"-" form:"-"`
	//	UserID  int   `json:"-" form:"-"`
	//	ShiftID int   `json:"shift_id" form:"shift_id" binding:"required"`
	//	Cash    int64 `json:"cash" form:"cash" binding:"required"`
	//}

	//IStoreService interface {
	//	FloorList(ctx context.Context) (floors []*Floor, errData *utils.ServiceError)
	//	AddFloor(ctx context.Context, data *Floor) (floor *Floor, errData *utils.ServiceError)
	//	EditFloor(ctx context.Context, data *Floor) (floor *Floor, errData *utils.ServiceError)
	//	DeleteFloor(ctx context.Context, data *Floor) *utils.ServiceError
	//
	//	TableList(ctx context.Context) (table []*Table, errData *utils.ServiceError)
	//	AddTable(ctx context.Context, data *Table) (table *Table, errData *utils.ServiceError)
	//	EditTable(ctx context.Context, data *Table) (table *Table, errData *utils.ServiceError)
	//	DeleteTable(ctx context.Context, data *Table) *utils.ServiceError
	//
	//	FloorsWith(ctx context.Context, s any) (floors []*Floor, errData *utils.ServiceError)
	//}

	//IStoreShiftRepository interface {
	//	ICRUDRepository[Shift]
	//	OpenShift(ctx context.Context, form *StoreShiftForm) error
	//	CloseShift(ctx context.Context, form *StoreShiftForm) error
	//}

	//IStoreShiftService interface {
	//	ShiftList(ctx context.Context) (shifts []*Shift, errData *utils.ServiceError)
	//	ShiftDetail(ctx context.Context, id int) (shift *Shift, errData *utils.ServiceError)
	//}
)
