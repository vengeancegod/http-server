package handlers

import (
	"encoding/json"
	"http-server/internal/entities"
	"http-server/internal/entities/sqlite"
	"log"
	"net/http"
)

type App struct {
	Accounts     *sqlite.AccountEntity
	Integrations *sqlite.AccountIntegrationEntity
}

func (app *App) handleAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.Accounts.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleIntegrations(w http.ResponseWriter, r *http.Request) {
	integrations, err := app.Integrations.GetAllIntegrations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Integrations:", integrations)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(integrations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account entities.Account
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&account)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = app.Accounts.CreateAccount(account.AccessToken, account.RefreshToken, account.Expires)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Account created successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleCreateIntegration(w http.ResponseWriter, r *http.Request) {
	var integration entities.AccountIntegration
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&integration)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = app.Integrations.CreateIntegration(integration.AccountId, integration.SecretKey, integration.ClientId, integration.RedirectURL, integration.AuthenticationCode, integration.AuthorizationCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Integration created successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleAccounts)
	mux.HandleFunc("/integrations", app.handleIntegrations)
	mux.HandleFunc("/createAccount", app.handleCreateAccount)
	mux.HandleFunc("/createIntegration", app.handleCreateIntegration)
	return mux
}
