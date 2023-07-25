package console

import (
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/app/di"
	"shop-smart-api/internal/infrastructure/seeder"
	"shop-smart-api/pkg"
)

type application struct {
	container    di.Container
	serverConfig pkg.Server
}

func CreateApplication(c di.Container, sc pkg.Server) app.Application {
	return &application{c, sc}
}

func (a *application) Run() error {
	userUseCase := a.container.ProvideUserUseCase()

	manager := seeder.CreateSeeder(userUseCase)
	return manager.Seed()
}
