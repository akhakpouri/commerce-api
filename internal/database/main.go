package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connection := "host=localhost user=commerce dbname=commerce port=5432 password=commerce@123 sslmode=disable search_path=commerce"
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}
	err = Migrate(database)
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database, %v", err))
	}
}
