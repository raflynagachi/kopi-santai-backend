package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	FindAll(tx *gorm.DB) ([]*model.Menu, error)
}

type menuRepository struct{}

func NewMenu() MenuRepository {
	return &menuRepository{}
}

func (r *menuRepository) FindAll(tx *gorm.DB) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Preload("Category").Find(&menus).Error
	return menus, err
}
