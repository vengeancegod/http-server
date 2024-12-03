package handlers

import (
	"encoding/json"
	"http-server/internal/entities"
	"http-server/internal/service"
	"net/http"
	"strconv"
)

type App struct {
	accountService     service.AccountService
	integrationService service.AccountIntegrationService
}

func NewApp(accountService service.AccountService, integrationService service.AccountIntegrationService) *App {
	return &App{accountService: accountService, integrationService: integrationService}
}

var lastAccountID int64
var lastIntegrationID int64

func (app *App) handleGetContacts(w http.ResponseWriter, _ *http.Request) {
	contacts, err := app.accountService.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Contect-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleAuthorization(w http.ResponseWriter, r *http.Request) {
	var authRequest entities.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := app.accountService.Authorization(authRequest)
	if err != nil {
		http.Error(w, "Authorization failed", http.StatusUnauthorized)
		return
	}

	if err := app.accountService.CreateAccount(authResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(authResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (app *App) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account entities.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastAccountID++
	account.ID = lastAccountID

	err := app.accountService.CreateAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	response := entities.Response{
		Status:  "success",
		Message: entities.AccountCreate,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleAccounts(w http.ResponseWriter, _ *http.Request) {

	account, err := app.accountService.GetAllAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	accountID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, entities.ErrIncorrectType, http.StatusBadRequest)
		return
	}

	var account entities.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account.ID = accountID

	err = app.accountService.UpdateAccount(accountID, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := entities.Response{
		Status:  "success",
		Message: entities.AccountUpdate,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	accountID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.accountService.DeleteAccount(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := entities.Response{
		Status:  "success",
		Message: entities.AccountDelete,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleCreateIntegration(w http.ResponseWriter, r *http.Request) {
	var integration entities.AccountIntegration

	if err := json.NewDecoder(r.Body).Decode(&integration); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastIntegrationID++
	integration.ID = lastIntegrationID

	err := app.integrationService.CreateIntegration(integration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := entities.Response{
		Status:  "success",
		Message: entities.IntegrationCreate,
	}

	w.Header().Set("Content-Type", "applications/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleIntegrations(w http.ResponseWriter, _ *http.Request) {

	accountIntegration, err := app.integrationService.GetAllIntegrations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accountIntegration); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleUpdateIntegration(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	integrationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, entities.ErrIncorrectType, http.StatusBadRequest)
	}

	var integration entities.AccountIntegration
	if err := json.NewDecoder(r.Body).Decode(&integration); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	integration.ID = integrationID

	err = app.integrationService.UpdateIntegration(integrationID, integration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	response := entities.Response{
		Status:  "success",
		Message: entities.IntegrationUpdate,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleDeleteIntegration(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	integrationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.integrationService.DeleteIntegration(integrationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := entities.Response{
		Status:  "success",
		Message: entities.IntegrationDelete,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
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
	mux.HandleFunc("/auth", app.handleAuthorization)
	mux.HandleFunc("/getContacts", app.handleGetContacts)
	return mux
}
