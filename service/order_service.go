package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrderItem(req *dto.OrderItemPostReq, userID uint) (*dto.OrderItemRes, error)
	FindOrderItemByUserID(userID uint) ([]*dto.OrderItemRes, error)
}

type orderService struct {
	db              *gorm.DB
	orderRepository repository.OrderRepository
	menuRepository  repository.MenuRepository
}

type OrderConfig struct {
	DB              *gorm.DB
	OrderRepository repository.OrderRepository
	MenuRepository  repository.MenuRepository
}

func NewOrder(c *OrderConfig) OrderService {
	return &orderService{
		db:              c.DB,
		orderRepository: c.OrderRepository,
		menuRepository:  c.MenuRepository,
	}
}

func (s *orderService) CreateOrderItem(req *dto.OrderItemPostReq, userID uint) (*dto.OrderItemRes, error) {
	oi := &model.OrderItem{
		UserID:      userID,
		MenuID:      req.MenuID,
		Quantity:    req.Quantity,
		Description: req.Description,
	}

	tx := s.db.Begin()
	item, err := s.orderRepository.CreateOrderItem(tx, oi)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	cart := &model.Cart{
		UserID:      userID,
		OrderItemID: item.ID,
	}
	_, err = s.orderRepository.CreateCart(tx, cart)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}
	helper.CommitOrRollback(tx, err)

	menuRes := new(dto.MenuRes).FromMenu(item.Menu)
	orderItemRes := new(dto.OrderItemRes).From(item, menuRes)
	return orderItemRes, nil
}

func (s *orderService) FindOrderItemByUserID(userID uint) ([]*dto.OrderItemRes, error) {
	tx := s.db.Begin()
	orderItems, err := s.orderRepository.FindOrderItemByUserID(tx, userID)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	var orderItemsRes []*dto.OrderItemRes
	for _, item := range orderItems {
		menuRes := new(dto.MenuRes).FromMenu(item.Menu)
		orderItemsRes = append(orderItemsRes, new(dto.OrderItemRes).From(item, menuRes))
	}

	return orderItemsRes, nil
}
