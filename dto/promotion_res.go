package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type PromotionRes struct {
	Name        string
	Description string
	Image       []byte
	MinSpent    uint
	Coupon      *CouponRes
}

func (_ *PromotionRes) FromPromotion(p *model.Promotion) *PromotionRes {
	coupon := new(CouponRes).FromCoupon(p.Coupon)
	return &PromotionRes{
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		MinSpent:    p.MinSpent,
		Coupon:      coupon,
	}
}
