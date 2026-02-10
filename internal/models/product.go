package models

type Product struct {
	Base
	Name        string
	Price       float32
	Description string
	CategoryId  uint
}

func (Product) TableName() string {
	return "products"
}
