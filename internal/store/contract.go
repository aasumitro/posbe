package store

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type IStoreSettingRepository interface {
	FindByKey(ctx context.Context, k string) (*model.StoreSetting, error)
	GetAll(ctx context.Context) (*model.StoreSetting, error)
	Update(ctx context.Context, setting *model.StoreSetting) error
}

type IStoreSettingService interface {
	AllSetting(ctx context.Context) (*model.StoreSetting, *utils.ServiceError)
	UpdateSetting(ctx context.Context, form *SettingForm) *utils.ServiceError
}

type IStoreShiftRepository interface {
	// GetAll FindByID Create Update Delete action for table `shifts`
	GetAll(ctx context.Context) ([]*model.Shift, error)
	FindByID(ctx context.Context, id int64) (*model.Shift, error)
	Create(ctx context.Context, form *ShiftForm) error
	Update(ctx context.Context, form *ShiftForm) error
	Delete(ctx context.Context, id int64) error

	// OpenShift CloseShift for `active_shifts` table
	// OpenShift()
	// CloseShift()
}

type IStoreShiftService interface {
	ShiftList(ctx context.Context) ([]*model.Shift, *utils.ServiceError)
	ShiftDetail(ctx context.Context, id int64) (*model.Shift, *utils.ServiceError)
	CreateShift(ctx context.Context, form *ShiftForm) *utils.ServiceError
}
