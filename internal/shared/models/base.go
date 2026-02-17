package models

import "time"

type Base struct {
	Id          uint      `gorm:"primaryKey"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	DeletedDate time.Time `gorm:"index"`
}
