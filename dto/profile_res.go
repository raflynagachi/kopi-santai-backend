package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
)

type ProfileRes struct {
	FullName       string       `json:"fullName"`
	Phone          string       `json:"phone"`
	Email          string       `json:"email"`
	Username       string       `json:"username"`
	Address        string       `json:"address"`
	ProfilePicture string       `json:"profilePicture"`
	Coupons        []*CouponRes `json:"coupon"`
}

func couponsToCouponsRes(coupons []*model.Coupon) []*CouponRes {
	var couponsRes []*CouponRes
	for _, coupon := range coupons {
		couponsRes = append(couponsRes, new(CouponRes).FromCoupon(coupon))
	}
	return couponsRes
}

func (_ *ProfileRes) FromUser(u *model.User) *ProfileRes {
	coupons := couponsToCouponsRes(u.Coupons)
	return &ProfileRes{
		FullName:       u.FullName,
		Phone:          u.Phone,
		Email:          u.Email,
		Username:       u.Username,
		Address:        u.Address,
		ProfilePicture: u.ProfilePicture,
		Coupons:        coupons,
	}
}
