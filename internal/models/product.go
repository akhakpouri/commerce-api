package models

type Product struct {
	Base
	Name        string
	Price       float32
	Description string
	CategoryId  uint
}
