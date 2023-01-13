package domain

import "context"

type FindWith int64

const (
	FindWithID FindWith = iota
	FindWithRelationID
	FindWithName
	// for user data
	FindWithUsername
	FindWithEmail
	FindWithPhone
	// for product data
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
