package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ID          uint
	UserID      uint `gorm:"primaryKey"`
	MenuID      uint `gorm:"primaryKey"`
	Description string
	Rating      float64
	User        *User
}

func (_ *Review) TableName() string {
	return "reviews_tab"
}
