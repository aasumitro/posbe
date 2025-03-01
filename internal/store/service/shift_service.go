package service

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type shiftService struct {
	shiftRepo model.IStoreShiftRepository
}

func (s shiftService) ShiftList(
	ctx context.Context,
) (shifts []*model.Shift, errData *utils.ServiceError) {
	data, err := s.shiftRepo.All(ctx)
	return utils.ValidateDataRows[model.Shift](data, err)
}

func (s shiftService) ShiftDetail(
	ctx context.Context,
	id int,
) (shift *model.Shift, errData *utils.ServiceError) {
	data, err := s.shiftRepo.Find(ctx, model.FindWithID, id)
	return utils.ValidateDataRow[model.Shift](data, err)
}

func NewStoreShiftService(
	shiftRepo model.IStoreShiftRepository,
) model.IStoreShiftService {
	return &shiftService{
		shiftRepo: shiftRepo,
	}
}
