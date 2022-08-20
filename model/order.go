package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	CouponID        uint
	DeliveryID      uint
	PaymentOptionID uint
	CartID          uint
	OrderedDate     time.Time
	TotalPrice      float64
	IsCompleted     bool
	Coupon          *Coupon
	Delivery        *Delivery
	Payment         *PaymentOption
	OrderItems      []*OrderItem `gorm:"many2many:carts_tab;"`
}

func (_ *Order) TableName() string {
	return "orders_tab"
}
