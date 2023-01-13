package service

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/errors"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
	"reflect"
)

type storeService struct {
	ctx       context.Context
	floorRepo domain.ICRUDRepository[domain.Floor]
	tableRepo domain.ICRUDAddOnRepository[domain.Table]
	roomRepo  domain.ICRUDAddOnRepository[domain.Room]
}

func (service storeService) FloorList() (floors []*domain.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Floor](data, err)
}

func (service storeService) AddFloor(data *domain.Floor) (floor *domain.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Floor](data, err)
}

func (service storeService) EditFloor(data *domain.Floor) (floor *domain.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Floor](data, err)
}

func (service storeService) DeleteFloor(data *domain.Floor) *utils.ServiceError {
	floor, err := service.floorRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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
			Message: errors.ErrorUnableToDelete,
		}
	}

	err = service.floorRepo.Delete(service.ctx, floor)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service storeService) TableList() (table []*domain.Table, errData *utils.ServiceError) {
	data, err := service.tableRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Table](data, err)
}

func (service storeService) AddTable(data *domain.Table) (table *domain.Table, errData *utils.ServiceError) {
	data, err := service.tableRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Table](data, err)
}

func (service storeService) EditTable(data *domain.Table) (table *domain.Table, errData *utils.ServiceError) {
	data, err := service.tableRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Table](data, err)
}

func (service storeService) DeleteTable(data *domain.Table) *utils.ServiceError {
	table, err := service.tableRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	err = service.tableRepo.Delete(service.ctx, table)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service storeService) RoomList() (rooms []*domain.Room, errData *utils.ServiceError) {
	data, err := service.roomRepo.All(service.ctx)

	return utils.ValidateDataRows[domain.Room](data, err)
}

func (service storeService) AddRoom(data *domain.Room) (room *domain.Room, errData *utils.ServiceError) {
	data, err := service.roomRepo.Create(service.ctx, data)

	return utils.ValidateDataRow[domain.Room](data, err)
}

func (service storeService) EditRoom(data *domain.Room) (room *domain.Room, errData *utils.ServiceError) {
	data, err := service.roomRepo.Update(service.ctx, data)

	return utils.ValidateDataRow[domain.Room](data, err)
}

func (service storeService) DeleteRoom(data *domain.Room) *utils.ServiceError {
	room, err := service.roomRepo.Find(service.ctx, domain.FindWithID, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
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

	err = service.roomRepo.Delete(service.ctx, room)
	if err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (service storeService) FloorsWith(s any) (floors []*domain.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.All(service.ctx)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	sName := reflect.TypeOf(s).Name()
	var fa []*domain.Floor
	for _, floor := range data {
		if sName == "Table" && floor.TotalTables >= 1 {
			f := floor

			if tables, err := service.tableRepo.AllWhere(
				service.ctx,
				domain.FindWithRelationID,
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
				service.ctx,
				domain.FindWithRelationID,
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
	ctx context.Context,
	floorRepo domain.ICRUDRepository[domain.Floor],
	tableRepo domain.ICRUDAddOnRepository[domain.Table],
	roomRepo domain.ICRUDAddOnRepository[domain.Room],
) domain.IStoreService {
	return &storeService{
		ctx:       ctx,
		tableRepo: tableRepo,
		floorRepo: floorRepo,
		roomRepo:  roomRepo,
	}
}
