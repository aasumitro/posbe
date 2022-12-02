package service

import (
	"context"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/errors"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
)

type storeService struct {
	ctx       context.Context
	floorRepo domain.ICRUDRepository[domain.Floor]
	tableRepo domain.ICRUDAddOnRepository[domain.Table]
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
	floor, err := service.floorRepo.Find(service.ctx, domain.FindWithId, data.ID)
	if err != nil {
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
	table, err := service.tableRepo.Find(service.ctx, domain.FindWithId, data.ID)
	if err != nil {
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

func (service storeService) FloorsWithTables() (floors []*domain.Floor, errData *utils.ServiceError) {
	data, err := service.floorRepo.All(service.ctx)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	for _, floor := range data {
		var f = floor
		tables, err := service.tableRepo.AllWhere(service.ctx,
			domain.FindWithRelationId, floor.ID)
		if err != nil {
			f.Tables = nil
		}
		f.Tables = tables
		floors = append(floors, f)
	}

	return data, nil
}

func NewStoreService(
	ctx context.Context,
	floorRepo domain.ICRUDRepository[domain.Floor],
	tableRepo domain.ICRUDAddOnRepository[domain.Table],
) domain.IStoreService {
	return &storeService{
		ctx:       ctx,
		tableRepo: tableRepo,
		floorRepo: floorRepo,
	}
}
