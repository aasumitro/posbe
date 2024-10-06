package config

import (
	"database/sql"
	"log"
	"time"

	// postgresql
	_ "github.com/lib/pq"
)

const (
	maxOpenConn = 50
	maxIdleConn = 10
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
			conn.SetMaxOpenConns(maxOpenConn)
			conn.SetMaxIdleConns(maxIdleConn)
			conn.SetConnMaxLifetime(time.Hour)
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
