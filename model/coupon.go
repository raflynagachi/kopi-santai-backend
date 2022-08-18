package model

import (
	"gorm.io/gorm"
	"time"
)

type Coupon struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Amount      float64
	IsAvailable bool
	MinSpent    float64
	ExpiredDate time.Time
}

func (c *Coupon) TableName() string {
	return "coupons_tab"
}
