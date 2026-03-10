package models

type ProductCategory struct {
	Base
	ProductId  uint     `gorm:"not null;"`
	CategoryId uint     `gorm:"not null;"`
	Product    Product  `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	Category   Category `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
