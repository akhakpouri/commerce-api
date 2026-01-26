package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	connection := "host=localhost user=postgres dbname=commerce port=5432 password=commerce@123"
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}
	database.AutoMigrate(&Product{})
	Db = database
}
