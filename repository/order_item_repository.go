package repository

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	FindOrderItemByUserIDAndOrderID(tx *gorm.DB, userID, orderID uint) ([]*model.OrderItem, error)
	FindOrderItemByUserID(tx *gorm.DB, userID uint) ([]*model.OrderItem, error)
	FindOrderItemByOrderID(tx *gorm.DB, orderID uint) ([]*model.OrderItem, error)
	CreateOrderItem(tx *gorm.DB, oi *model.OrderItem) (*model.OrderItem, error)
	UpdateOrderItemByID(tx *gorm.DB, id uint, oi *model.OrderItem) (*model.OrderItem, error)
	DeleteOrderItemByID(tx *gorm.DB, id uint) (bool, error)
	IsOrderItemOfUserID(tx *gorm.DB, id, userID uint) (bool, error)
}

type orderItemRepository struct{}

func NewOrderItem() OrderItemRepository {
	return &orderItemRepository{}
}

func (r *orderItemRepository) FindOrderItemByUserIDAndOrderID(tx *gorm.DB, userID, orderID uint) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	result := tx.Preload("Menu").Preload("Menu.Category").Where("user_id = ? AND order_id = ?", userID, orderID).Find(&orderItems)
	if result.RowsAffected == 0 {
		return orderItems, new(apperror.OrderItemNotFoundError)
	}
	return orderItems, result.Error
}

func (r *orderItemRepository) FindOrderItemByUserID(tx *gorm.DB, userID uint) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	result := tx.Preload("Menu").Preload("Menu.Category").Where("user_id = ? AND is_active = ? AND order_id IS NULL", userID, true).Find(&orderItems)
	if result.RowsAffected == 0 {
		return orderItems, new(apperror.OrderItemNotFoundError)
	}
	return orderItems, result.Error
}

func (r *orderItemRepository) FindOrderItemByOrderID(tx *gorm.DB, orderID uint) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	result := tx.Preload("Menu").Preload("Menu.Category").Where("order_id = ?", orderID).Find(&orderItems)
	if result.RowsAffected == 0 {
		return orderItems, new(apperror.OrderItemNotFoundError)
	}
	return orderItems, result.Error
}

func (r *orderItemRepository) CreateOrderItem(tx *gorm.DB, oi *model.OrderItem) (*model.OrderItem, error) {
	err := tx.Preload("Menu").Preload("Menu.Category").Create(&oi).First(&oi).Error
	return oi, err
}

func (r *orderItemRepository) IsOrderItemOfUserID(tx *gorm.DB, id, userID uint) (bool, error) {
	var oi *model.OrderItem
	result := tx.First(&oi, id)
	if result.RowsAffected == 0 {
		return false, new(apperror.OrderItemNotFoundError)
	}

	result = tx.Where("user_id = ?", userID).First(&oi, id)
	if result.RowsAffected == 0 {
		return false, new(apperror.UserUnauthorizedError)
	}
	return true, nil
}

func (r *orderItemRepository) UpdateOrderItemByID(tx *gorm.DB, id uint, oi *model.OrderItem) (*model.OrderItem, error) {
	var updatedOrderItem *model.OrderItem
	err := tx.First(&updatedOrderItem, id).Updates(&oi).Error
	if err != nil {
		return nil, err
	}
	_ = tx.Preload("Menu").Preload("Menu.Category").First(&updatedOrderItem, id)
	return updatedOrderItem, nil
}

func (r *orderItemRepository) DeleteOrderItemByID(tx *gorm.DB, id uint) (bool, error) {
	var updatedOrderItem *model.OrderItem
	err := tx.Delete(&updatedOrderItem, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
