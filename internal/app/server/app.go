package server

import (
	"database/sql"
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/controller"
	di "shop-smart-api/internal/infrastructure/container"
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

	userService := container.ProvideUserService()
	otpService := container.ProvideOTPService()
	organizationService := container.ProvideOrganizationService()
	transactionService := container.ProvideTransactionService()

	httpServer := controller.CreateServer(
		a.serverConfig,
		otpService,
		userService,
		organizationService,
		transactionService,
	)

	return httpServer.RunServer()
}
