package store

import (
	"context"
	"net/http"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type shiftService struct {
	repository IStoreShiftRepository
}

func (service shiftService) ShiftList(ctx context.Context) ([]*model.Shift, *utils.ServiceError) {
	data, err := service.repository.GetAll(ctx)
	return utils.HandleMultipleResults[model.Shift]("shifts", data, err)
}

func (service shiftService) ShiftDetail(ctx context.Context, id int64) (*model.Shift, *utils.ServiceError) {
	//TODO implement me
	panic("implement me")
}

func (service shiftService) CreateShift(ctx context.Context, form *ShiftForm) *utils.ServiceError {
	if err := service.repository.Create(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func NewShiftService(repository IStoreShiftRepository) IStoreShiftService {
	return &shiftService{repository: repository}
}
