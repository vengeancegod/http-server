package account

import (
	"errors"
	"http-server-fixed/internal/entities"
	rep "http-server-fixed/internal/repository"
	repoModel "http-server-fixed/internal/repository/account/model"
	"strconv"
	"sync"
)

var _ rep.AccountRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.Account
	mu   sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.Account),
	}
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

	// account.Id = id
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
