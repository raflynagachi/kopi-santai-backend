package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type DeliveryRepository interface {
	Create(tx *gorm.DB, d *model.Delivery) (*model.Delivery, error)
	FindAll(tx *gorm.DB) ([]*model.Delivery, error)
	FindByID(tx *gorm.DB, id uint) (*model.Delivery, error)
}

type deliveryRepository struct{}

func NewDelivery() DeliveryRepository {
	return &deliveryRepository{}
}

func (r *deliveryRepository) Create(tx *gorm.DB, d *model.Delivery) (*model.Delivery, error) {
	err := tx.Create(&d).Error
	return d, err
}

func (r *deliveryRepository) FindAll(tx *gorm.DB) ([]*model.Delivery, error) {
	var d []*model.Delivery
	err := tx.Find(&d).Error
	return d, err
}

func (r *deliveryRepository) FindByID(tx *gorm.DB, id uint) (*model.Delivery, error) {
	var d *model.Delivery
	err := tx.First(&d, id).Error
	return d, err
}
