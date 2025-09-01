package model

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type IPgxPool interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

type FindWith int64

const (
	FindWithID FindWith = iota
	FindWithRelationID

	FindWithUsername
	FindWithEmail

	FindWithSKU
	FindWithCategoryID
	FindWithSubcategoryID
	FindWithPriceInRange
)

type ICRUDRepository[T any] interface {
	All(ctx context.Context) (data []*T, err error)
	Find(ctx context.Context, key FindWith, val any) (data *T, err error)
	Create(ctx context.Context, params *T) (data *T, err error)
	Update(ctx context.Context, params *T) (data *T, err error)
	Delete(ctx context.Context, params *T) error
}

// ICRUDAddOnRepository TODO: rename
type ICRUDAddOnRepository[T any] interface {
	AllWhere(ctx context.Context, key FindWith, val any) (data []*T, err error)
	ICRUDRepository[T]
}

type ICRUDWithSearchRepository[T any] interface {
	Search(ctx context.Context, keys []FindWith, values []any) (data []*T, err error)
	ICRUDRepository[T]
}
