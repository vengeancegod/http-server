package integration

import (
	"errors"
	"http-server/internal/entities"
	"http-server/internal/infrastructure/database/sql"
	rep "http-server/internal/repository"

	"gorm.io/gorm"
)

var _ rep.AccountIntegrationRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository() *repository {
	db, err := sql.InitDB()
	if err != nil {
		return nil
	}
	return &repository{
		DB: db,
	}
}

func (repo *repository) CreateIntegration(integration entities.AccountIntegration) error {
	return repo.DB.Create(&integration).Error
}

func (repo *repository) GetAllIntegrations() ([]entities.AccountIntegration, error) {
	var integrations []entities.AccountIntegration

	if err := repo.DB.Find(&integrations).Error; err != nil {
		return nil, err
	}

	if len(integrations) == 0 {
		return nil, errors.New(entities.ErrNotFoundInt)
	}

	return integrations, nil
}

func (repo *repository) UpdateIntegration(id int64, integration entities.AccountIntegration) error {
	var integrations entities.AccountIntegration
	if err := repo.DB.First(&integrations, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(entities.ErrNotFoundInt)
		}
		return err
	}

	if err := repo.DB.Model(&integrations).Updates(integration).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) DeleteIntegration(id int64) error {
	var integration entities.AccountIntegration

	if err := repo.DB.First(&integration, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(entities.ErrNotFoundInt)
		}
		return err
	}

	if err := repo.DB.Delete(&integration).Error; err != nil {
		return err
	}
	return nil
}
