package domain

import "context"

type FindWith int64

const (
	FindWithId FindWith = iota
	FindWithRelationId
	FindWithName
	FindWithUsername
	FindWithEmail
	FindWithPhone
)

type ICRUDRepository[T any] interface {
	All(ctx context.Context) (data []*T, err error)
	Find(ctx context.Context, key FindWith, val any) (data *T, err error)
	Create(ctx context.Context, params *T) (data *T, err error)
	Update(ctx context.Context, params *T) (data *T, err error)
	Delete(ctx context.Context, params *T) error
}

// ICRUDAddOnRepository
// TODO: rename
type ICRUDAddOnRepository[T any] interface {
	AllWhere(ctx context.Context, key FindWith, val any) (data []*T, err error)
	ICRUDRepository[T]
}
