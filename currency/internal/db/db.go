package db

import (
	"currency_service/currency/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDatabaseConnection(c *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", c.GetDSN())

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
