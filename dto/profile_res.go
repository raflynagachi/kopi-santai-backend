package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
)

type ProfileRes struct {
	User    *UserRes     `json:"user"`
	Coupons []*CouponRes `json:"coupons"`
}

func couponsToCouponsRes(coupons []*model.Coupon) []*CouponRes {
	var couponsRes []*CouponRes
	for _, coupon := range coupons {
		couponsRes = append(couponsRes, new(CouponRes).FromCoupon(coupon))
	}
	return couponsRes
}

func (_ *ProfileRes) FromUser(u *model.User) *ProfileRes {
	user := new(UserRes).FromUser(u)
	coupons := couponsToCouponsRes(u.Coupons)
	return &ProfileRes{
		User:    user,
		Coupons: coupons,
	}
}
