package models

type Category struct {
	Base
	Name              string            `gorm:"type:text;size:255" sql:"type:text"`
	Description       string            `gorm:"type:text;size:255" sql:"type:text"`
	Slug              string            `gorm:"type:text;size:100" sql:"type:text"`
	ParentId          *uint             `gorm:"index"`
	IsActive          bool              `gorm:"default:true"`
	Children          []Category        `gorm:"foreignKey:ParentId;constraint:OnDelete:CASCADE"`
	ProductCategories []ProductCategory `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE"`
}

func (Category) TableName() string {
	return "categories"
}
