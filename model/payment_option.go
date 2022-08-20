package model

import "gorm.io/gorm"

type PaymentOption struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (_ *PaymentOption) TableName() string {
	return "payment_options_tab"
}
