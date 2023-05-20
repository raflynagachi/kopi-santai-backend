package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

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
