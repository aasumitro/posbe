package sql

import "database/sql"

type UserSQLRepository struct {
	Db *sql.DB
}
