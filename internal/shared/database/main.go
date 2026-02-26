package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(cfg DbConfig) {
	connection := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%d password=%s sslmode=%s search_path=%s",
		cfg.Host, cfg.User, cfg.DbName, cfg.Port, cfg.Password, cfg.SSLMode, cfg.Schema)
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
