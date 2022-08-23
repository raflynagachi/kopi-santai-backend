package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
)

type CouponRes struct {
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"`
	IsAvailable bool    `json:"isAvailable"`
}

func (_ *CouponRes) FromCoupon(c *model.Coupon) *CouponRes {
	return &CouponRes{
		Name:        c.Name,
		Amount:      c.Amount,
		IsAvailable: c.IsAvailable,
	}
}
