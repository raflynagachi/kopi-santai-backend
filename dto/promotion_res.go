package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

type PromotionRes struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       []byte     `json:"image"`
	MinSpent    uint       `json:"minSpent"`
	Coupon      *CouponRes `json:"coupon"`
}

func (_ *PromotionRes) FromPromotion(p *model.Promotion) *PromotionRes {
	coupon := new(CouponRes).FromCoupon(p.Coupon)
	return &PromotionRes{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		MinSpent:    p.MinSpent,
		Coupon:      coupon,
	}
}
