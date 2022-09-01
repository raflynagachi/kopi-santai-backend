package service

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type MenuService interface {
	FindAll(q *model.QueryParamMenu) ([]*dto.MenuRes, error)
	FindAllUnscoped(q *model.QueryParamMenu) ([]*dto.MenuRes, error)
	GetMenuDetail(id uint) (*dto.MenuDetailRes, error)
	Create(req *dto.MenuPostReq) (*dto.MenuRes, error)
	Update(id uint, req *dto.MenuUpdateReq) (*dto.MenuRes, error)
	DeleteByID(id uint) (gin.H, error)
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

func menuOptionsCategoriesToMenuOptionsRes(menusOptCats []*model.MenuOptionsCategories) []*dto.MenuOptionRes {
	var menuOptRes []*dto.MenuOptionRes
	for _, menusOptCat := range menusOptCats {
		menuOptRes = append(menuOptRes, new(dto.MenuOptionRes).FromMenuOptionsCategories(menusOptCat))
	}
	return menuOptRes
}

func validateQueryParamMenu(q *model.QueryParamMenu) {
	if q.Sort != model.Desc && q.Sort != model.Asc {
		q.Sort = model.Desc
	}

	if q.SortBy != model.Price {
		q.SortBy = model.ID
	}

	q.Category = strings.ToLower(q.Category)
	if q.Category != model.CategoryCoffee && q.Category != model.CategoryNonCoffee && q.Category != model.CategoryBread {
		q.Category = "%%"
	}

	if q.Search != "" {
		q.Search = fmt.Sprintf("%%%s%%", strings.ToLower(q.Search))
	} else {
		q.Search = "%%"
	}
}

func (s *menuService) FindAll(q *model.QueryParamMenu) ([]*dto.MenuRes, error) {
	validateQueryParamMenu(q)

	tx := s.db.Begin()
	menus, err := s.menuRepository.FindAll(tx, q)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	menusRes := menusToMenusRes(menus)
	return menusRes, err
}

func (s *menuService) FindAllUnscoped(q *model.QueryParamMenu) ([]*dto.MenuRes, error) {
	validateQueryParamMenu(q)

	tx := s.db.Begin()
	menus, err := s.menuRepository.FindAllUnscoped(tx, q)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	menusRes := menusToMenusRes(menus)
	return menusRes, err
}

func (s *menuService) GetMenuDetail(id uint) (*dto.MenuDetailRes, error) {
	tx := s.db.Begin()
	menu, err := s.menuRepository.FindByID(tx, id)
	if err != nil {
		tx.Rollback()
		return nil, apperror.NotFoundError(err.Error())
	}
	menuOptions, err := s.menuRepository.FindMenuOptions(tx, menu.CategoryID)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	menuRes := new(dto.MenuRes).FromMenu(menu)
	menuOptionsRes := menuOptionsCategoriesToMenuOptionsRes(menuOptions)
	menuDetailRes := new(dto.MenuDetailRes).From(menuRes, menuOptionsRes)
	return menuDetailRes, nil
}

func (s *menuService) Create(req *dto.MenuPostReq) (*dto.MenuRes, error) {
	d := &model.Menu{
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Price:      req.Price,
		Image:      req.Image,
	}

	tx := s.db.Begin()
	menu, err := s.menuRepository.Create(tx, d)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	menuRes := new(dto.MenuRes).FromMenu(menu)
	return menuRes, nil
}

func (s *menuService) Update(id uint, req *dto.MenuUpdateReq) (*dto.MenuRes, error) {
	d := &model.Menu{
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Price:      req.Price,
		Image:      req.Image,
	}

	tx := s.db.Begin()
	menu, err := s.menuRepository.Update(tx, id, d)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	menuRes := new(dto.MenuRes).FromMenu(menu)
	return menuRes, nil
}

func (s *menuService) DeleteByID(id uint) (gin.H, error) {
	tx := s.db.Begin()
	isDeleted, err := s.menuRepository.DeleteByID(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return gin.H{"isDeleted": false}, apperror.BadRequestError(err.Error())
	}

	return gin.H{"isDeleted": isDeleted}, nil
}
