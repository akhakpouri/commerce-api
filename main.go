package main

import (
	"commerce/api/internal/database"
	"fmt"
)

func main() {
	fmt.Println("hello, world!")
	database.Migrate()
}
