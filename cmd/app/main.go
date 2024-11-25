package main

import (
	"http-server-fixed/internal/handlers"
	"http-server-fixed/internal/repository/account"
	"http-server-fixed/internal/repository/integration"
	aService "http-server-fixed/internal/service/account"
	iService "http-server-fixed/internal/service/integration"
	"log"
	"net/http"
)

func main() {
	accountRepo := account.NewRepository()
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
