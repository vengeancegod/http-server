package sql

import (
	"http-server/internal/entities"

	"gorm.io/gorm"
)

func CreateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.Account{},
		&entities.AccountIntegration{},
		&entities.Contacts{},
		&entities.UnisenderIntegration{},
	)
}
