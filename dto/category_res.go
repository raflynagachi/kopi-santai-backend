package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type CategoryRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (_ *CategoryRes) From(c *model.Category) *CategoryRes {
	return &CategoryRes{
		ID:   c.ID,
		Name: c.Name,
	}
}
