package service

import "http-server/internal/entities"

type AccountService interface {
	CreateAccount(account entities.Account) error
	GetAllAccounts() ([]entities.Account, error)
	Authorization(request entities.AuthRequest) (entities.Account, error)
	GetAllContacts() ([]entities.Contacts, error)
}

type AccountIntegrationService interface {
	CreateIntegration(integration entities.AccountIntegration) error
	GetAllIntegrations() ([]entities.AccountIntegration, error)
	UpdateIntegration(id int64, integration entities.AccountIntegration) error
	DeleteIntegration(id int64) error
}
