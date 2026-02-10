package models

type Category struct {
	Base
	Name        string
	Description string
	Products    []Product `gorm:"foreignKey:CategoryId"`
}

func (Category) TableName() string {
	return "categories"
}
