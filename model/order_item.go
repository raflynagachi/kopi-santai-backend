package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	MenuID      uint
	OrderID     *uint
	Quantity    int
	Description string
	Menu        *Menu
}

func (_ *OrderItem) TableName() string {
	return "order_items_tab"
}
