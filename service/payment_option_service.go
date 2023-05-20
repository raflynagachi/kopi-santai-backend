package service

import (
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"gorm.io/gorm"
)

type PaymentOptionService interface {
	FindAll() ([]*dto.PaymentOptionRes, error)
}

type paymentOptionService struct {
	db             *gorm.DB
	paymentOptRepo repository.PaymentOptionRepository
}

type PaymentOptConfig struct {
	DB             *gorm.DB
	PaymentOptRepo repository.PaymentOptionRepository
}

func NewPaymentOpt(c *PaymentOptConfig) PaymentOptionService {
	return &paymentOptionService{
		db:             c.DB,
		paymentOptRepo: c.PaymentOptRepo,
	}
}

func (s *paymentOptionService) FindAll() ([]*dto.PaymentOptionRes, error) {
	var paymentOptRes []*dto.PaymentOptionRes
	tx := s.db.Begin()
	paymentOpts, err := s.paymentOptRepo.FindAll(tx)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	for _, opt := range paymentOpts {
		paymentOptRes = append(paymentOptRes, new(dto.PaymentOptionRes).FromPaymentOption(opt))
	}
	return paymentOptRes, nil
}
