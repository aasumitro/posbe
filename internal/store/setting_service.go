package store

import (
	"context"
	"net/http"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type settingService struct {
	repository IStoreSettingRepository
}

func (service settingService) AllSetting(ctx context.Context) (*model.StoreSetting, *utils.ServiceError) {
	data, err := utils.CacheFirstData(ctx, config.RedisCache,
		&utils.CacheDataSupplied[*model.StoreSetting]{
			Key: model.StoreSettingCacheKey, TTL: time.Hour * 1,
			CbF: func() (*model.StoreSetting, error) {
				return service.repository.GetAll(ctx)
			},
		},
	)

	return utils.HandleSingleResult[model.StoreSetting]("settings", data, err)
}

func (service settingService) UpdateSetting(ctx context.Context, form *SettingForm) *utils.ServiceError {
	settings := form.ToStoreModelMap()
	if len(settings) == 0 {
		return &utils.ServiceError{
			Code:    http.StatusBadRequest,
			Message: "form value is required",
		}
	}

	if err := service.repository.Update(ctx, &settings); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	config.RedisCache.Del(ctx, model.StoreSettingCacheKey)

	return nil
}

func NewSettingService(repository IStoreSettingRepository) IStoreSettingService {
	return &settingService{repository: repository}
}
