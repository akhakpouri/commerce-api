package models

type Product struct {
	Base
	Name              string            `gorm:"type:text;size:150" sql:"type:text"`
	Price             float32           `gorm:"type:decimal(10,2)" sql:"type:decimal(10,2)"`
	Description       string            `gorm:"type:text;size:255" sql:"type:text"`
	Sku               string            `gorm:"type:text;size:100;uniqueIndex" sql:"type:text"`
	Stock             int               `gorm:"default:0"`
	IsActive          bool              `gorm:"default:true"`
	IsFeatured        bool              `gorm:"default:false"`
	ProductCategories []ProductCategory `gorm:"foreignKey:ProductId"`
}

func (Product) TableName() string {
	return "products"
}
