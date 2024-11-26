package account

import (
	"errors"
	"http-server/internal/entities"
)

func (s *Service) CreateAccount(account entities.Account) error {
	err := s.accountRepository.CreateAccount(account)
	if err != nil {
		return errors.New(entities.ErrCreateAcc)
	}
	return nil
}

func (s *Service) GetAllAccounts() ([]entities.Account, error) {
	account, err := s.accountRepository.GetAllAccounts()
	if err != nil {
		return nil, errors.New(entities.ErrNotFoundAcc)
	}
	return account, nil
}

func (s *Service) UpdateAccount(id int64, account entities.Account) error {
	err := s.accountRepository.UpdateAccount(id, account)

	if err != nil {
		return errors.New(entities.ErrUpdateAcc)
	}

	return nil
}

func (s *Service) DeleteAccount(id int64) error {
	err := s.accountRepository.DeleteAccount(id)

	if err != nil {
		return errors.New(entities.ErrFailedDelete)
	}

	return nil
}
