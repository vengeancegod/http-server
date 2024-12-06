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

const contactsURL = "https://volkovkirill.amocrm.ru/api/v4/contacts"

func (s *Service) GetAllContacts() ([]entities.Contacts, error) {

	accounts, err := s.accountRepository.GetAllAccounts()
	if err != nil {
		return nil, errors.New(entities.ErrNotFoundAcc)
	}

	var allContacts []entities.Contacts

	for _, account := range accounts {
		req, err := http.NewRequest("GET", contactsURL, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+account.AccessToken)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, errors.New(entities.ErrNotFoundContacts)
		}

		defer resp.Body.Close()

		log.Println("Status Code:", resp.StatusCode)
		log.Println("Using AccessToken:", account.AccessToken)

		if resp.StatusCode != http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, errors.New(entities.ErrReadBody)
			}
			log.Printf("API Response: %s\n", string(body))
			return nil, errors.New(entities.ErrGetContacts)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(entities.ErrReadBody)
		}
		log.Printf("API Response Body: %s\n", string(body))

		var apiResponse entities.ContactAPIResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			return nil, errors.New(entities.ErrParse)
		}
		for _, contact := range apiResponse.Embedded.Contacts {

			var email string
			for _, field := range contact.CustomFieldsValues {
				if field.FieldName == "Email" && len(field.Values) > 0 {
					email = field.Values[0].Value
					break
				}
			}

			allContacts = append(allContacts, entities.Contacts{
				Name:  contact.Name,
				Email: email,
			})
		}

	}
	return allContacts, nil
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

	if err := s.accountRepository.CreateAccount(account); err != nil {
		log.Printf("Error creating account: %v", err)
		return entities.Account{}, errors.New(entities.ErrCreateAcc)
	}

	return account, nil
}

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
