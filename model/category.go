package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID                    uint `gorm:"primaryKey"`
	Name                  string
	MenuOptionsCategories *MenuOptionsCategories `gorm:"many2many:menu_options_categories_tab"`
}

func (c *Category) TableName() string {
	return "categories_tab"
}
