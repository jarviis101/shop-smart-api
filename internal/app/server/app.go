package server

import (
	"database/sql"
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/controller/http"
	"shop-smart-api/internal/infrastructure/container"
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
	di := container.CreateContainer(a.database, a.serverConfig)

	userUseCase := di.ProvideUserUseCase()
	otpUseCase := di.ProvideOTPUseCase()
	httpServer := http.CreateServer(a.serverConfig, userUseCase, otpUseCase)

	return httpServer.RunServer()
}
