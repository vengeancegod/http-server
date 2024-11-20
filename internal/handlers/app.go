package handlers

import (
	"encoding/json"
	"http-server/internal/entities"
	"http-server/internal/entities/sqlite"
	"log"
	"net/http"
	"strconv"
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

func (app *App) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Enter id", http.StatusBadRequest)
		return
	}
	accountId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Incorecct type", http.StatusBadRequest)
		return
	}
	err = app.Accounts.DeleteAccount(accountId)
	if err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Account deleted successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleDeleteIntegration(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Enter id", http.StatusBadRequest)
		return
	}
	integrationId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Incorrect type", http.StatusBadRequest)
		return
	}
	err = app.Integrations.DeleteIntegration(integrationId)
	if err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Integration deleted successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Enter id", http.StatusBadRequest)
		return
	}

	accountId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Incorrect type of ID", http.StatusBadRequest)
		return
	}

	var acc *entities.Account

	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = app.Accounts.UpdateAccount(accountId, acc.AccessToken, acc.RefreshToken, acc.Expires)
	if err != nil {
		http.Error(w, "Failed to update account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Account updated successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleUpdateIntegration(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(("id"))
	if id == "" {
		http.Error(w, "Enter id", http.StatusBadRequest)
		return
	}

	integrationId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Incorrect type of ID", http.StatusBadRequest)
		return
	}

	var ai *entities.AccountIntegration

	if err := json.NewDecoder(r.Body).Decode(&ai); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = app.Integrations.UpdateIntegration(integrationId, ai.AccountId, ai.SecretKey, ai.ClientId, ai.RedirectURL, ai.AuthenticationCode, ai.AuthorizationCode)
	if err != nil {
		http.Error(w, "Failed to update account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{"status": "Integration updated successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleAccounts)
	mux.HandleFunc("/integrations", app.handleIntegrations)
	mux.HandleFunc("/createAccount", app.handleCreateAccount)
	mux.HandleFunc("/createIntegration", app.handleCreateIntegration)
	mux.HandleFunc("/deleteAccount", app.handleDeleteAccount)
	mux.HandleFunc("/deleteIntegration", app.handleDeleteIntegration)
	mux.HandleFunc("/updateAccount", app.handleUpdateAccount)
	mux.HandleFunc("/updateIntegration", app.handleUpdateIntegration)
	return mux
}
