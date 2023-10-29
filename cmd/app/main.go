package main

import (
	"cookbook/internal/server"
	"cookbook/internal/config"
	"cookbook/internal/usecase"
	"cookbook/internal/storage"
	"log"
)


func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	dishStorage, err := storage.New(appConfig.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	cookInteractor := usecase.NewTaskInteractor(dishStorage)

	handlers := server.NewHandler(cookInteractor)
	srv := new(server.Server)

	if err := srv.Run(appConfig.HostAddr, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}