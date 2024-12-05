package main

import (
	"errors"
	"http-server/internal/handlers"
	"http-server/internal/repository/account"
	"http-server/internal/repository/integration"
	aService "http-server/internal/service/account"
	iService "http-server/internal/service/integration"
	"log"
	"net/http"
)

func main() {
	accountRepo, err := account.NewRepository()
	if err != nil {
		errors.New("Error create repository")
	}
	integrationRepo := integration.NewRepository()
	accountService := aService.NewService(accountRepo)
	integrationService := iService.NewService(integrationRepo)
	app := handlers.NewApp(accountService, integrationService)

	server := http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}

	log.Println(("Listen on :8080"))
	server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
