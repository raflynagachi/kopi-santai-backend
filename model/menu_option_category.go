package model

type MenuOptionsCategories struct {
	ID           uint `gorm:"primaryKey"`
	CategoryID   uint
	MenuOptionID uint
	MenuOption   *MenuOption
	Name         string
}

func (_ *MenuOptionsCategories) TableName() string {
	return "menu_options_categories_tab"
}
