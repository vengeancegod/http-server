package service

import "http-server/internal/entities"

type AccountService interface {
	CreateAccount(account *entities.Account) error
	GetAllAccounts() ([]entities.Account, error)
	Authorization(request entities.AuthRequest) (entities.Account, error)
	DeleteAccount(id int64) error
	GetAccountByID(id int64) (entities.Account, error)
}

type AccountIntegrationService interface {
	CreateIntegration(integration entities.AccountIntegration) error
	GetAllIntegrations() ([]entities.AccountIntegration, error)
	UpdateIntegration(id int64, integration entities.AccountIntegration) error
	DeleteIntegration(id int64) error
}

type ContactsService interface {
	GetContactsByAccountID(accountID int64) ([]entities.Contacts, error)
	GetAndSaveContactsByAccountID(accountID int64) ([]entities.Contacts, error)
	GetAllContacts() ([]entities.Contacts, error)
	DeleteContact(id int64) error
	SendToUnisender(contacts []entities.Contacts) error
	UpdateContact(contact entities.Contacts) error
	CreateContact(contact entities.Contacts) error
	CreateContacts(contacts []entities.Contacts) error
	GetContactByID(id int64) (entities.Contacts, error)
}

type UnisenderIntegrationService interface {
	GetUnisenderKey() ([]entities.UnisenderIntegration, error)
	SaveUnisenderKey(unisenderKey *entities.UnisenderIntegration) error
}
