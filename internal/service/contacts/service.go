package contacts

import (
	"http-server/internal/repository"
)

type Service struct {
	contactsRepository  repository.ContactsRepository
	accountRepositry    repository.AccountRepository
	unisenderRepository repository.UnisenderIntegrationRepository
}

func NewService(contactsRepository repository.ContactsRepository, accountRepository repository.AccountRepository,
	unisenderRepository repository.UnisenderIntegrationRepository) *Service {
	return &Service{
		contactsRepository:  contactsRepository,
		accountRepositry:    accountRepository,
		unisenderRepository: unisenderRepository,
	}
}
