package dto

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
)

type MenuRes struct {
	ID           uint    `json:"id"`
	CategoryID   uint    `json:"categoryID"`
	CategoryName string  `json:"categoryName"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	Rating       float64 `json:"rating"`
}

func (_ *MenuRes) FromMenu(m *model.Menu) *MenuRes {
	var rating float64
	if len(m.Reviews) != 0 {
		for _, review := range m.Reviews {
			rating += review.Rating
		}
		rating /= float64(len(m.Reviews))
	}
	return &MenuRes{
		ID:           m.ID,
		CategoryID:   m.CategoryID,
		CategoryName: m.Category.Name,
		Name:         m.Name,
		Price:        m.Price,
		Image:        m.Image,
		Rating:       helper.ToFixed(rating, 2),
	}
}
