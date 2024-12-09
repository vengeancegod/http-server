package contacts

import (
	"http-server/internal/repository"
)

type Service struct {
	contactsRepository repository.ContactsRepository
	accountRepositry   repository.AccountRepository
}

func NewService(contactsRepository repository.ContactsRepository, accountRepository repository.AccountRepository) *Service {
	return &Service{
		contactsRepository: contactsRepository,
		accountRepositry:   accountRepository,
	}
}
