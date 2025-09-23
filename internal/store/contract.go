package store

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type IStoreSettingRepository interface {
	FindByKey(ctx context.Context, k string) (*model.StoreSetting, error)
	GetAll(ctx context.Context) (*model.StoreSetting, error)
	Update(ctx context.Context, setting *model.StoreSetting) error
}

type IStoreSettingService interface {
	AllSetting(ctx context.Context) (*model.StoreSetting, *utils.ServiceError)
	UpdateSetting(ctx context.Context, form *SettingForm) *utils.ServiceError
}

type IStoreShiftRepository interface {
	GetAll(ctx context.Context) ([]*model.Shift, error)
	FindByID(ctx context.Context, id int64) (*model.Shift, error)
	Create(ctx context.Context, form *ShiftForm) error
	Update(ctx context.Context, form *ShiftForm) error
	Delete(ctx context.Context, id int64) error
	Open(ctx context.Context, form *ActiveShiftForm) error
	Close(ctx context.Context, form *ActiveShiftForm) error
}

type IStoreShiftService interface {
	ShiftList(ctx context.Context) ([]*model.Shift, *utils.ServiceError)
	ShiftDetail(ctx context.Context, id int64) (*model.Shift, *utils.ServiceError)
	CreateShift(ctx context.Context, form *ShiftForm) *utils.ServiceError
	UpdateShift(ctx context.Context, form *ShiftForm) *utils.ServiceError
	DeleteShift(ctx context.Context, id int64) *utils.ServiceError
	ActiveShiftAction(ctx context.Context, form *ActiveShiftForm) *utils.ServiceError
}

type IStoreSeatingRepository interface {
	GetAllFloors(ctx context.Context) ([]*model.Floor, error)
	GetFloorByID(ctx context.Context, id int64) (*model.Floor, error)
	CreateFloor(ctx context.Context, form *SeatingFloorRequest) error
	UpdateFloor(ctx context.Context, form *SeatingFloorRequest) error
	DeleteFloor(ctx context.Context, id int64) error

	GetAllTable(ctx context.Context, floorID int64) ([]*model.Table, error)
	GetTableByID(ctx context.Context, id int64) (*model.Table, error)
	CreateTable(ctx context.Context, form *SeatingTableRequest) error
	UpdateTable(ctx context.Context, form *SeatingTableRequest) error
	DeleteTable(ctx context.Context, id int64) error
}

type IStoreSeatingService interface {
	FloorList(ctx context.Context) ([]*model.Floor, *utils.ServiceError)
	FloorDetail(ctx context.Context, id int64) (*model.Floor, *utils.ServiceError)
	CreateFloor(ctx context.Context, form *SeatingFloorRequest) *utils.ServiceError
	UpdateFloor(ctx context.Context, form *SeatingFloorRequest) *utils.ServiceError
	DeleteFloor(ctx context.Context, id int64) *utils.ServiceError

	TableList(ctx context.Context, floorID int64) ([]*model.Table, *utils.ServiceError)
	TableDetail(ctx context.Context, tableID int64) (*model.Table, *utils.ServiceError)
	CreateTable(ctx context.Context, form *SeatingTableRequest) *utils.ServiceError
	UpdateTable(ctx context.Context, form *SeatingTableRequest) *utils.ServiceError
	DeleteTable(ctx context.Context, tableID int64) *utils.ServiceError
}
