package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	FindByMinSpent(tx *gorm.DB, spent uint) ([]*model.Promotion, error)
	FindAll(tx *gorm.DB) ([]*model.Promotion, error)
	FindAllUnscoped(tx *gorm.DB) ([]*model.Promotion, error)
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
