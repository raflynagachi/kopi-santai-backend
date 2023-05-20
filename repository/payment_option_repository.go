package repository

import (
	"github.com/raflynagachi/kopi-santai-backend/model"
	"gorm.io/gorm"
)

type PaymentOptionRepository interface {
	FindAll(tx *gorm.DB) ([]*model.PaymentOption, error)
	FindByID(tx *gorm.DB, id uint) (*model.PaymentOption, error)
}

type paymentOptionRepository struct{}

func NewPaymentOption() PaymentOptionRepository {
	return &paymentOptionRepository{}
}

func (r *paymentOptionRepository) FindAll(tx *gorm.DB) ([]*model.PaymentOption, error) {
	var po []*model.PaymentOption
	err := tx.Find(&po).Error
	return po, err
}

func (r *paymentOptionRepository) FindByID(tx *gorm.DB, id uint) (*model.PaymentOption, error) {
	var po *model.PaymentOption
	err := tx.First(&po, id).Error
	return po, err
}
