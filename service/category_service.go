package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
)

type CategoryService interface {
	FindAll() ([]*dto.CategoryRes, error)
}

type categoryService struct {
	db           *gorm.DB
	categoryRepo repository.CategoryRepository
}

type CategoryConfig struct {
	DB           *gorm.DB
	CategoryRepo repository.CategoryRepository
}

func NewCategory(c *CategoryConfig) CategoryService {
	return &categoryService{
		db:           c.DB,
		categoryRepo: c.CategoryRepo,
	}
}

func (s *categoryService) FindAll() ([]*dto.CategoryRes, error) {
	var categoriesRes []*dto.CategoryRes
	tx := s.db.Begin()
	categories, err := s.categoryRepo.FindAll(tx)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	for _, cat := range categories {
		categoriesRes = append(categoriesRes, new(dto.CategoryRes).From(cat))
	}
	return categoriesRes, nil
}
