package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
	"time"
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
