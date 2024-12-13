package unisender_integration

import (
	"http-server/internal/entities"
	rep "http-server/internal/repository"

	"gorm.io/gorm"
)

var _ rep.UnisenderIntegrationRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	return &repository{
		DB: db,
	}, nil
}

func (repo *repository) GetUnisenderKey() ([]entities.UnisenderIntegration, error) {
	var unisenderKey []entities.UnisenderIntegration

	err := repo.DB.Find(&unisenderKey).Error
	if err != nil {
		return nil, err
	}
	return unisenderKey, nil
}

func (repo *repository) SaveUnisenderKey(unisenderKey *entities.UnisenderIntegration) error {
	if err := repo.DB.Create(unisenderKey).Error; err != nil {
		return err
	}
	return nil
}
