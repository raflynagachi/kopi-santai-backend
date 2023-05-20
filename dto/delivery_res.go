package dto

import (
	"time"

	"github.com/raflynagachi/kopi-santai-backend/model"
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
