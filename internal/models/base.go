package models

import "time"

type Base struct {
	Id          uint `gorm:"primaryKey"`
	CreatedDate time.Time
	UpdatedDate time.Time
}
