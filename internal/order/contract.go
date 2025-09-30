package order

import (
	"context"

	"github.com/aasumitro/posbe/internal/utils"
)

type IShiftRepository interface {
	Open(ctx context.Context, form *ActiveShiftForm) error
	Close(ctx context.Context, form *ActiveShiftForm) error
}

type IShiftService interface {
	ActiveShiftAction(ctx context.Context, form *ActiveShiftForm) *utils.ServiceError
}
