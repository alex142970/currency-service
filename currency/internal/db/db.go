package db

import (
	"currency_service/currency/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect(c *config.DatabaseConfig) *sql.DB {
	db, err := sql.Open("postgres", c.GetDSN())

	if err != nil {
		panic(err)
	}

	return db
}
