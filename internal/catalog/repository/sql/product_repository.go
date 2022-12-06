package sql

import "database/sql"

type ProductSQLRepository struct {
	Db *sql.DB
}
