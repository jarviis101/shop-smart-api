package server

import (
	"database/sql"
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/controller"
	di "shop-smart-api/internal/infrastructure/container"
	"shop-smart-api/pkg"
)

type application struct {
	database *sql.DB
	appCfg   *pkg.AppConfig
}

func CreateApplication(db *sql.DB, a *pkg.AppConfig) app.Application {
	return &application{db, a}
}

func (a *application) Run() error {
	container := di.CreateContainer(a.database, a.appCfg.Server, a.appCfg.Mailer)

	userService := container.ProvideUserService()
	otpService := container.ProvideOTPService()
	organizationService := container.ProvideOrganizationService()
	transactionService := container.ProvideTransactionService()

	httpServer := controller.CreateServer(
		a.appCfg.Server,
		otpService,
		userService,
		organizationService,
		transactionService,
	)

	return httpServer.RunServer()
}
