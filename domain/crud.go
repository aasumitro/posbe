package domain

import "context"

type ICRUDRepository[T any] interface {
	All(ctx context.Context) (data []*T, err error)
	Find(ctx context.Context, id int) (data *T, err error)
	Create(ctx context.Context, params *T) (data *T, err error)
	Update(ctx context.Context, params *T) (data *T, err error)
	Delete(ctx context.Context, params *T) error
}
