package sql

import "database/sql"

type AddonSQLRepository struct {
	Db *sql.DB
}
