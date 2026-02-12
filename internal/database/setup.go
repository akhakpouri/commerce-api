package database

import (
	"commerce/api/internal/models"

	"gorm.io/gorm"
)

func setup(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Address{},
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.ProductCategory{},
		&models.Review{},
	)
}
