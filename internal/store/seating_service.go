package store

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type seatingService struct {
	repository IStoreSeatingRepository
}

func (service seatingService) FloorList(ctx context.Context) ([]*model.Floor, *utils.ServiceError) {
	data, err := service.repository.GetAllFloors(ctx)
	return utils.HandleMultipleResults[model.Floor]("floors", data, err)
}

func (service seatingService) FloorDetail(ctx context.Context, id int64) (*model.Floor, *utils.ServiceError) {
	data, err := service.repository.GetFloorByID(ctx, id)
	return utils.HandleSingleResult[model.Floor]("floors", data, err)
}

func (service seatingService) CreateFloor(ctx context.Context, form *SeatingFloorRequest) *utils.ServiceError {
	if err := service.repository.CreateFloor(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service seatingService) UpdateFloor(ctx context.Context, form *SeatingFloorRequest) *utils.ServiceError {
	if err := service.repository.UpdateFloor(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service seatingService) DeleteFloor(ctx context.Context, id int64) *utils.ServiceError {
	floors, err := service.FloorList(ctx)
	if err != nil {
		return err
	}

	if len(floors) == 1 {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "cannot delete this floor",
		}
	}

	if err := service.repository.DeleteFloor(ctx, id); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service seatingService) TableList(ctx context.Context, floorID int64) ([]*model.Table, *utils.ServiceError) {
	data, err := service.repository.GetAllTable(ctx, floorID)
	return utils.HandleMultipleResults[model.Table]("tables", data, err)
}

func (service seatingService) TableDetail(ctx context.Context, id int64) (*model.Table, *utils.ServiceError) {
	data, err := service.repository.GetTableByID(ctx, id)
	return utils.HandleSingleResult[model.Table]("table", data, err)
}

func (service seatingService) CreateTable(ctx context.Context, form *SeatingTableRequest) *utils.ServiceError {
	if err := service.repository.CreateTable(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service seatingService) UpdateTable(ctx context.Context, form *SeatingTableRequest) *utils.ServiceError {
	if err := service.repository.UpdateTable(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service seatingService) DeleteTable(ctx context.Context, tableID int64) *utils.ServiceError {
	// TODO: validate in use or not (has active order)!
	if err := service.repository.DeleteTable(ctx, tableID); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (service seatingService) ProceedEvent(ctx context.Context, data string) {
	var payload SeatingTableRequest

	form, err := payload.Bind(data)
	if err != nil {
		log.Println("Failed to bind seating table request", err)
		return
	}

	if err := service.repository.UpdateTable(ctx, form); err != nil {
		log.Printf("error updating table: %v", err)
		return
	}

	if form.Status != "" {
		cacheKey := fmt.Sprintf("table:status:%d", form.ID)
		config.RedisCache.Set(ctx, cacheKey, form.Status, -1)
	}

	log.Printf("[ProceedEvent] table updated successfully: ID=%d", form.ID)
}

func NewSeatingService(repository IStoreSeatingRepository) IStoreSeatingService {
	return &seatingService{repository: repository}
}
