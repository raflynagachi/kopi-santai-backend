package model

type MenuOption struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (mo *MenuOption) TableName() string {
	return "menu_options_tab"
}
