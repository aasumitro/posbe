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
