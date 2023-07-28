package console

import (
	"database/sql"
	"shop-smart-api/internal/app"
	di "shop-smart-api/internal/infrastructure/container"
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
	container := di.CreateContainer(a.database, a.serverConfig)

	userUseCase := container.ProvideUserService()

	manager := seeder.CreateSeeder(userUseCase)
	return manager.Seed()
}
