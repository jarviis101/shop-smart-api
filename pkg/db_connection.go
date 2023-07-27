package pkg

import (
	"database/sql"
)

func CreateDatabaseConnection(cfg Database) (*sql.DB, error) {
	return sql.Open("postgres", cfg.URL)
}
