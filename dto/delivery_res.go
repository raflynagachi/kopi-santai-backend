package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"time"
)

type DeliveryRes struct {
	ID           uint      `json:"id"`
	DeliveryDate time.Time `json:"deliveryDate"`
	Status       string    `json:"status"`
}

func (_ *DeliveryRes) FromDelivery(d *model.Delivery) *DeliveryRes {
	return &DeliveryRes{
		ID:           d.ID,
		DeliveryDate: d.DeliveryDate,
		Status:       d.Status,
	}
}
