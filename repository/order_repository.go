package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrderByUserID(tx *gorm.DB, userID uint) (*model.Order, error)
	FindOrderItemByUserID(tx *gorm.DB, userID uint) ([]*model.OrderItem, error)
	CreateCart(tx *gorm.DB, cart *model.Cart) (*model.Cart, error)
	CreateOrderItem(tx *gorm.DB, oi *model.OrderItem) (*model.OrderItem, error)
	UpdateOrderItemByID(tx *gorm.DB, id uint, oi *model.OrderItem) (*model.OrderItem, error)
	IsOrderItemOfUserID(tx *gorm.DB, id, userID uint) bool
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

func (r *orderRepository) FindOrderItemByUserID(tx *gorm.DB, userID uint) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	err := tx.Preload("Menu").Preload("Menu.Category").Where("user_id = ?", userID).Find(&orderItems).Error
	return orderItems, err
}

func (r *orderRepository) CreateCart(tx *gorm.DB, cart *model.Cart) (*model.Cart, error) {
	err := tx.Create(&cart).Error
	return cart, err
}

func (r *orderRepository) CreateOrderItem(tx *gorm.DB, oi *model.OrderItem) (*model.OrderItem, error) {
	err := tx.Preload("Menu").Preload("Menu.Category").Create(&oi).First(&oi).Error
	return oi, err
}

func (r *orderRepository) IsOrderItemOfUserID(tx *gorm.DB, id, userID uint) bool {
	var oi *model.OrderItem
	result := tx.Where("user_id = ?", userID).First(&oi, id)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (r *orderRepository) UpdateOrderItemByID(tx *gorm.DB, id uint, oi *model.OrderItem) (*model.OrderItem, error) {
	var updatedOrderItem *model.OrderItem
	err := tx.First(&updatedOrderItem, id).Updates(&oi).Error
	if err != nil {
		return nil, err
	}
	_ = tx.Preload("Menu").Preload("Menu.Category").First(&updatedOrderItem, id)
	return updatedOrderItem, nil
}
