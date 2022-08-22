package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
	"time"
)

type OrderRepository interface {
	CreateOrder(tx *gorm.DB, o *model.Order) (*model.Order, error)
	FindAll(tx *gorm.DB, t *time.Time) ([]*model.Order, error)
	FindOrderByIDAndUserID(tx *gorm.DB, id, userID uint) (*model.Order, error)
	FindOrderByUserID(tx *gorm.DB, userID uint) ([]*model.Order, error)
}

type orderRepository struct{}

func NewOrder() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) CreateOrder(tx *gorm.DB, o *model.Order) (*model.Order, error) {
	err := tx.Preload("Coupon").Preload("Delivery").Preload("PaymentOption").Create(&o).First(&o).Error
	return o, err
}

func (r *orderRepository) FindAll(tx *gorm.DB, t *time.Time) ([]*model.Order, error) {
	var orders []*model.Order

	result := tx.Preload("Coupon").Preload("Delivery").Preload("PaymentOption").Where("ordered_date BETWEEN ? AND ?", t, time.Now()).Find(&orders)
	if result.Error != nil && result.RowsAffected == 0 {
		return nil, new(apperror.OrderNotFoundError)
	}
	return orders, nil
}

func (r *orderRepository) FindOrderByIDAndUserID(tx *gorm.DB, id, userID uint) (*model.Order, error) {
	var o *model.Order
	result := tx.Preload("Coupon").Preload("Delivery").Preload("PaymentOption").Where("user_id = ?", userID).First(&o, id)
	if result.Error != nil && result.RowsAffected == 0 {
		return nil, new(apperror.OrderNotFoundError)
	}
	return o, result.Error
}

func (r *orderRepository) FindOrderByUserID(tx *gorm.DB, userID uint) ([]*model.Order, error) {
	var orders []*model.Order

	result := tx.Preload("Coupon").Preload("Delivery").Preload("PaymentOption").Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil && result.RowsAffected == 0 {
		return nil, new(apperror.OrderNotFoundError)
	}
	return orders, result.Error
}
