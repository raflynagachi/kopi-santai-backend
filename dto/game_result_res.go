package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type UserCouponRes struct {
	ID        uint       `json:"id"`
	UserRes   *UserRes   `json:"userRes"`
	CouponRes *CouponRes `json:"couponRes"`
}

func (_ *UserCouponRes) From(uc *model.UserCoupon) *UserCouponRes {
	return &UserCouponRes{
		ID:        uc.ID,
		UserRes:   new(UserRes).FromUser(uc.User),
		CouponRes: new(CouponRes).FromCoupon(uc.Coupon),
	}
}
