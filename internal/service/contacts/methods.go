package contacts

import (
	"encoding/json"
	"errors"
	"http-server/internal/entities"
	"io"
	"log"
	"net/http"
)

const contactsURL = "https://volkovkirill.amocrm.ru/api/v4/contacts"

func (s *Service) GetContactsByAccountID(accountID int64) ([]entities.Contacts, error) {
	contacts, err := s.contactsRepository.GetContactsByAccountID(accountID)
	if err != nil {
		return nil, errors.New(entities.ErrNotFoundContacts)
	}

	return contacts, nil
}

func (s *Service) GetAndSaveContactsByAccountID(accountID int64) ([]entities.Contacts, error) {
	contacts, err := s.getContactsFromAPI(accountID)
	if err != nil {
		return nil, err
	}

	if err := s.contactsRepository.CreateContacts(contacts); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *Service) getContactsFromAPI(accountID int64) ([]entities.Contacts, error) {
	var allContacts []entities.Contacts

	account, err := s.accountRepositry.GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}

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
			ID:        contact.ID,
			Name:      contact.Name,
			Email:     email,
			AccountID: accountID,
		})
	}
	return allContacts, nil
}

func (s *Service) GetAllContacts() ([]entities.Contacts, error) {
	contacts, err := s.contactsRepository.GetAllContacts()
	if err != nil {
		return nil, errors.New(entities.ErrNotFoundContacts)
	}
	return contacts, nil
}

func (s *Service) DeleteContact(id int64) error {
	err := s.contactsRepository.DeleteContact(id)
	if err != nil {
		return err
	}

	return nil
}
