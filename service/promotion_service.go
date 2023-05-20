package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"gorm.io/gorm"
)

type PromotionService interface {
	FindAll() ([]*dto.PromotionRes, error)
	FindAllUnscoped() ([]*dto.PromotionRes, error)
	CreatePromotion(req *dto.PromotionPostReq) (*dto.PromotionRes, error)
	DeletePromotionByID(id uint) (gin.H, error)
}

type promotionService struct {
	db        *gorm.DB
	promoRepo repository.PromotionRepository
}

type PromoConfig struct {
	DB        *gorm.DB
	PromoRepo repository.PromotionRepository
}

func NewPromo(c *PromoConfig) PromotionService {
	return &promotionService{
		db:        c.DB,
		promoRepo: c.PromoRepo,
	}
}

func (s *promotionService) FindAll() ([]*dto.PromotionRes, error) {
	tx := s.db.Begin()
	promos, err := s.promoRepo.FindAll(tx)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var promosRes []*dto.PromotionRes
	for _, promo := range promos {
		if promo.Coupon != nil {
			promosRes = append(promosRes, new(dto.PromotionRes).FromPromotion(promo))
		}
	}

	return promosRes, nil
}

func (s *promotionService) FindAllUnscoped() ([]*dto.PromotionRes, error) {
	tx := s.db.Begin()
	promos, err := s.promoRepo.FindAllUnscoped(tx)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var promosRes []*dto.PromotionRes
	for _, promo := range promos {
		if promo.Coupon != nil {
			promosRes = append(promosRes, new(dto.PromotionRes).FromPromotion(promo))
		}
	}

	return promosRes, nil
}

func (s *promotionService) CreatePromotion(req *dto.PromotionPostReq) (*dto.PromotionRes, error) {
	p := &model.Promotion{
		CouponID:    req.CouponID,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		MinSpent:    req.MinSpent,
	}

	tx := s.db.Begin()
	promo, err := s.promoRepo.Create(tx, p)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	promotionRes := new(dto.PromotionRes).FromPromotion(promo)
	return promotionRes, nil
}

func (s *promotionService) DeletePromotionByID(id uint) (gin.H, error) {
	tx := s.db.Begin()
	isDeleted, err := s.promoRepo.Delete(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return gin.H{"isDeleted": false}, apperror.BadRequestError(err.Error())
	}

	return gin.H{"isDeleted": isDeleted}, nil
}
