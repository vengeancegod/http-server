package main

import (
	"errors"
	unsubService "http-server/internal/api/unsubscribe"
	"http-server/internal/handlers"
	bService "http-server/internal/infrastructure/beanstalk"
	"http-server/internal/infrastructure/database/sql"
	"http-server/internal/repository/account"
	"http-server/internal/repository/contacts"
	"http-server/internal/repository/integration"
	"http-server/internal/repository/unisender_integration"
	aService "http-server/internal/service/account"
	cService "http-server/internal/service/contacts"
	iService "http-server/internal/service/integration"
	uiService "http-server/internal/service/unisender_integration"
	desc "http-server/pkg/unsubscribe"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
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
	beanstalkService, err := bService.NewService()
	if err != nil {
		log.Fatal("Error initializing beanstalk service: ", err)
	}
	app := handlers.NewApp(accountService, integrationService, contactsService, unisenderService, beanstalkService)
	db := accountRepo.DB
	err = sql.CreateMigration(db)
	if err != nil {
		errors.New("Error create migration")
	}

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}

	go func() {
		log.Println(("Listen on :8080"))
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	grpcServer := grpc.NewServer()

	unsubscribeService := unsubService.NewImplementation(accountService)
	desc.RegisterUnsubscribeServer(grpcServer, unsubscribeService)

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println(("Listen on :8081"))
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
