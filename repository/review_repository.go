package repository

import (
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewRepository interface {
	Create(tx *gorm.DB, review *model.Review) (*model.Review, error)
	FindByMenuID(tx *gorm.DB, menuID uint) ([]*model.Review, error)
}

type reviewRepository struct {
}

func NewReview() ReviewRepository {
	return &reviewRepository{}
}

func (r *reviewRepository) Create(tx *gorm.DB, review *model.Review) (*model.Review, error) {
	result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&review)
	if int(result.RowsAffected) == 0 {
		return nil, new(apperror.ReviewCreatedError)
	}
	tx.Preload("User").First(&review)
	return review, result.Error
}

func (r *reviewRepository) FindByMenuID(tx *gorm.DB, menuID uint) ([]*model.Review, error) {
	var reviews []*model.Review
	result := tx.Preload("User").Where("menu_id = ?", menuID).Find(&reviews)
	if result.RowsAffected == 0 {
		return nil, new(apperror.ReviewNotFoundError)
	}
	return reviews, result.Error
}
