package model

import (
	"gorm.io/gorm"
	"time"
)

const LimitDefaultValue = "10"
const PageDefaultValue = "1"

type QueryParamOrder struct {
	Date  string
	Limit int
	Page  int
}

type Order struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	CouponID        *uint
	DeliveryID      uint
	PaymentOptionID uint
	OrderedDate     time.Time
	TotalPrice      float64
	Coupon          *Coupon
	Delivery        *Delivery
	PaymentOption   *PaymentOption
	OrderItems      []*OrderItem
}

func (_ *Order) TableName() string {
	return "orders_tab"
}
