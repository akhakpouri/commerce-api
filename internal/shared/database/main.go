package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	connection := "host=localhost user=commerce dbname=commerce port=5432 password=commerce@123 sslmode=disable search_path=commerce"
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to the database")
	}

	log.Println("Connected to the database successfully.")
	log.Println("Running migration.")

	err = setup(database)
	if err != nil {
		log.Fatal("Migration failed: ", err)
		panic(fmt.Sprintf("Failed to migrate database, %v", err))
	}
	log.Println("Migration completed successfully.")
}
