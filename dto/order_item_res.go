package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

type OrderItemRes struct {
	ID          uint     `json:"id"`
	Menu        *MenuRes `json:"menu" `
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
}

func (_ *OrderItemRes) From(oi *model.OrderItem, m *MenuRes) *OrderItemRes {
	return &OrderItemRes{
		ID:          oi.ID,
		Menu:        m,
		Quantity:    oi.Quantity,
		Description: oi.Description,
	}
}
