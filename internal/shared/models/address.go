package models

type Address struct {
	Base
	UserId     uint   `gorm:"not null;foreignKey:UserId"`
	Street     string `gorm:"type:text;size:255" sql:"type:text"`
	City       string `gorm:"type:text;size:100" sql:"type:text"`
	State      string `gorm:"type:text;size:50" sql:"type:text"`
	PostalCode string `gorm:"type:text;size:15" sql:"type:text"`
	Country    string `gorm:"type:text;size:75" sql:"type:text"`
	IsDefault  bool   `gorm:"default:false"`
	User       User   `gorm:"foreignKey:UserId"`
}

func (Address) TableName() string {
	return "addresses"
}
