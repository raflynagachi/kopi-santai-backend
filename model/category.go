package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (c *Category) TableName() string {
	return "categories_tab"
}
