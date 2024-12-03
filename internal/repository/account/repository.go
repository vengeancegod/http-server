package account

import (
	"errors"
	"http-server/internal/entities"
	rep "http-server/internal/repository"
	repoModel "http-server/internal/repository/account/model"
	"strconv"
	"sync"
)

var _ rep.AccountRepository = (*repository)(nil)

type repository struct {
	data     map[string]*repoModel.Account
	contacts map[string]*repoModel.Contacts
	mu       sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data:     make(map[string]*repoModel.Account),
		contacts: make(map[string]*repoModel.Contacts),
	}
}

func (repo *repository) GetAllContacts() ([]entities.Contacts, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.contacts) == 0 {
		return nil, errors.New(entities.ErrNotFoundContacts)
	}

	var contacts []entities.Contacts

	for _, contact := range repo.contacts {
		contacts = append(contacts, entities.Contacts(*contact))
	}

	return contacts, nil
}

func (repo *repository) Authorization(request entities.AuthRequest) (entities.AuthResponse, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	account, exists := repo.data[request.Code]
	if !exists {
		return entities.AuthResponse{}, errors.New(entities.ErrNotFoundAcc)
	}

	authResponse := entities.AuthResponse{
		TokenType:    "Bearer",
		ExpiresIn:    int(account.Expires),
		AccessToken:  account.AccessToken,
		RefreshToken: account.RefreshToken,
	}

	return authResponse, nil
}

func (repo *repository) CreateAccount(account entities.Account) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	accountID := strconv.FormatInt(account.ID, 10)
	if _, exists := repo.data[accountID]; exists {
		return nil
	}

	repo.data[accountID] = &repoModel.Account{
		ID:           account.ID,
		AccessToken:  account.AccessToken,
		RefreshToken: account.RefreshToken,
		Expires:      account.Expires,
	}

	return nil
}

func (repo *repository) GetAllAccounts() ([]entities.Account, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.data) == 0 {
		return nil, errors.New(entities.ErrNotFoundAcc)
	}

	var accounts []entities.Account

	for _, account := range repo.data {
		accounts = append(accounts, entities.Account(*account))
	}

	return accounts, nil
}

func (repo *repository) UpdateAccount(id int64, account entities.Account) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	accountID := strconv.FormatInt(account.ID, 10)
	if _, exists := repo.data[accountID]; !exists {
		return errors.New(entities.ErrUpdateAcc)
	}

	repo.data[accountID] = &repoModel.Account{
		ID:           account.ID,
		AccessToken:  account.AccessToken,
		RefreshToken: account.RefreshToken,
		Expires:      account.Expires,
	}
	return nil
}

func (repo *repository) DeleteAccount(id int64) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	accountID := strconv.FormatInt(id, 10)
	if _, exists := repo.data[accountID]; !exists {
		return errors.New(entities.ErrFailedDelete)
	}

	delete(repo.data, accountID)
	return nil

}
