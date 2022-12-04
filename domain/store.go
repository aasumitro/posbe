package domain

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
		FloorId   int           `json:"floor_id" form:"floor_id" binding:"required"`
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
		FloorId   int           `json:"floor_id" form:"floor_id" binding:"required"`
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

	// StorePref store setting
	StorePref struct {
		Key       string `json:"key" form:"key" binding:"required"`
		Value     string `json:"value" form:"value" binding:"required"`
		CreatedAt sql.NullInt64
		UpdatedAt sql.NullInt64
	}

	StoreSetting map[string]string

	IStoreService interface {
		FloorList() (floors []*Floor, errData *utils.ServiceError)
		AddFloor(data *Floor) (floor *Floor, errData *utils.ServiceError)
		EditFloor(data *Floor) (floor *Floor, errData *utils.ServiceError)
		DeleteFloor(data *Floor) *utils.ServiceError

		TableList() (table []*Table, errData *utils.ServiceError)
		AddTable(data *Table) (table *Table, errData *utils.ServiceError)
		EditTable(data *Table) (table *Table, errData *utils.ServiceError)
		DeleteTable(data *Table) *utils.ServiceError

		RoomList() (rooms []*Room, errData *utils.ServiceError)
		AddRoom(data *Room) (room *Room, errData *utils.ServiceError)
		EditRoom(data *Room) (room *Room, errData *utils.ServiceError)
		DeleteRoom(data *Room) *utils.ServiceError

		FloorsWith(s any) (floors []*Floor, errData *utils.ServiceError)
	}

	IStorePrefRepository interface {
		Find(ctx context.Context, key string) (pref *StoreSetting, err error)
		All(ctx context.Context) (prefs *StoreSetting, err error)
		Update(ctx context.Context, key, value string) (prefs *StoreSetting, err error)
	}

	IStorePrefService interface {
		AllPrefs() (prefs *StoreSetting, errData *utils.ServiceError)
		UpdatePrefs(key, value string) (prefs *StoreSetting, errData *utils.ServiceError)
	}
)
