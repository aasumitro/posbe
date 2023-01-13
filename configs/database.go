package configs

import (
	"database/sql"
	"fmt"
	"log"

	// postgresql
	_ "github.com/lib/pq"
)

func (cfg Config) InitDbConn() {
	log.Println("Trying to open database connection pool . . . .")

	dbOnce.Do(func() {
		conn, err := sql.Open(cfg.DBDriver, cfg.DBDsnURL)
		if err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		DbPool = conn

		if err := DbPool.Ping(); err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		log.Printf("Database connection pool created with %s driver . . . .", cfg.DBDriver)
	})
}
