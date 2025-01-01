package account

import (
	"errors"
	"http-server/internal/entities"
	rep "http-server/internal/repository"

	"gorm.io/gorm"
)

var _ rep.AccountRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	return &repository{
		DB: db,
	}, nil
}

func (repo *repository) GetAllContacts() ([]entities.Contacts, error) {
	var contacts []entities.Contacts

	if err := repo.DB.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (repo *repository) Authorization(request entities.AuthRequest) (entities.AuthResponse, error) {
	var account entities.Account

	if err := repo.DB.Where("code = ?", request.Code).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.AuthResponse{}, errors.New(entities.ErrNotFoundAcc)
		}
		return entities.AuthResponse{}, err
	}

	return entities.AuthResponse{
		AccessToken:  account.AccessToken,
		RefreshToken: account.RefreshToken,
		ExpiresIn:    int(account.Expires),
	}, nil
}

func (repo *repository) CreateAccount(account *entities.Account) error {
	return repo.DB.Create(account).Error
}

func (repo *repository) GetAllAccounts() ([]entities.Account, error) {
	var accounts []entities.Account

	if err := repo.DB.Find(&accounts).Error; err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, errors.New(entities.ErrNotFoundAcc)
	}
	return accounts, nil
}

func (repo *repository) DeleteAccount(id int64) error {
	var account entities.Account

	if err := repo.DB.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(entities.ErrNotFoundAcc)
		}
		return err
	}

	if err := repo.DB.Where("account_id = ?", id).Delete(&entities.Contacts{}).Error; err != nil {
		return err
	}

	if err := repo.DB.Where("account_id = ?", id).Delete(&entities.AccountIntegration{}).Error; err != nil {
		return err
	}

	if err := repo.DB.Delete(&account).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetContactsByAccountID(id int64) (entities.Contacts, error) {
	var contacts entities.Contacts

	if err := repo.DB.Where("id = ?", id).First(&contacts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.Contacts{}, errors.New(entities.ErrNotFoundAcc)
		}
		return entities.Contacts{}, err
	}
	return contacts, nil
}

func (repo *repository) GetAccountByID(id int64) (entities.Account, error) {
	var account entities.Account

	if err := repo.DB.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return account, errors.New(entities.ErrNotFoundAcc)
		}
		return account, err
	}

	return account, nil
}
