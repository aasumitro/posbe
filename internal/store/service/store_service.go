package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"reflect"

	"github.com/aasumitro/posbe/common"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/aasumitro/posbe/pkg/utils"
)

type storeService struct {
	floorRepo model.ICRUDRepository[model.Floor]
	tableRepo model.ICRUDAddOnRepository[model.Table]
	roomRepo  model.ICRUDAddOnRepository[model.Room]
}

func (service storeService) FloorList(
	ctx context.Context,
) (floors []*model.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.All(ctx)
	return utils.ValidateDataRows[model.Floor](data, err)
}

func (service storeService) AddFloor(
	ctx context.Context,
	item *model.Floor,
) (floor *model.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Floor](data, err)
}

func (service storeService) EditFloor(
	ctx context.Context,
	item *model.Floor,
) (floor *model.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Floor](data, err)
}

func (service storeService) DeleteFloor(
	ctx context.Context,
	data *model.Floor,
) *utils.ServiceError {
	floor, err := service.floorRepo.Find(
		ctx, model.FindWithID, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if floor.TotalTables >= 1 {
		return &utils.ServiceError{
			Code:    http.StatusForbidden,
			Message: common.ErrorUnableToDelete,
		}
	}
	if err := service.floorRepo.Delete(ctx, floor); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service storeService) TableList(
	ctx context.Context,
) (table []*model.Table, errData *utils.ServiceError) {
	data, err := service.tableRepo.All(ctx)
	return utils.ValidateDataRows[model.Table](data, err)
}

func (service storeService) AddTable(
	ctx context.Context,
	item *model.Table,
) (table *model.Table, errData *utils.ServiceError) {
	data, err := service.tableRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Table](data, err)
}

func (service storeService) EditTable(
	ctx context.Context,
	item *model.Table,
) (table *model.Table, errData *utils.ServiceError) {
	data, err := service.tableRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Table](data, err)
}

func (service storeService) DeleteTable(
	ctx context.Context,
	data *model.Table,
) *utils.ServiceError {
	table, err := service.tableRepo.Find(ctx, model.FindWithID, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := service.tableRepo.Delete(ctx, table); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service storeService) RoomList(
	ctx context.Context,
) (rooms []*model.Room, errData *utils.ServiceError) {
	data, err := service.roomRepo.All(ctx)
	return utils.ValidateDataRows[model.Room](data, err)
}

func (service storeService) AddRoom(
	ctx context.Context,
	item *model.Room,
) (room *model.Room, errData *utils.ServiceError) {
	data, err := service.roomRepo.Create(ctx, item)
	return utils.ValidateDataRow[model.Room](data, err)
}

func (service storeService) EditRoom(
	ctx context.Context,
	item *model.Room,
) (room *model.Room, errData *utils.ServiceError) {
	data, err := service.roomRepo.Update(ctx, item)
	return utils.ValidateDataRow[model.Room](data, err)
}

func (service storeService) DeleteRoom(
	ctx context.Context,
	data *model.Room,
) *utils.ServiceError {
	room, err := service.roomRepo.Find(
		ctx, model.FindWithID, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := service.roomRepo.Delete(ctx, room); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service storeService) FloorsWith(
	ctx context.Context,
	s any) (floors []*model.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.All(ctx)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sName := reflect.TypeOf(s).Name()
	var fa []*model.Floor
	for _, floor := range data {
		if sName == "Table" && floor.TotalTables >= 1 {
			f := floor

			if tables, err := service.tableRepo.AllWhere(
				ctx,
				model.FindWithRelationID,
				floor.ID,
			); err != nil {
				f.Tables = nil
			} else {
				f.Tables = tables
			}

			fa = append(fa, f)
		}

		if sName == "Room" && floor.TotalRooms >= 1 {
			f := floor

			if rooms, err := service.roomRepo.AllWhere(
				ctx,
				model.FindWithRelationID,
				floor.ID,
			); err != nil {
				f.Rooms = nil
			} else {
				f.Rooms = rooms
			}

			fa = append(fa, f)
		}
	}

	return fa, nil
}

func NewStoreService(
	floorRepo model.ICRUDRepository[model.Floor],
	tableRepo model.ICRUDAddOnRepository[model.Table],
	roomRepo model.ICRUDAddOnRepository[model.Room],
) model.IStoreService {
	return &storeService{
		tableRepo: tableRepo,
		floorRepo: floorRepo,
		roomRepo:  roomRepo,
	}
}
