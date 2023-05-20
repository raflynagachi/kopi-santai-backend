package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

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
