package unisender_integration

import (
	"http-server/internal/entities"
	"http-server/internal/infrastructure/database/sql"
	rep "http-server/internal/repository"

	"gorm.io/gorm"
)

var _ rep.UnisenderIntegrationRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository() (*repository, error) {
	db, err := sql.InitDB()
	if err != nil {
		return nil, err
	}

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
