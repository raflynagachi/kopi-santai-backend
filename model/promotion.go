package model

import "gorm.io/gorm"

type Promotion struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	CouponID    uint
	Name        string
	Description string
	Image       []byte
	MinSpent    uint
}

func (_ *Promotion) TableName() string {
	return "promotions_tab"
}
