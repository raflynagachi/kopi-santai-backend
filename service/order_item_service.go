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

type OrderItemService interface {
	CreateOrderItem(req *dto.OrderItemPostReq, userID uint) (*dto.OrderItemRes, error)
	FindOrderItemByUserID(userID uint) ([]*dto.OrderItemRes, error)
	UpdateOrderItemByID(id, userID uint, req *dto.OrderItemPatchReq) (*dto.OrderItemRes, error)
	DeleteOrderItemByID(id, userID uint) (gin.H, error)
}

type orderItemService struct {
	db                  *gorm.DB
	orderItemRepository repository.OrderItemRepository
	menuRepository      repository.MenuRepository
}

type OrderItemConfig struct {
	DB                  *gorm.DB
	OrderItemRepository repository.OrderItemRepository
	MenuRepository      repository.MenuRepository
}

func NewOrderItem(c *OrderItemConfig) OrderItemService {
	return &orderItemService{
		db:                  c.DB,
		orderItemRepository: c.OrderItemRepository,
		menuRepository:      c.MenuRepository,
	}
}

func checkOrderItem(ok bool, err error) error {
	if !ok && errors.Is(err, new(apperror.UserUnauthorizedError)) {
		return apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error())
	}
	if !ok && errors.Is(err, new(apperror.OrderItemNotFoundError)) {
		return apperror.NotFoundError(err.Error())
	}
	return nil
}

func (s *orderItemService) CreateOrderItem(req *dto.OrderItemPostReq, userID uint) (*dto.OrderItemRes, error) {
	oi := &model.OrderItem{
		UserID:      userID,
		MenuID:      req.MenuID,
		Quantity:    req.Quantity,
		Description: req.Description,
	}

	tx := s.db.Begin()
	item, err := s.orderItemRepository.CreateOrderItem(tx, oi)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	menuRes := new(dto.MenuRes).FromMenu(item.Menu)
	orderItemRes := new(dto.OrderItemRes).From(item, menuRes)
	return orderItemRes, nil
}

func (s *orderItemService) FindOrderItemByUserID(userID uint) ([]*dto.OrderItemRes, error) {
	tx := s.db.Begin()
	orderItems, err := s.orderItemRepository.FindOrderItemByUserID(tx, userID)
	helper.CommitOrRollback(tx, err)
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

func (s *orderItemService) UpdateOrderItemByID(id, userID uint, req *dto.OrderItemPatchReq) (*dto.OrderItemRes, error) {
	orderItem := &model.OrderItem{
		Quantity:    req.Quantity,
		Description: req.Description,
	}

	tx := s.db.Begin()
	ok, err := s.orderItemRepository.IsOrderItemOfUserID(tx, id, userID)
	err = checkOrderItem(ok, err)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	oi, err := s.orderItemRepository.UpdateOrderItemByID(tx, id, orderItem)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	menuRes := new(dto.MenuRes).FromMenu(oi.Menu)
	orderItemRes := new(dto.OrderItemRes).From(oi, menuRes)
	return orderItemRes, nil
}

func (s *orderItemService) DeleteOrderItemByID(id, userID uint) (gin.H, error) {
	tx := s.db.Begin()
	ok, err := s.orderItemRepository.IsOrderItemOfUserID(tx, id, userID)
	err = checkOrderItem(ok, err)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	isDeleted, err := s.orderItemRepository.DeleteOrderItemByID(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return gin.H{"isDeleted": false}, apperror.BadRequestError(err.Error())
	}

	return gin.H{"isDeleted": isDeleted}, nil
}
