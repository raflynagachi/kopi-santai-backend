package model

const (
	Price             = "price"
	ID                = "id"
	Desc              = "desc"
	Asc               = "asc"
	CategoryCoffee    = "coffee"
	CategoryNonCoffee = "non-coffee"
	CategoryBread     = "bread"
)

type QueryParamMenu struct {
	Search   string
	SortBy   string
	Sort     string
	Category string
}

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
