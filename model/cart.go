package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	OrderItemID uint
}

func (_ *Cart) TableName() string {
	return "carts_tab"
}
