package contacts

import (
	"errors"
	"http-server/internal/entities"
	rep "http-server/internal/repository"
	repoModel "http-server/internal/repository/contacts/model"
	"sync"
)

var _ rep.ContactsRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.Contacts
	mu   sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.Contacts),
	}
}

func (repo *repository) GetAllContacts() ([]entities.Contacts, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.data) == 0 {
		return nil, errors.New(entities.ErrNotFoundContacts)
	}

	var contacts []entities.Contacts

	for _, contact := range repo.data {
		contacts = append(contacts, entities.Contacts{
			Name:  contact.Name,
			Email: contact.Email,
		})
	}

	return contacts, nil
}
