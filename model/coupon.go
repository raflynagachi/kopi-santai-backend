package model

import (
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	Name   string
	Amount float64
}

func (c *Coupon) TableName() string {
	return "coupons_tab"
}
