package service

import (
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"gorm.io/gorm"
)

type ReviewService interface {
	Create(req *dto.ReviewPostReq, userID uint) (*dto.ReviewRes, error)
	FindByMenuID(menuID uint) ([]*dto.ReviewRes, error)
}

type reviewService struct {
	db         *gorm.DB
	reviewRepo repository.ReviewRepository
}

type ReviewConfig struct {
	DB         *gorm.DB
	ReviewRepo repository.ReviewRepository
}

func NewReview(c *ReviewConfig) ReviewService {
	return &reviewService{
		db:         c.DB,
		reviewRepo: c.ReviewRepo,
	}
}

func (s *reviewService) Create(req *dto.ReviewPostReq, userID uint) (*dto.ReviewRes, error) {
	r := &model.Review{
		UserID:      userID,
		MenuID:      req.MenuID,
		Description: req.Description,
		Rating:      req.Rating,
	}

	tx := s.db.Begin()
	review, err := s.reviewRepo.Create(tx, r)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	reviewRes := new(dto.ReviewRes).FromReview(review)
	return reviewRes, nil
}

func (s *reviewService) FindByMenuID(menuID uint) ([]*dto.ReviewRes, error) {
	tx := s.db.Begin()
	reviews, err := s.reviewRepo.FindByMenuID(tx, menuID)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	var reviewsRes []*dto.ReviewRes
	for _, review := range reviews {
		reviewsRes = append(reviewsRes, new(dto.ReviewRes).FromReview(review))
	}
	return reviewsRes, nil
}
