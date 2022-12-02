package domain

import (
	"database/sql"
	"github.com/aasumitro/posbe/pkg/utils"
)

type (
	Floor struct {
		ID          int            `json:"id"`
		Name        string         `json:"name" form:"name" binding:"required"`
		TotalTables int            `json:"total_tables,omitempty"`
		TotalRooms  int            `json:"total_rooms,omitempty"`
		CreatedAt   sql.NullString `json:"created_at"`
		UpdatedAt   sql.NullString `json:"updated_at"`
		Tables      []*Table       `json:"tables,omitempty" binding:"-"`
		Rooms       []*Room        `json:"rooms,omitempty" binding:"-"`
	}

	// Table Case Study Restaurant Dine in
	Table struct {
		ID        int            `json:"id"`
		FloorId   int            `json:"floor_id" form:"floor_id" binding:"required"`
		Name      string         `json:"name" form:"name" binding:"required"`
		XPos      float32        `json:"x_pos" form:"x_pos" binding:"required"`
		YPos      float32        `json:"y_pos" form:"y_pos" binding:"required"`
		WSize     float32        `json:"w_size" form:"w_size" binding:"required"`
		HSize     float32        `json:"h_size" form:"h_size" binding:"required"`
		Capacity  int            `json:"capacity" form:"capacity" binding:"required"`
		Type      string         `json:"type" form:"type"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at,omitempty"`
	}

	// Room Case Study Karaoke
	Room struct {
		ID        int            `json:"id"`
		FloorId   int            `json:"floor_id" form:"floor_id" binding:"required"`
		Name      string         `json:"name" form:"name" binding:"required"`
		XPos      float32        `json:"x_pos" form:"x_pos" binding:"required"`
		YPos      float32        `json:"y_pos" form:"y_pos" binding:"required"`
		WSize     float32        `json:"w_size" form:"w_size" binding:"required"`
		HSize     float32        `json:"h_size" form:"h_size" binding:"required"`
		Capacity  int            `json:"capacity" form:"capacity" binding:"required"`
		Price     float32        `json:"price" form:"type"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at,omitempty"`
	}

	IStoreService interface {
		FloorList() (floors []*Floor, errData *utils.ServiceError)
		AddFloor(data *Floor) (floor *Floor, errData *utils.ServiceError)
		EditFloor(data *Floor) (floor *Floor, errData *utils.ServiceError)
		DeleteFloor(data *Floor) *utils.ServiceError

		TableList() (table []*Table, errData *utils.ServiceError)
		AddTable(data *Table) (table *Table, errData *utils.ServiceError)
		EditTable(data *Table) (table *Table, errData *utils.ServiceError)
		DeleteTable(data *Table) *utils.ServiceError

		FloorsWith(s any) (floors []*Floor, errData *utils.ServiceError)
	}
)
