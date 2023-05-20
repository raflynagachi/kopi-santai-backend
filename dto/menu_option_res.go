package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type MenuOptionRes struct {
	TypeMenuOption string `json:"typeMenuOption"`
	Name           string `json:"name"`
}

func (_ *MenuOptionRes) FromMenuOptionsCategories(m *model.MenuOptionsCategories) *MenuOptionRes {
	return &MenuOptionRes{
		TypeMenuOption: m.MenuOption.Name,
		Name:           m.Name,
	}
}
