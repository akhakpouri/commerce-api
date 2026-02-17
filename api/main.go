package main

import (
	"commerce/internal/shared/database"
	"fmt"
)

func main() {
	fmt.Println("hello, world!")
	database.Migrate()
}
