package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
)

type MenuService interface {
	FindAll() ([]*dto.MenuRes, error)
}

type menuService struct {
	db             *gorm.DB
	menuRepository repository.MenuRepository
}

type MenuConfig struct {
	DB             *gorm.DB
	MenuRepository repository.MenuRepository
}

func NewMenu(c *MenuConfig) MenuService {
	return &menuService{
		db:             c.DB,
		menuRepository: c.MenuRepository,
	}
}

func menusToMenusRes(menus []*model.Menu) []*dto.MenuRes {
	var menusRes []*dto.MenuRes
	for _, menu := range menus {
		menusRes = append(menusRes, new(dto.MenuRes).FromMenu(menu))
	}
	return menusRes
}

func (s *menuService) FindAll() ([]*dto.MenuRes, error) {
	tx := s.db.Begin()
	menus, err := s.menuRepository.FindAll(tx)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	menusRes := menusToMenusRes(menus)
	return menusRes, err
}
