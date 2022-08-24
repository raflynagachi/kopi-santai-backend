package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
)

type PromotionService interface {
	FindAll() ([]*dto.PromotionRes, error)
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
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var promosRes []*dto.PromotionRes
	for _, promo := range promos {
		promosRes = append(promosRes, new(dto.PromotionRes).FromPromotion(promo))
	}

	return promosRes, nil
}
