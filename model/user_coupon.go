package model

type UserCoupon struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	CouponID uint
	User     *User
	Coupon   *Coupon
}

func (uc *UserCoupon) TableName() string {
	return "users_coupons_tab"
}
