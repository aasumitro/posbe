package model

import (
	"context"
	"database/sql"

	"github.com/aasumitro/posbe/pkg/utils"
)

type (
	Floor struct {
		ID          int           `json:"id"`
		Name        string        `json:"name" form:"name" binding:"required"`
		TotalTables int           `json:"total_tables,omitempty"`
		TotalRooms  int           `json:"total_rooms,omitempty"`
		CreatedAt   sql.NullInt64 `json:"created_at"`
		UpdatedAt   sql.NullInt64 `json:"updated_at"`
		Tables      []*Table      `json:"tables,omitempty" binding:"-"`
		Rooms       []*Room       `json:"rooms,omitempty" binding:"-"`
	}

	// Table Case Study Restaurant Dine in
	Table struct {
		ID        int           `json:"id"`
		FloorID   int           `json:"floor_id" form:"floor_id" binding:"required"`
		Name      string        `json:"name" form:"name" binding:"required"`
		XPos      float32       `json:"x_pos" form:"x_pos" binding:"required"`
		YPos      float32       `json:"y_pos" form:"y_pos" binding:"required"`
		WSize     float32       `json:"w_size" form:"w_size" binding:"required"`
		HSize     float32       `json:"h_size" form:"h_size" binding:"required"`
		Capacity  int           `json:"capacity" form:"capacity" binding:"required"`
		Type      string        `json:"type" form:"type"`
		CreatedAt sql.NullInt64 `json:"created_at"`
		UpdatedAt sql.NullInt64 `json:"updated_at,omitempty"`
	}

	// Room Case Study Karaoke
	Room struct {
		ID        int           `json:"id"`
		FloorID   int           `json:"floor_id" form:"floor_id" binding:"required"`
		Name      string        `json:"name" form:"name" binding:"required"`
		XPos      float32       `json:"x_pos" form:"x_pos" binding:"required"`
		YPos      float32       `json:"y_pos" form:"y_pos" binding:"required"`
		WSize     float32       `json:"w_size" form:"w_size" binding:"required"`
		HSize     float32       `json:"h_size" form:"h_size" binding:"required"`
		Capacity  int           `json:"capacity" form:"capacity" binding:"required"`
		Price     float32       `json:"price" form:"price" binding:"required"`
		CreatedAt sql.NullInt64 `json:"created_at"`
		UpdatedAt sql.NullInt64 `json:"updated_at,omitempty"`
	}

	// Shift is reference section data for store
	Shift struct {
		ID           int           `json:"id"`
		Name         string        `json:"name"`
		StartTime    int64         `json:"start_time"`
		EndTime      int64         `json:"end_time"`
		CreatedAt    sql.NullInt64 `json:"created_at"`
		UpdatedAt    sql.NullInt64 `json:"updated_at,omitempty"`
		CurrentShift *StoreShift   `json:"current_shift,omitempty" binding:"-"`
	}

	// StorePref store setting
	StorePref struct {
		Key       string `json:"key" form:"key" binding:"required"`
		Value     string `json:"value" form:"value" binding:"required"`
		CreatedAt sql.NullInt64
		UpdatedAt sql.NullInt64
	}

	StoreShift struct {
		ID        int           `json:"id"`
		ShiftID   int           `json:"shift_id"`
		OpenAt    int64         `json:"open_at"`
		OpenBy    sql.NullInt64 `json:"open_by"`
		OpenCash  sql.NullInt64 `json:"open_cash"`
		CloseAt   sql.NullInt64 `json:"close_at"`
		CloseBy   sql.NullInt64 `json:"close_by"`
		CloseCash sql.NullInt64 `json:"close_cash"`
		CreatedAt sql.NullInt64 `json:"created_at"`
		UpdatedAt sql.NullInt64 `json:"updated_at,omitempty"`
		Shift     *Shift        `json:"shift,omitempty" binding:"-"`
	}

	StoreShiftTransaction struct {
		ID         int
		OrderCount int
	}

	StoreShiftForm struct {
		ID      int   `json:"-" form:"-"`
		UserID  int   `json:"-" form:"-"`
		ShiftID int   `json:"shift_id" form:"shift_id" binding:"required"`
		Cash    int64 `json:"cash" form:"cash" binding:"required"`
	}

	StoreSetting map[string]interface{}

	IStoreService interface {
		FloorList(ctx context.Context) (floors []*Floor, errData *utils.ServiceError)
		AddFloor(ctx context.Context, data *Floor) (floor *Floor, errData *utils.ServiceError)
		EditFloor(ctx context.Context, data *Floor) (floor *Floor, errData *utils.ServiceError)
		DeleteFloor(ctx context.Context, data *Floor) *utils.ServiceError

		TableList(ctx context.Context) (table []*Table, errData *utils.ServiceError)
		AddTable(ctx context.Context, data *Table) (table *Table, errData *utils.ServiceError)
		EditTable(ctx context.Context, data *Table) (table *Table, errData *utils.ServiceError)
		DeleteTable(ctx context.Context, data *Table) *utils.ServiceError

		RoomList(ctx context.Context) (rooms []*Room, errData *utils.ServiceError)
		AddRoom(ctx context.Context, data *Room) (room *Room, errData *utils.ServiceError)
		EditRoom(ctx context.Context, data *Room) (room *Room, errData *utils.ServiceError)
		DeleteRoom(ctx context.Context, data *Room) *utils.ServiceError

		FloorsWith(ctx context.Context, s any) (floors []*Floor, errData *utils.ServiceError)
	}

	IStorePrefRepository interface {
		Find(ctx context.Context, key string) (pref *StoreSetting, err error)
		All(ctx context.Context) (prefs *StoreSetting, err error)
		Update(ctx context.Context, key, value string) (prefs *StoreSetting, err error)
	}

	IStorePrefService interface {
		AllPrefs(ctx context.Context) (prefs *StoreSetting, errData *utils.ServiceError)
		UpdatePrefs(ctx context.Context, key, value string) (prefs *StoreSetting, errData *utils.ServiceError)
	}

	IStoreShiftRepository interface {
		ICRUDRepository[Shift]
		OpenShift(ctx context.Context, form *StoreShiftForm) error
		CloseShift(ctx context.Context, form *StoreShiftForm) error
	}
)
