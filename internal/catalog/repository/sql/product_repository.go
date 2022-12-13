package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/config"
)

type ProductSQLRepository struct {
	Db *sql.DB
}

func (repo ProductSQLRepository) AllWhere(ctx context.Context, key domain.FindWith, val any) (data []*domain.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductSQLRepository) All(ctx context.Context) (data []*domain.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductSQLRepository) Find(ctx context.Context, key domain.FindWith, val any) (data *domain.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductSQLRepository) Create(ctx context.Context, params *domain.Product) (data *domain.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductSQLRepository) Update(ctx context.Context, params *domain.Product) (data *domain.Product, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo ProductSQLRepository) Delete(ctx context.Context, params *domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func NewProductSQLRepository(db *sql.DB) domain.ICRUDAddOnRepository[domain.Product] {
	return &ProductSQLRepository{Db: config.DbPool}
}
