package console

import (
	"database/sql"
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/infrastructure/repository"
	"shop-smart-api/internal/infrastructure/seeder"
	"shop-smart-api/pkg"
)

type application struct {
	database     *sql.DB
	serverConfig pkg.Server
}

func CreateApplication(db *sql.DB, sc pkg.Server) app.Application {
	return &application{db, sc}
}

func (a *application) Run() error {
	userRepository := repository.CreateUserRepository(a.database)

	manager := seeder.CreateSeeder(userRepository)
	return manager.Seed()
}
