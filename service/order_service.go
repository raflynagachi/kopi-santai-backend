package service

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrderItem(req *dto.OrderItemPostReq, userID uint) (*dto.OrderItemRes, error)
	FindOrderItemByUserID(userID uint) ([]*dto.OrderItemRes, error)
	UpdateOrderItemByID(id, userID uint, req *dto.OrderItemPatchReq) (*dto.OrderItemRes, error)
	DeleteOrderItemByID(id, userID uint) (gin.H, error)
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

func checkOrderItem(ok bool, err error) error {
	if !ok && errors.Is(err, new(apperror.UserUnauthorizedError)) {
		return apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error())
	}
	if !ok && errors.Is(err, new(apperror.OrderNotFoundError)) {
		return apperror.NotFoundError(err.Error())
	}
	return nil
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
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

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

func (s *orderService) UpdateOrderItemByID(id, userID uint, req *dto.OrderItemPatchReq) (*dto.OrderItemRes, error) {
	orderItem := &model.OrderItem{
		Quantity:    req.Quantity,
		Description: req.Description,
	}

	tx := s.db.Begin()
	ok, err := s.orderRepository.IsOrderItemOfUserID(tx, id, userID)
	err = checkOrderItem(ok, err)
	if err != nil {
		return nil, err
	}

	oi, err := s.orderRepository.UpdateOrderItemByID(tx, id, orderItem)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	menuRes := new(dto.MenuRes).FromMenu(oi.Menu)
	orderItemRes := new(dto.OrderItemRes).From(oi, menuRes)
	return orderItemRes, nil
}

func (s *orderService) DeleteOrderItemByID(id, userID uint) (gin.H, error) {
	tx := s.db.Begin()
	ok, err := s.orderRepository.IsOrderItemOfUserID(tx, id, userID)
	err = checkOrderItem(ok, err)
	if err != nil {
		return nil, err
	}

	isDeleted, err := s.orderRepository.DeleteOrderItemByID(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return gin.H{"isDeleted": false}, apperror.BadRequestError(err.Error())
	}

	return gin.H{"isDeleted": isDeleted}, nil
}
