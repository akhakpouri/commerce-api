package database

import (
	"commerce/internal/shared/models"

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
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
	)
}
