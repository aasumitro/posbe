package sql

import "database/sql"

type VariantSQLRepository struct {
	Db *sql.DB
}
