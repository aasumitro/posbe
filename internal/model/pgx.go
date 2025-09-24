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
	Begin(ctx context.Context) (pgx.Tx, error)
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
