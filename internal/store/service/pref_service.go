package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type prefService struct {
	prefRepo model.IStorePrefRepository
}

func (service prefService) AllPrefs(
	ctx context.Context,
) (prefs *model.StoreSetting, errData *utils.ServiceError) {
	data, err := service.prefRepo.All(ctx)
	return utils.ValidateDataRow[model.StoreSetting](data, err)
}

func (service prefService) UpdatePrefs(
	ctx context.Context,
	key, value string,
) (prefs *model.StoreSetting, errData *utils.ServiceError) {
	_, err := service.prefRepo.Find(ctx, key)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &utils.ServiceError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		}

		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	data, err := service.prefRepo.Update(ctx, key, value)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return data, nil
}

func NewPreferenceService(
	prefRepo model.IStorePrefRepository,
) model.IStorePrefService {
	return &prefService{
		prefRepo: prefRepo,
	}
}
