package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type OrderItemRes struct {
	Menu        *MenuRes `json:"menu" `
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
}

func (_ *OrderItemRes) From(oi *model.OrderItem, m *MenuRes) *OrderItemRes {
	return &OrderItemRes{
		Menu:        m,
		Quantity:    oi.Quantity,
		Description: oi.Description,
	}
}
