package catalog

import (
	"context"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
)

type IAttributeRepository interface {
	GetAllUnit(ctx context.Context) ([]*model.Unit, error)
	CreateUnit(ctx context.Context, form *UnitForm) error
	UpdateUnit(ctx context.Context, form *UnitForm) error
	DeleteUnit(ctx context.Context, id int64) error
}

type IAttributeService interface {
	UnitList(ctx context.Context) ([]*model.Unit, *utils.ServiceError)
	CreateUnit(ctx context.Context, form *UnitForm) *utils.ServiceError
	UpdateUnit(ctx context.Context, form *UnitForm) *utils.ServiceError
	DeleteUnit(ctx context.Context, id int64) *utils.ServiceError
}
