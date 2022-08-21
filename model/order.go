package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	CouponID        *uint
	DeliveryID      uint
	PaymentOptionID uint
	OrderedDate     time.Time
	TotalPrice      float64
	IsActive        bool
	Coupon          *Coupon
	Delivery        *Delivery
	PaymentOption   *PaymentOption
	OrderItems      []*OrderItem
}

func (_ *Order) TableName() string {
	return "orders_tab"
}
