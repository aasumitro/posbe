package config

import (
	"database/sql"
	"log"

	// postgresql
	_ "github.com/lib/pq"
)

func PostgresConnection() Option {
	return func(cfg *Config) {
		const dbDriverName = "postgres"
		postgresSingleton.Do(func() {
			log.Println("Trying to open database connection pool . . . .")
			conn, err := sql.Open(dbDriverName,
				cfg.PostgresDsnURL)
			if err != nil {
				log.Fatalf("DATABASE_ERROR: %s\n",
					err.Error())
			}
			PostgresPool = conn
			if err := PostgresPool.Ping(); err != nil {
				log.Fatalf("DATABASE_ERROR: %s\n",
					err.Error())
			}
			log.Printf("Database connected with %s driver . . . .\n",
				dbDriverName)
		})
	}
}
