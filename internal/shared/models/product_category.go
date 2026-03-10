package models

type ProductCategory struct {
	Base
	ProductId  uint     `gorm:"not null;foreignKey:ProductId;constraint:OnDelete:SET NULL"`
	CategoryId uint     `gorm:"not null;foreignKey:CategoryId;constraint:OnDelete:SET NULL"`
	Product    Product  `gorm:"foreignKey:ProductId"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
