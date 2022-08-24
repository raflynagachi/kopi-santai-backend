package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	FindByMinSpent(tx *gorm.DB, spent uint) ([]*model.Promotion, error)
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
