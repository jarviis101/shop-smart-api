package pkg

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func CreateDatabaseConnection(cfg Database) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.URL)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err := m.Up(); err != nil {
		return nil, err
	}

	return db, err
}
