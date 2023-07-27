package pkg

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func CreateDatabaseConnection(cfg Database) (*sql.DB, error) {
	return sql.Open("postgres", cfg.URL)
}
