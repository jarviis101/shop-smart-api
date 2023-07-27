package main

import (
	_ "shop-smart-api/internal/app/console"
)

func main() {
	//config, err := pkg.CreateConfig()
	//if err != nil {
	//	log.Panic(err.Error())
	//}
	//
	//db := pkg.CreateDatabaseConnection(context.Background(), config.Database)
	//defer func() {
	//	if err := pkg.CloseConnection(db.Client()); err != nil {
	//		log.Panic(err.Error())
	//	}
	//}()
	//
	//container := di.CreateContainer(db, config.Server)
	//application := console.CreateApplication(container, config.Server)
	//if err := application.Run(); err != nil {
	//	log.Panic(err.Error())
	//}
}
