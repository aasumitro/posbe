package store

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type shiftService struct {
	repository IStoreShiftRepository
}

func (service shiftService) ShiftList(ctx context.Context) ([]*model.Shift, *utils.ServiceError) {
	data, err := utils.CacheFirstData(ctx, config.RdpPool,
		&utils.CacheDataSupplied[[]*model.Shift]{
			Key: model.StoreShiftCacheKey, TTL: time.Minute * 30,
			CbF: func() ([]*model.Shift, error) {
				return service.repository.GetAll(ctx)
			},
		},
	)

	return utils.HandleMultipleResults[model.Shift]("shifts", data, err)
}

func (service shiftService) ShiftDetail(ctx context.Context, id int64) (*model.Shift, *utils.ServiceError) {
	data, err := service.repository.FindByID(ctx, id)

	// TODO: proceed the data
	//      data counting
	//      last/current active shift
	//      profit/loss

	return utils.HandleSingleResult[model.Shift]("shift", data, err)
}

func (service shiftService) CreateShift(ctx context.Context, form *ShiftForm) *utils.ServiceError {
	if err := service.repository.Create(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	config.RdpPool.Del(ctx, model.StoreShiftCacheKey)

	return nil
}

func (service shiftService) UpdateShift(ctx context.Context, form *ShiftForm) *utils.ServiceError {
	if err := service.repository.Update(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	config.RdpPool.Del(ctx, model.StoreShiftCacheKey)

	return nil
}

func (service shiftService) DeleteShift(ctx context.Context, id int64) *utils.ServiceError {
	if slices.Contains([]int64{1, 2}, id) {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "Cannot delete default shift",
		}
	}

	if err := service.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, ErrShiftHasOrders) {
			return &utils.ServiceError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	config.RdpPool.Del(ctx, model.StoreShiftCacheKey)

	return nil
}

func (service shiftService) ActiveShiftAction(ctx context.Context, form *ActiveShiftForm) *utils.ServiceError {
	if form.Action == "open" {
		if err := service.repository.Open(ctx, form); err != nil {
			return &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	if form.Action == "close" {
		if err := service.repository.Close(ctx, form); err != nil {
			return &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	config.RdpPool.Del(ctx, model.StoreShiftCacheKey)

	return nil
}

func NewShiftService(repository IStoreShiftRepository) IStoreShiftService {
	return &shiftService{repository: repository}
}
