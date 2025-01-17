package handlers

import (
	"encoding/json"
	"http-server/internal/entities"
	"http-server/internal/infrastructure/beanstalk"
	"http-server/internal/service"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	accountService     service.AccountService
	integrationService service.AccountIntegrationService
	contactsService    service.ContactsService
	unisenderService   service.UnisenderIntegrationService
	beanstalkService   beanstalk.BeanstalkService
}

func NewApp(accountService service.AccountService, integrationService service.AccountIntegrationService, contactsService service.ContactsService,
	unisenderService service.UnisenderIntegrationService, beanstalkService beanstalk.BeanstalkService) *App {
	return &App{accountService: accountService, integrationService: integrationService, contactsService: contactsService, unisenderService: unisenderService,
		beanstalkService: beanstalkService}
}

func (app *App) handleAuthorization(w http.ResponseWriter, r *http.Request) {
	var authRequest entities.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := app.accountService.Authorization(authRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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

	err := app.accountService.CreateAccount(&account)
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

func (app *App) handleCreateIntegration(w http.ResponseWriter, r *http.Request) {
	var integration entities.AccountIntegration

	if err := json.NewDecoder(r.Body).Decode(&integration); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
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

	response := entities.Response{
		Status:  "success",
		Message: entities.AccountDelete,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleAccountByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	accountID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	account, err := app.accountService.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (app *App) handleGetContactByAccountID(w http.ResponseWriter, r *http.Request) {
	account_id := r.URL.Query().Get("account_id")

	accountID, err := strconv.ParseInt(account_id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contacts, err := app.contactsService.GetContactsByAccountID(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func (app *App) handleGetAndSaveContactsByAccountID(w http.ResponseWriter, r *http.Request) {
	account_id := r.URL.Query().Get("account_id")

	accountID, err := strconv.ParseInt(account_id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contacts, err := app.contactsService.GetAndSaveContactsByAccountID(int64(accountID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func (app *App) handleGetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := app.contactsService.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *App) handleDeleteContact(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	contactID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.contactsService.DeleteContact(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := entities.Response{
		Status:  "success",
		Message: entities.AccountDelete,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) handleGetUnisenderKey(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		unisenderKey, err := app.unisenderService.GetUnisenderKey()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(unisenderKey); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		unisenderKey := r.FormValue("unisender_key")
		accountIDString := r.FormValue("account_id")
		accountID, err := strconv.ParseInt(accountIDString, 10, 64)

		unisenderIntegration := &entities.UnisenderIntegration{
			UnisenderKey: unisenderKey,
			AccountID:    accountID,
		}

		err = app.unisenderService.SaveUnisenderKey(unisenderIntegration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		task := entities.SynchronizationTask{
			UnisenderKey: unisenderKey,
			AccountID:    accountID,
		}

		taskData, err := json.Marshal(task)
		if err != nil {
			http.Error(w, entities.ErrMarshalTask, http.StatusInternalServerError)
			return
		}

		err = app.beanstalkService.PutTask(taskData)
		if err != nil {
			http.Error(w, entities.ErrPutTask, http.StatusInternalServerError)
			return
		}

		contacts, err := app.contactsService.GetAllContacts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = app.contactsService.SendToUnisender(contacts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(unisenderIntegration); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (app *App) handleContactsHook(w http.ResponseWriter, r *http.Request) {
	var task *entities.ContactsTask

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteContacts, ok := r.Form["contacts[delete][0][id]"]
	if ok && len(deleteContacts) > 0 {
		contactIDStr := deleteContacts[0]
		log.Printf("Delete contact with ID: %s", contactIDStr)

		contactID, err := strconv.ParseInt(contactIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Incorrect contactID", http.StatusBadRequest)
			return
		}

		task = &entities.ContactsTask{
			Action:    "delete",
			ContactID: contactID,
		}

		taskData, err := json.Marshal(task)
		if err != nil {
			http.Error(w, entities.ErrMarshalTask, http.StatusInternalServerError)
			return
		}

		if err := app.beanstalkService.PutTask(taskData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	addContact, ok := r.Form["contacts[add][0][id]"]
	if ok && len(addContact) > 0 {
		contactData := make(map[string]string)
		contactData["id"] = addContact[0]
		contactData["name"] = r.Form.Get("contacts[add][0][name]")
		contactData["email"] = r.Form.Get("contacts[add][0][custom_fields][0][values][0][value]")

		accountIDStr := r.Form.Get("contacts[add][0][account_id]")
		if accountIDStr == "" {
			accountIDStr = r.Form.Get("account[id]")
		}

		accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid account ID", http.StatusBadRequest)
			return
		}

		var contact entities.Contacts
		contact.ID, _ = strconv.ParseInt(contactData["id"], 10, 64)
		contact.Name = contactData["name"]
		contact.Email = contactData["email"]
		contact.AccountID = accountID

		task = &entities.ContactsTask{
			Action:    "add",
			ContactID: contact.ID,
			Name:      contact.Name,
			Email:     contact.Email,
			AccountID: accountID,
		}

		taskData, err := json.Marshal(task)
		if err != nil {
			http.Error(w, entities.ErrMarshalTask, http.StatusInternalServerError)
			return
		}

		if err := app.beanstalkService.PutTask(taskData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Contact with ID %d сreated", contact.ID)
		w.WriteHeader(http.StatusOK)
		return
	}

	updateContacts, ok := r.Form["contacts[update][0][id]"]
	if ok && len(updateContacts) > 0 {
		contactIDStr := updateContacts[0]
		contactID, err := strconv.ParseInt(contactIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Incorrect contactID", http.StatusBadRequest)
			return
		}

		contact, err := app.contactsService.GetContactByID(contactID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		contact.Name = r.Form.Get("contacts[update][0][name]")
		contact.Email = r.Form.Get("contacts[update][0][custom_fields][0][values][0][value]")

		task = &entities.ContactsTask{
			Action:    "update",
			ContactID: contactID,
			Name:      contact.Name,
			Email:     contact.Email,
		}

		taskData, err := json.Marshal(task)
		if err != nil {
			http.Error(w, entities.ErrMarshalTask, http.StatusInternalServerError)
			return
		}

		if err := app.beanstalkService.PutTask(taskData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Contact with ID %d updated", contact.ID)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleAccounts)
	mux.HandleFunc("/integrations", app.handleIntegrations)
	mux.HandleFunc("/createAccount", app.handleCreateAccount)
	mux.HandleFunc("/createIntegration", app.handleCreateIntegration)
	mux.HandleFunc("/deleteIntegration", app.handleDeleteIntegration)
	mux.HandleFunc("/updateIntegration", app.handleUpdateIntegration)
	mux.HandleFunc("/auth", app.handleAuthorization)
	mux.HandleFunc("/deleteAccount", app.handleDeleteAccount)
	mux.HandleFunc("/getAccountByID", app.handleAccountByID)
	mux.HandleFunc("/getContactsByAccountID", app.handleGetContactByAccountID)
	mux.HandleFunc("/getContactsFromAPI", app.handleGetAndSaveContactsByAccountID)
	mux.HandleFunc("/getContacts", app.handleGetAllContacts)
	mux.HandleFunc("/deleteContact", app.handleDeleteContact)
	mux.HandleFunc("/getUnisenderKey", app.handleGetUnisenderKey)
	mux.HandleFunc("/contactsHook", app.handleContactsHook)
	return mux
}
