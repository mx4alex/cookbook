package main

import (
	"cookbook/internal/server"
	"cookbook/internal/config"
	"cookbook/internal/usecase"
	"cookbook/internal/storage"
	"log"
	"os"
	"os/signal"
	"context"
)


func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.New(appConfig.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(db)
	services := usecase.NewService(storage)

	handlers := server.NewHandler(services)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(appConfig.HostAddr, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("error occured on server shutting down: %s", err.Error())
	}
}