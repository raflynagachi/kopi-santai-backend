package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type MenuRes struct {
	ID           uint    `json:"id"`
	CategoryID   uint    `json:"categoryID"`
	CategoryName string  `json:"categoryName"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
}

func (_ *MenuRes) FromMenu(m *model.Menu) *MenuRes {
	return &MenuRes{
		ID:           m.ID,
		CategoryID:   m.CategoryID,
		CategoryName: m.Category.Name,
		Name:         m.Name,
		Price:        m.Price,
		Image:        m.Image,
	}
}
