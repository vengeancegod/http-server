package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"http-server/internal/entities"
	"io"
	"log"
	"net/http"
)

const authURL = "https://volkovkirill.amocrm.ru/oauth2/access_token"

func (s *Service) GetAccountByID(accountID int64) (entities.Account, error) {
	account, err := s.accountRepository.GetAccountByID(accountID)
	if err != nil {
		return entities.Account{}, errors.New(entities.ErrNotFoundAcc)
	}

	return account, nil
}

func (s *Service) Authorization(request entities.AuthRequest) (entities.Account, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return entities.Account{}, err
	}

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return entities.Account{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return entities.Account{}, err
		}

		log.Printf("API Error: response body: %s\n", string(body))

		return entities.Account{}, errors.New(entities.ErrAuthFailed)
	}

	var authResponse entities.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return entities.Account{}, err
	}

	account := entities.Account{
		AccessToken:  authResponse.AccessToken,
		RefreshToken: authResponse.RefreshToken,
		Expires:      int64(authResponse.ExpiresIn),
	}

	if err := s.accountRepository.CreateAccount(&account); err != nil {
		log.Printf("Error creating account: %v", err)
		return entities.Account{}, errors.New(entities.ErrCreateAcc)
	}

	return account, nil
}

func (s *Service) CreateAccount(account *entities.Account) error {
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

func (s *Service) DeleteAccount(id int64) error {
	err := s.accountRepository.DeleteAccount(id)

	if err != nil {
		return errors.New(entities.ErrFailedDelete)
	}

	return nil
}
