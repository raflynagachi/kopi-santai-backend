package repository

import (
	"errors"
	"fmt"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	FindAll(tx *gorm.DB, q *model.QueryParamMenu) ([]*model.Menu, error)
	FindAllUnscoped(tx *gorm.DB, q *model.QueryParamMenu) ([]*model.Menu, error)
	FindByID(tx *gorm.DB, id uint) (*model.Menu, error)
	FindMenuOptions(tx *gorm.DB, categoryID uint) ([]*model.MenuOptionsCategories, error)
	Create(tx *gorm.DB, menu *model.Menu) (*model.Menu, error)
	Update(tx *gorm.DB, id uint, d *model.Menu) (*model.Menu, error)
	DeleteByID(tx *gorm.DB, id uint) (bool, error)
}

type menuRepository struct{}

func NewMenu() MenuRepository {
	return &menuRepository{}
}

func (r *menuRepository) FindAllUnscoped(tx *gorm.DB, q *model.QueryParamMenu) ([]*model.Menu, error) {
	var categories []*model.Category
	var menus []*model.Menu
	var err error
	orderStatement := fmt.Sprintf("%s %s", q.SortBy, q.Sort)

	_ = tx.Distinct().Select("id").Where("LOWER(name) LIKE ?", q.Category).Find(&categories).Error
	var ids []uint
	for _, category := range categories {
		ids = append(ids, category.ID)
	}

	err = tx.Unscoped().Preload("Category").Preload("Reviews").Where("category_id IN (?) AND LOWER(name) LIKE ?", ids, q.Search).Order(orderStatement).Find(&menus).Error
	return menus, err
}

func (r *menuRepository) FindAll(tx *gorm.DB, q *model.QueryParamMenu) ([]*model.Menu, error) {
	var categories []*model.Category
	var menus []*model.Menu
	var err error
	orderStatement := fmt.Sprintf("%s %s", q.SortBy, q.Sort)

	_ = tx.Distinct().Select("id").Where("LOWER(name) LIKE ?", q.Category).Find(&categories).Error
	var ids []uint
	for _, category := range categories {
		ids = append(ids, category.ID)
	}

	err = tx.Preload("Category").Preload("Reviews").Where("category_id IN (?) AND LOWER(name) LIKE ?", ids, q.Search).Order(orderStatement).Find(&menus).Error
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

func (r *menuRepository) Create(tx *gorm.DB, menu *model.Menu) (*model.Menu, error) {
	err := tx.Preload("Category").Create(&menu).First(&menu).Error
	return menu, err
}

func (r *menuRepository) Update(tx *gorm.DB, id uint, d *model.Menu) (*model.Menu, error) {
	var updatedMenu *model.Menu
	err := tx.First(&updatedMenu, id).Updates(&d).Error
	if err != nil {
		return nil, err
	}
	_ = tx.Preload("Category").First(&updatedMenu, id)
	return updatedMenu, nil
}

func (r *menuRepository) DeleteByID(tx *gorm.DB, id uint) (bool, error) {
	var deletedMenu *model.Menu
	err := tx.Delete(&deletedMenu, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
