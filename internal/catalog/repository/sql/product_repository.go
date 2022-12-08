package sql

import "database/sql"

type ProductSQLRepository struct {
	Db *sql.DB
}

func NewProductSQLRepository(db *sql.DB) {
	//return &ProductSQLRepository{db: db}
}
