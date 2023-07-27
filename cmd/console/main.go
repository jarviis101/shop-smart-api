package main

import (
	"log"
	"shop-smart-api/internal/app/console"
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

	application := console.CreateApplication(db, config.Server)
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
