package service

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"net/http"
)

type storePrefService struct {
	ctx      context.Context
	prefRepo domain.IStorePrefRepository
}

func (service storePrefService) AllPrefs() (prefs *domain.StoreSetting, errData *utils.ServiceError) {
	data, err := service.prefRepo.All(service.ctx)

	return utils.ValidateDataRow[domain.StoreSetting](data, err)
}

func (service storePrefService) UpdatePrefs(key, value string) (prefs *domain.StoreSetting, errData *utils.ServiceError) {
	_, err := service.prefRepo.Find(service.ctx, key)
	if err != nil {
		if err == sql.ErrNoRows {
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

	data, err := service.prefRepo.Update(service.ctx, key, value)
	if err != nil {
		return nil, &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return data, nil
}

func NewStorePrefService(
	ctx context.Context,
	prefRepo domain.IStorePrefRepository,
) domain.IStorePrefService {
	return &storePrefService{
		ctx:      ctx,
		prefRepo: prefRepo,
	}
}
