package service

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
	"time"
)

type OrderService interface {
	CreateOrder(req *dto.OrderPostReq, userID uint) (*dto.OrderRes, error)
	FindAll(q *model.QueryParamOrder) ([]*dto.OrderRes, error)
	FindOrderByIDAndUserID(id, userID uint) (*dto.OrderRes, error)
	FindOrderByUserID(userID uint) ([]*dto.OrderRes, error)
}

type orderService struct {
	db             *gorm.DB
	deliveryRepo   repository.DeliveryRepository
	paymentOptRepo repository.PaymentOptionRepository
	couponRepo     repository.CouponRepository
	orderRepo      repository.OrderRepository
	orderItemRepo  repository.OrderItemRepository
}

type OrderConfig struct {
	DB             *gorm.DB
	DeliveryRepo   repository.DeliveryRepository
	PaymentOptRepo repository.PaymentOptionRepository
	CouponRepo     repository.CouponRepository
	OrderRepo      repository.OrderRepository
	OrderItemRepo  repository.OrderItemRepository
}

func NewOrder(c *OrderConfig) OrderService {
	return &orderService{
		db:             c.DB,
		deliveryRepo:   c.DeliveryRepo,
		paymentOptRepo: c.PaymentOptRepo,
		couponRepo:     c.CouponRepo,
		orderRepo:      c.OrderRepo,
		orderItemRepo:  c.OrderItemRepo,
	}
}

func validateQueryParamOrder(q *model.QueryParamOrder) time.Time {
	if q.Date == "lastWeek" {
		return time.Now().AddDate(0, 0, -7)
	}
	if q.Date == "lastMonth" {
		return time.Now().AddDate(0, -1, 0)
	}
	if q.Date == "lastYear" {
		return time.Now().AddDate(-1, 0, 0)
	}
	return time.Time{}
}

func (s *orderService) CreateOrder(req *dto.OrderPostReq, userID uint) (*dto.OrderRes, error) {
	o := &model.Order{
		UserID:          userID,
		PaymentOptionID: req.PaymentOptID,
		OrderedDate:     req.OrderedDate,
		IsActive:        true,
	}

	tx := s.db.Begin()

	var totalPrice float64
	orderItems, err := s.orderItemRepo.FindOrderItemByUserID(tx, userID)
	if err != nil {
		if errors.Is(err, new(apperror.OrderItemNotFoundError)) {
			return nil, apperror.BadRequestError(new(apperror.OrderItemsEmptyError).Error())
		}
		return nil, apperror.InternalServerError(err.Error())
	}
	for _, item := range orderItems {
		totalPrice += float64(item.Quantity) * item.Menu.Price
	}

	d := &model.Delivery{
		DeliveryDate: req.OrderedDate,
		Status:       model.StatusDefault,
	}
	delivery, err := s.deliveryRepo.Create(tx, d)
	o.DeliveryID = delivery.ID
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	_, err = s.paymentOptRepo.FindByID(tx, req.PaymentOptID)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	if req.CouponID != 0 {
		coupon, err := s.couponRepo.FindByID(tx, req.CouponID)
		if err != nil {
			return nil, apperror.NotFoundError(err.Error())
		}
		if time.Now().Before(coupon.ExpiredDate) && coupon.IsAvailable && coupon.MinSpent <= totalPrice {
			totalPrice -= coupon.Amount
			o.CouponID = &req.CouponID
		} else {
			return nil, apperror.BadRequestError(new(apperror.CouponFailedError).Error())
		}
	}

	o.TotalPrice = totalPrice
	order, err := s.orderRepo.CreateOrder(tx, o)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	for _, item := range orderItems {
		item.IsActive = false
		item.OrderID = &order.ID
		item, err = s.orderItemRepo.UpdateOrderItemByID(tx, item.ID, item)
		if err != nil {
			return nil, apperror.InternalServerError(err.Error())
		}
	}
	helper.CommitOrRollback(tx, err)

	o.OrderItems = orderItems
	orderRes := new(dto.OrderRes).From(order)
	return orderRes, nil
}

func (s *orderService) FindOrderByIDAndUserID(id, userID uint) (*dto.OrderRes, error) {
	tx := s.db.Begin()
	order, err := s.orderRepo.FindOrderByIDAndUserID(tx, id, userID)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}
	orderItems, err := s.orderItemRepo.FindOrderItemByUserIDAndOrderID(tx, userID, id)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	order.OrderItems = orderItems
	orderRes := new(dto.OrderRes).From(order)
	return orderRes, nil
}

func (s *orderService) FindOrderByUserID(userID uint) ([]*dto.OrderRes, error) {
	var ordersRes []*dto.OrderRes

	tx := s.db.Begin()
	order, err := s.orderRepo.FindOrderByUserID(tx, userID)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	for _, o := range order {
		orderItems, err := s.orderItemRepo.FindOrderItemByUserIDAndOrderID(tx, userID, o.ID)
		if err != nil {
			return nil, apperror.NotFoundError(err.Error())
		}
		o.OrderItems = orderItems
		ordersRes = append(ordersRes, new(dto.OrderRes).From(o))
	}

	return ordersRes, nil
}

func (s *orderService) FindAll(q *model.QueryParamOrder) ([]*dto.OrderRes, error) {
	var ordersRes []*dto.OrderRes
	t := validateQueryParamOrder(q)

	tx := s.db.Begin()
	order, err := s.orderRepo.FindAll(tx, &t)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	for _, o := range order {
		orderItems, err := s.orderItemRepo.FindOrderItemByOrderID(tx, o.ID)
		if err != nil {
			return nil, apperror.NotFoundError(err.Error())
		}
		o.OrderItems = orderItems
		ordersRes = append(ordersRes, new(dto.OrderRes).From(o))
	}

	return ordersRes, nil
}
