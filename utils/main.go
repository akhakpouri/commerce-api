package main

import (
	"commerce/internal/shared/database"
	"commerce/utils/internal/managers"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Commerce Utility Application!")

	cfg, err := managers.NewDbConfig("configs/config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	database.Migrate(cfg)
}
