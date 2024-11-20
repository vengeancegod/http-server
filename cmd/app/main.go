package main

import (
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
		Addr:    ":8080",
		Handler: app.Routes(),
	}

	log.Println(("Listining on :8080"))
	server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
