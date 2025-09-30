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

func NewShiftService(repository IShiftRepository) IShiftService {
	return &shiftService{repository: repository}
}
