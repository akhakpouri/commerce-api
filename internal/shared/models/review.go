package models

type Review struct {
	Base
	ProductId uint    `gorm:"not null;"`
	UserId    uint    `gorm:"not null;"`
	Rating    int     `gorm:"not null"`
	Title     string  `gorm:"type:text;size:100" sql:"type:text"`
	Comment   string  `gorm:"type:text;size:500" sql:"type:text"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

func (Review) TableName() string {
	return "reviews"
}
