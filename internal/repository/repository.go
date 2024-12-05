package repository

import "http-server/internal/entities"

type AccountRepository interface {
	CreateAccount(account entities.Account) error
	GetAllAccounts() ([]entities.Account, error)
	// UpdateAccount(id int64, account entities.Account) error
	// DeleteAccount(id int64) error
	Authorization(request entities.AuthRequest) (entities.AuthResponse, error)
}

type AccountIntegrationRepository interface {
	CreateIntegration(integration entities.AccountIntegration) error
	GetAllIntegrations() ([]entities.AccountIntegration, error)
	UpdateIntegration(id int64, integration entities.AccountIntegration) error
	DeleteIntegration(id int64) error
}

type ContactsRepository interface {
	GetAllContacts() ([]entities.Contacts, error)
}
