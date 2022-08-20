package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrderByUserID(tx *gorm.DB, userID uint) (*model.Order, error)
	CreateCart(tx *gorm.DB, cart *model.Cart) (*model.Cart, error)
	CreateOrderItem(tx *gorm.DB, oi *model.OrderItem) (*model.OrderItem, error)
}

type orderRepository struct{}

func NewOrder() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) FindOrderByUserID(tx *gorm.DB, userID uint) (*model.Order, error) {
	var o *model.Order
	result := tx.Where("user_id = ? AND is_completed = false", userID).First(&o)
	if result.Error != nil && result.RowsAffected == 0 {
		return nil, new(apperror.OrderNotFoundError)
	}
	return o, result.Error
}

func (r *orderRepository) CreateCart(tx *gorm.DB, cart *model.Cart) (*model.Cart, error) {
	err := tx.Create(&cart).Error
	return cart, err
}

func (r *orderRepository) CreateOrderItem(tx *gorm.DB, oi *model.OrderItem) (*model.OrderItem, error) {
	err := tx.Create(&oi).Error
	return oi, err
}
