package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewRepository interface {
	Create(tx *gorm.DB, review *model.Review) (*model.Review, error)
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
	return review, result.Error
}
