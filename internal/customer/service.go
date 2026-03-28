package customer

import (
	"context"
	"net/http"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type service struct {
	repository ICustomerRepository
}

func (s service) List(
	ctx context.Context, query *RequestQuery,
) ([]*model.Customer, *utils.ServiceError) {
	data, err := s.repository.GetAll(ctx, query)
	return utils.HandleMultipleResults[model.Customer]("customers", data, err)
}

func (s service) Create(
	ctx context.Context, form *RequestForm,
) *utils.ServiceError {
	if err := s.repository.Insert(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (s service) Update(
	ctx context.Context, form *RequestForm,
) *utils.ServiceError {
	if err := s.repository.Update(ctx, form); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (s service) Destroy(
	ctx context.Context, id int64,
) *utils.ServiceError {
	if err := s.repository.Delete(ctx, id); err != nil {
		return &utils.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func NewService(repository ICustomerRepository) ICustomerService {
	return &service{repository: repository}
}
