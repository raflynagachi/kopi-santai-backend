package model

type Menu struct {
	ID         uint `gorm:"primaryKey"`
	CategoryID uint
	Name       string
	Price      float64
	Image      string
	Category   *Category
}

func (_ *Menu) TableName() string {
	return "menus_tab"
}
