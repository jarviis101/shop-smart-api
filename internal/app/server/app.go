package server

import (
	"shop-smart-api/internal/app"
	"shop-smart-api/internal/app/di"
	"shop-smart-api/internal/controller/http"
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
	otpUseCase := a.container.ProvideOTPUseCase()
	httpServer := http.CreateServer(a.serverConfig, userUseCase, otpUseCase)

	return httpServer.RunServer()
}
