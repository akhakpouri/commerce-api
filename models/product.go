package models

import "time"

type Product struct {
	Id          int
	Name        string
	Price       float32
	Description string
	CreatedDate time.Time
}
