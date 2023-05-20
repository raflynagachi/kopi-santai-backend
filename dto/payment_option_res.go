package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

type PaymentOptionRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (_ *PaymentOptionRes) FromPaymentOption(po *model.PaymentOption) *PaymentOptionRes {
	return &PaymentOptionRes{
		ID:   po.ID,
		Name: po.Name,
	}
}
