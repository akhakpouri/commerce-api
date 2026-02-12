package models

type Review struct {
	Base
	ProductId uint    `gorm:"not null;foreignKey:ProductId"`
	UserId    uint    `gorm:"not null;foreignKey:UserId"`
	Rating    int     `gorm:"not null"`
	Title     string  `gorm:"type:text;size:100" sql:"type:text"`
	Comment   string  `gorm:"type:text;size:500" sql:"type:text"`
	Product   Product `gorm:"foreignKey:ProductId"`
	User      User    `gorm:"foreignKey:UserId"`
}

func (Review) TableName() string {
	return "reviews"
}
