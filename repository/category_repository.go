package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(tx *gorm.DB) ([]*model.Category, error)
}

type categoryRepository struct{}

func NewCategory() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAll(tx *gorm.DB) ([]*model.Category, error) {
	var cat []*model.Category
	err := tx.Find(&cat).Error
	return cat, err
}
