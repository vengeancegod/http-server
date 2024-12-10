package main

import (
	"errors"
	"http-server/internal/handlers"
	"http-server/internal/infrastructure/database/sql"
	"http-server/internal/repository/account"
	"http-server/internal/repository/contacts"
	"http-server/internal/repository/integration"
	"http-server/internal/repository/unisender_integration"
	aService "http-server/internal/service/account"
	cService "http-server/internal/service/contacts"
	iService "http-server/internal/service/integration"
	uiService "http-server/internal/service/unisender_integration"
	"log"
	"net/http"
)

func main() {
	accountRepo, err := account.NewRepository()
	if err != nil {
		errors.New("Error create repository")
	}

	integrationRepo := integration.NewRepository()
	if err != nil {
		errors.New("Error create repository")
	}
	contactsRepo, err := contacts.NewRepository()
	if err != nil {
		errors.New("Error create repository")
	}
	unisenderRepo, err := unisender_integration.NewRepository()
	if err != nil {
		errors.New("Errors create repository")
	}
	unisenderService := uiService.NewService(unisenderRepo)
	contactsService := cService.NewService(contactsRepo, accountRepo, unisenderRepo)
	accountService := aService.NewService(accountRepo)
	integrationService := iService.NewService(integrationRepo)
	app := handlers.NewApp(accountService, integrationService, contactsService, unisenderService)
	db := accountRepo.DB
	err = sql.CreateMigration(db)
	if err != nil {
		errors.New("Error create migration")
	}

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
