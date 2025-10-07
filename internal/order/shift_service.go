package order

import (
	"context"
	"net/http"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type shiftService struct {
	repository IShiftRepository
}

func (service shiftService) CurrentActiveShift(ctx context.Context) (*model.ActiveShift, *utils.ServiceError) {
	data, err := service.repository.ActiveShift(ctx)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return utils.HandleSingleResult[model.ActiveShift]("active_shift", data, err)
}

func (service shiftService) ActiveShiftAction(ctx context.Context, form *ActiveShiftForm) *utils.ServiceError {
	if form.Action == ShiftActionOpen {
		if err := service.repository.Open(ctx, form); err != nil {
			return &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	if form.Action == ShiftActionClose {
		// TODO: order/transactions validations
		// if still has open orders then they should completed it first.

		if err := service.repository.Close(ctx, form); err != nil {
			return &utils.ServiceError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	config.RedisCache.Del(ctx, model.StoreShiftCacheKey)

	return nil
}

func NewShiftService(repository IShiftRepository) IShiftService {
	return &shiftService{repository: repository}
}
