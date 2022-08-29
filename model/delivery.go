package model

import (
	"gorm.io/gorm"
	"time"
)

const StatusDefault = "ON PROCESS"

type Delivery struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	DeliveryDate time.Time
	Status       string
}

func (_ *Delivery) TableName() string {
	return "deliveries_tab"
}
