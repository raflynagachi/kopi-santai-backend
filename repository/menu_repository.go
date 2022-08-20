package repository

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	FindAll(tx *gorm.DB) ([]*model.Menu, error)
	FindByID(tx *gorm.DB, id uint) (*model.Menu, error)
	FindMenuOptions(tx *gorm.DB, categoryID uint) ([]*model.MenuOptionsCategories, error)
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

func (r *menuRepository) FindByID(tx *gorm.DB, id uint) (*model.Menu, error) {
	var menu *model.Menu
	err := tx.Preload("Category").First(&menu, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, new(apperror.MenuNotFoundError)
	}
	return menu, err
}

func (r *menuRepository) FindMenuOptions(tx *gorm.DB, categoryID uint) ([]*model.MenuOptionsCategories, error) {
	var menuOptionCategories []*model.MenuOptionsCategories
	err := tx.Preload("MenuOption").Find(&menuOptionCategories, "category_id = ?", categoryID).Error
	return menuOptionCategories, err
}
