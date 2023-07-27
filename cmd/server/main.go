package main

import (
	"log"
	"shop-smart-api/internal/app/server"
	"shop-smart-api/pkg"
)

func main() {
	config, err := pkg.CreateConfig()
	if err != nil {
		log.Panic(err.Error())
	}
	db, err := pkg.CreateDatabaseConnection(config.Database)
	if err != nil {
		log.Panic(err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Panic(err.Error())
		}
	}()

	application := server.CreateApplication(db, config.Server)
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
