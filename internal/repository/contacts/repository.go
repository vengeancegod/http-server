package contacts

import (
	"errors"
	"http-server/internal/entities"
	rep "http-server/internal/repository"
	"log"

	"gorm.io/gorm"
)

var _ rep.ContactsRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	return &repository{
		DB: db,
	}, nil
}

func (repo *repository) CreateContacts(contacts []entities.Contacts) error {
	for _, contact := range contacts {
		log.Printf("Save contact: %+v", contact)
		if err := repo.DB.Create(&contacts).Error; err != nil {
			log.Printf("Error saving contacts: %v", err)
			return err
		}
	}
	return nil
}

func (repo *repository) GetAllContacts() ([]entities.Contacts, error) {
	var contacts []entities.Contacts

	if err := repo.DB.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (repo *repository) GetContactsByAccountID(account_id int64) ([]entities.Contacts, error) {
	var contacts []entities.Contacts

	if err := repo.DB.Where("account_id = ?", account_id).Find(&contacts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(entities.ErrNotFoundContacts)
		}
		return nil, err
	}
	return contacts, nil
}

func (repo *repository) DeleteContact(id int64) error {

	var contact entities.Contacts

	if err := repo.DB.First(&contact, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(entities.ErrNotFoundContacts)
		}
		return err
	}

	if err := repo.DB.Delete(&contact).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) UpdateContact(contact entities.Contacts) error {
	if err := repo.DB.Model(&entities.Contacts{}).Where("id = ?", contact.ID).Updates(contact).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) CreateContact(contact entities.Contacts) error {
	log.Printf("Save contact: %+v", contact)
	if err := repo.DB.Create(&contact).Error; err != nil {
		log.Printf("Error saving contact: %v", err)
		return err
	}
	return nil
}

func (repo *repository) GetContactByID(id int64) (entities.Contacts, error) {
	var contact entities.Contacts
	if err := repo.DB.First(&contact, id).Error; err != nil {
		return entities.Contacts{}, err
	}
	return contact, nil
}
