package models

type ProductCategory struct {
	Base
	ProductId  uint     `gorm:"not null;foreignKey:ProductId"`
	CategoryId uint     `gorm:"not null;foreignKey:CategoryId"`
	Product    Product  `gorm:"foreignKey:ProductId"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
