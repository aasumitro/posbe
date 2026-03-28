package customer

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type ICustomerRepository interface {
	GetAll(ctx context.Context, query *RequestQuery) ([]*model.Customer, error)
	Insert(ctx context.Context, form *RequestForm) error
	Update(ctx context.Context, form *RequestForm) error
	Delete(ctx context.Context, id int64) error
}

type ICustomerService interface {
	List(ctx context.Context, query *RequestQuery) ([]*model.Customer, *utils.ServiceError)
	Create(ctx context.Context, form *RequestForm) *utils.ServiceError
	Update(ctx context.Context, form *RequestForm) *utils.ServiceError
	Destroy(ctx context.Context, id int64) *utils.ServiceError
}
