package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func (cfg Config) InitDbConn() {
	log.Println("Trying to open database connection . . . .")

	dbOnce.Do(func() {
		conn, err := sql.Open("postgres", cfg.DBDsnUrl)
		if err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		Db = conn
		log.Printf("Database connected with %s driver . . . .", cfg.DBDriver)
	})
}
