package repository

import (
	"github.com/raflynagachi/kopi-santai-backend/model"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	FindByMinSpent(tx *gorm.DB, spent uint) ([]*model.Promotion, error)
	FindAll(tx *gorm.DB) ([]*model.Promotion, error)
	FindAllUnscoped(tx *gorm.DB) ([]*model.Promotion, error)
	Create(tx *gorm.DB, p *model.Promotion) (*model.Promotion, error)
	Delete(tx *gorm.DB, id uint) (bool, error)
}

type promotionRepository struct {
}

func NewPromo() PromotionRepository {
	return &promotionRepository{}
}

func (r *promotionRepository) FindByMinSpent(tx *gorm.DB, spent uint) ([]*model.Promotion, error) {
	var p []*model.Promotion
	result := tx.Where("min_spent <= ?", spent).Find(&p)
	return p, result.Error
}

func (r *promotionRepository) FindAll(tx *gorm.DB) ([]*model.Promotion, error) {
	var p []*model.Promotion
	result := tx.Preload("Coupon").Find(&p)
	return p, result.Error
}

func (r *promotionRepository) FindAllUnscoped(tx *gorm.DB) ([]*model.Promotion, error) {
	var p []*model.Promotion
	result := tx.Preload("Coupon", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Find(&p)
	return p, result.Error
}

func (r *promotionRepository) Create(tx *gorm.DB, p *model.Promotion) (*model.Promotion, error) {
	err := tx.Preload("Coupon").Create(&p).First(&p).Error
	return p, err
}

func (r *promotionRepository) Delete(tx *gorm.DB, id uint) (bool, error) {
	var deletedPromotion *model.Promotion
	err := tx.Delete(&deletedPromotion, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
