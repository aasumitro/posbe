package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func (cfg Config) InitDbConn() {
	log.Println("Trying to open database connection . . . .")
	conn, err := openConnection(cfg)

	if err != nil {
		log.Panicf("DATABASE_ERROR: %s", err.Error())
	}

	log.Printf("Database connected with %s driver . . . .", cfg.DBDriver)
	setConnection(conn)
}

func openConnection(cfg Config) (db *sql.DB, err error) {
	driver, err := sql.Open("postgres", cfg.DBDsnUrl)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func setConnection(conn *sql.DB) {
	db = conn
}

func (cfg Config) GetDbConn() *sql.DB {
	return db
}

func (cfg Config) DeferCloseDbConn() {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Panicf("DATABASE_ERROR: %s", err.Error())
		}
	}()
}
