package service

import (
	"time"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"gorm.io/gorm"
)

type DeliveryService interface {
	UpdateStatus(id uint, req *dto.DeliveryUpdateStatusReq) (*dto.DeliveryRes, error)
}

type deliveryService struct {
	db           *gorm.DB
	deliveryRepo repository.DeliveryRepository
}

type DeliveryConfig struct {
	DB           *gorm.DB
	DeliveryRepo repository.DeliveryRepository
}

func NewDelivery(c *DeliveryConfig) DeliveryService {
	return &deliveryService{
		db:           c.DB,
		deliveryRepo: c.DeliveryRepo,
	}
}

func (s *deliveryService) UpdateStatus(id uint, req *dto.DeliveryUpdateStatusReq) (*dto.DeliveryRes, error) {
	d := &model.Delivery{
		DeliveryDate: time.Now(),
		Status:       req.Status,
	}

	tx := s.db.Begin()
	delivery, err := s.deliveryRepo.Update(tx, id, d)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	deliveryRes := new(dto.DeliveryRes).FromDelivery(delivery)
	return deliveryRes, nil
}
