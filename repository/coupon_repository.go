package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type CouponRepository interface {
	FindByID(tx *gorm.DB, id uint) (*model.Coupon, error)
	Create(tx *gorm.DB, c *model.Coupon) (*model.Coupon, error)
	AddCouponToUser(tx *gorm.DB, uc *model.UserCoupon) (*model.UserCoupon, error)
	DeleteByID(tx *gorm.DB, id uint) (bool, error)
}

type couponRepository struct{}

func NewCoupon() CouponRepository {
	return &couponRepository{}
}

func (r *couponRepository) FindByID(tx *gorm.DB, id uint) (*model.Coupon, error) {
	var coupon *model.Coupon
	err := tx.Where("is_available = true").First(&coupon, id).Error
	return coupon, err
}

func (r *couponRepository) Create(tx *gorm.DB, c *model.Coupon) (*model.Coupon, error) {
	err := tx.Create(&c).Error
	return c, err
}

func (r *couponRepository) AddCouponToUser(tx *gorm.DB, uc *model.UserCoupon) (*model.UserCoupon, error) {
	err := tx.Preload("User").Preload("Coupon").Create(&uc).First(&uc).Error
	return uc, err
}

func (r *couponRepository) DeleteByID(tx *gorm.DB, id uint) (bool, error) {
	var deletedCoupon *model.Coupon
	err := tx.Delete(&deletedCoupon, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
