package order

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type IShiftRepository interface {
	ActiveShift(ctx context.Context) (*model.ActiveShift, error)
	Open(ctx context.Context, form *ActiveShiftForm) error
	Close(ctx context.Context, form *ActiveShiftForm) error
}

type IShiftService interface {
	CurrentActiveShift(ctx context.Context) (*model.ActiveShift, *utils.ServiceError)
	ActiveShiftAction(ctx context.Context, form *ActiveShiftForm) *utils.ServiceError
}

type ITransactionRepository interface{}

type ITransactionService interface{}
