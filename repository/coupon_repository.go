package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type CouponRepository interface {
	FindByID(tx *gorm.DB, id uint) (*model.Coupon, error)
}

type couponRepository struct{}

func NewCoupon() CouponRepository {
	return &couponRepository{}
}

func (r *couponRepository) FindByID(tx *gorm.DB, id uint) (*model.Coupon, error) {
	var coupon *model.Coupon
	err := tx.First(&coupon, id).Error
	return coupon, err
}
