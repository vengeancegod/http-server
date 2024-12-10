package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"http-server/internal/entities"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const contactsURL = "https://volkovkirill.amocrm.ru/api/v4/contacts"
const unisenderURL = "https://api.unisender.com/ru/api/importContacts"

func (s *Service) SendToUnisender(contacts []entities.Contacts) error {

	unisenderKeys, err := s.unisenderRepository.GetUnisenderKey()
	if err != nil {
		return errors.New(entities.ErrGetUnisenderKey)
	}
	unisenderKey := unisenderKeys[0].UnisenderKey

	fields := []string{"email", "Name"}
	params := url.Values{}
	params.Add("format", "json")
	params.Add("api_key", unisenderKey)

	for i, field := range fields {
		params.Add(fmt.Sprintf("field_names[%d]", i), field)
	}

	for i, contact := range contacts {
		params.Add(fmt.Sprintf("data[%d][0]", i), contact.Email)
		params.Add(fmt.Sprintf("data[%d][1]", i), contact.Name)
	}
	log.Printf("Request to Unisender: %s?%s", unisenderURL, params.Encode())
	resp, err := http.Post(unisenderURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.New(entities.ErrReadBody)
		}
		log.Printf("Unisender response body: %s", string(body))
		return nil
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(entities.ErrReadBody)
	}

	log.Printf("Unisender response: %s", string(responseBody))
	return nil
}

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
