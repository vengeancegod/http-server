package main

import (
<<<<<<< HEAD
	"http-server/internal/database"
	"http-server/internal/entities/sqlite"
	"http-server/internal/handlers"
	"log"
	"net/http"
)

// type app struct {
// 	accounts     *sqlite.AccountEntity
// 	integrations *sqlite.AccountIntegrationEntity
// }

// func (app *app) handleAccounts(w http.ResponseWriter, r *http.Request) {
// 	accounts, err := app.accounts.GetAll()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(accounts); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func (app *app) handleIntegrations(w http.ResponseWriter, r *http.Request) {
// 	integrations, err := app.integrations.GetAllIntegrations()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	log.Println("Integrations:", integrations)
// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(integrations); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func (app *app) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
// 	var account entities.Account
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&account)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	err = app.accounts.CreateAccount(account.AccessToken, account.RefreshToken, account.Expires)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Account created successfully"}); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func (app *app) handleCreateIntegration(w http.ResponseWriter, r *http.Request) {
// 	var integration entities.AccountIntegration
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&integration)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	err = app.integrations.CreateIntegration(integration.AccountId, integration.SecretKey, integration.ClientId, integration.RedirectURL, integration.AuthenticationCode, integration.AuthorizationCode)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Integration created successfully"}); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func (app *app) routes() http.Handler {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", app.handleAccounts)
// 	mux.HandleFunc("/integrations", app.handleIntegrations)
// 	mux.HandleFunc("/createAccount", app.handleCreateAccount)
// 	mux.HandleFunc("/createIntegration", app.handleCreateIntegration)
// 	return mux
// }

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	app := &handlers.App{
		Accounts:     &sqlite.AccountEntity{DB: db},
		Integrations: &sqlite.AccountIntegrationEntity{DB: db},
	}

	server := http.Server{
=======
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
	db, err := sql.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	accountRepo, err := account.NewRepository(db)
	if err != nil {
		errors.New("Error create repository")
	}

	integrationRepo := integration.NewRepository()
	if err != nil {
		errors.New("Error create repository")
	}
	contactsRepo, err := contacts.NewRepository(db)
	if err != nil {
		errors.New("Error create repository")
	}
	unisenderRepo, err := unisender_integration.NewRepository(db)
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

	err = sql.CreateMigration(db)
	if err != nil {
		errors.New("Error create migration")
	}

	httpServer := http.Server{
>>>>>>> feature/SCHOOL-1312
		Addr:    ":8080",
		Handler: app.Routes(),
	}

<<<<<<< HEAD
	log.Println(("Listining on :8080"))
	server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
=======
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
>>>>>>> feature/SCHOOL-1312
}
