package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type CouponRes struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Amount    float64        `json:"amount"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (_ *CouponRes) FromCoupon(c *model.Coupon) *CouponRes {
	return &CouponRes{
		ID:        c.ID,
		Name:      c.Name,
		Amount:    c.Amount,
		DeletedAt: c.DeletedAt,
	}
}
