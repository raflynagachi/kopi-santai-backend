package model

import (
	"gorm.io/gorm"
	"time"
)

const StatusDefault = "NO DELIVERY YET"
const StatusDelivered = "DELIVERED"

type Delivery struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	DeliveryDate time.Time
	Status       string
}

func (_ *Delivery) TableName() string {
	return "deliveries_tab"
}
