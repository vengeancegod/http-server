package repository

import "http-server/internal/entities"

type AccountRepository interface {
	CreateAccount(account *entities.Account) error
	GetAllAccounts() ([]entities.Account, error)
	Authorization(request entities.AuthRequest) (entities.AuthResponse, error)
	DeleteAccount(id int64) error
	GetAccountByID(id int64) (entities.Account, error)
}

type AccountIntegrationRepository interface {
	CreateIntegration(integration entities.AccountIntegration) error
	GetAllIntegrations() ([]entities.AccountIntegration, error)
	UpdateIntegration(id int64, integration entities.AccountIntegration) error
	DeleteIntegration(id int64) error
}

type ContactsRepository interface {
	GetAllContacts() ([]entities.Contacts, error)
	GetContactsByAccountID(account_id int64) ([]entities.Contacts, error)
	CreateContacts(contacts []entities.Contacts) error
	DeleteContact(id int64) error
}

type UnisenderIntegrationRepository interface {
	GetUnisenderKey() ([]entities.UnisenderIntegration, error)
	SaveUnisenderKey(unisenderKey *entities.UnisenderIntegration) error
}
