package service_test

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var orderID uint = 1

var order = &model.Order{
	ID:            1,
	CouponID:      &orderID,
	Coupon:        &model.Coupon{ID: 1},
	Delivery:      &delivery,
	PaymentOption: &model.PaymentOption{},
	OrderItems:    []*model.OrderItem{orderItem},
}

var orderRes = &dto.OrderRes{
	CouponID: 1,
	Coupon: &dto.CouponRes{
		ID: 1,
	},
	Delivery: &dto.DeliveryRes{
		Status: model.StatusDefault,
	},
	PaymentOption: &dto.PaymentOptionRes{},
	OrderItems: []*dto.OrderItemRes{{
		Menu: &dto.MenuRes{},
	}},
}

func TestOrderService_CreateOrder(t *testing.T) {
	t.Run("should return response when find orders by userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{CouponID: 1}
		expectedRes := orderRes
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(&model.UserCoupon{
			ID:     1,
			User:   &model.User{},
			Coupon: &model.Coupon{ID: 1},
		}, nil)
		couponMockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("CreateOrder", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Order")).Return(order, nil)
		orderItemMockRepository.On("UpdateOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0), orderItem).Return(nil, nil)
		promoMockRepository.On("FindByMinSpent", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return([]*model.Promotion{{}}, nil)
		couponMockRepository.On("AddCouponToUser", mock.AnythingOfType(testutils.GormDBPointerType), &model.UserCoupon{UserID: 1}).Return(nil, nil)

		menuRes, err := s.CreateOrder(req, uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when order item is empty", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{}
		dbErr := new(apperror.OrderItemNotFoundError)
		expectedErr := apperror.BadRequestError(new(apperror.OrderItemsEmptyError).Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to find order item by user id", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to create delivery", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to find payment options", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{}
		dbErr := errors.New("db error")
		expectedErr := apperror.NotFoundError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to find coupon", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{
			CouponID: 1,
		}
		dbErr := errors.New("db error")
		expectedErr := apperror.NotFoundError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to remove coupon", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{CouponID: 1}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(&model.UserCoupon{
			ID:     1,
			User:   &model.User{},
			Coupon: &model.Coupon{ID: 1},
		}, nil)
		couponMockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to create order", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{CouponID: 1}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(&model.UserCoupon{
			ID:     1,
			User:   &model.User{},
			Coupon: &model.Coupon{ID: 1},
		}, nil)
		couponMockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("CreateOrder", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Order")).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to update order item", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{CouponID: 1}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(&model.UserCoupon{
			ID:     1,
			User:   &model.User{},
			Coupon: &model.Coupon{ID: 1},
		}, nil)
		couponMockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("CreateOrder", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Order")).Return(&model.Order{}, nil)
		orderItemMockRepository.On("UpdateOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0), orderItem).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to apply promo", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{CouponID: 1}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(&model.UserCoupon{
			ID:     1,
			User:   &model.User{},
			Coupon: &model.Coupon{ID: 1},
		}, nil)
		couponMockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("CreateOrder", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Order")).Return(&model.Order{}, nil)
		orderItemMockRepository.On("UpdateOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0), orderItem).Return(nil, nil)
		promoMockRepository.On("FindByMinSpent", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to add coupon to user", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		req := &dto.OrderPostReq{CouponID: 1}
		dbErr := errors.New("db error")
		expectedErr := apperror.InternalServerError(dbErr.Error())
		orderItemMockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)
		deliveryMockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)
		paymentOptMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, nil)
		couponMockRepository.On("FindUserCouponByCouponID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(&model.UserCoupon{
			ID:     1,
			User:   &model.User{},
			Coupon: &model.Coupon{ID: 1},
		}, nil)
		couponMockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("CreateOrder", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Order")).Return(&model.Order{}, nil)
		orderItemMockRepository.On("UpdateOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(0), orderItem).Return(nil, nil)
		promoMockRepository.On("FindByMinSpent", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return([]*model.Promotion{{}}, nil)
		couponMockRepository.On("AddCouponToUser", mock.AnythingOfType(testutils.GormDBPointerType), &model.UserCoupon{UserID: 1}).Return(nil, dbErr)

		_, err := s.CreateOrder(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestOrderService_FindOrderByIDAndUserID(t *testing.T) {
	t.Run("should return response when find order by id and userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		expectedRes := orderRes
		mockRepository.On("FindOrderByIDAndUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(order, nil)
		orderItemMockRepository.On("FindOrderItemByUserIDAndOrderID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return([]*model.OrderItem{orderItem}, nil)

		menuRes, err := s.FindOrderByIDAndUserID(uint(1), uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when failed to find order", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		dbErr := errors.New("db error")
		expectedErr := apperror.NotFoundError(dbErr.Error())
		mockRepository.On("FindOrderByIDAndUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(nil, dbErr)

		_, err := s.FindOrderByIDAndUserID(uint(1), uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to find order item", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindOrderByIDAndUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(order, nil)
		orderItemMockRepository.On("FindOrderItemByUserIDAndOrderID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.FindOrderByIDAndUserID(uint(1), uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestOrderService_FindOrderByUserID(t *testing.T) {
	t.Run("should return response when find orders by userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		expectedRes := []*dto.OrderRes{orderRes}
		mockRepository.On("FindOrderByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.Order{order}, nil)
		orderItemMockRepository.On("FindOrderItemByUserIDAndOrderID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return([]*model.OrderItem{orderItem}, nil)

		menuRes, err := s.FindOrderByUserID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when failed to find order", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindOrderByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.Order{order}, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.FindOrderByUserID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to find order item", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindOrderByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.Order{order}, nil)
		orderItemMockRepository.On("FindOrderItemByUserIDAndOrderID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.FindOrderByUserID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestOrderService_FindAll(t *testing.T) {
	t.Run("should return response when find all orders success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		expectedRes := []*dto.OrderRes{orderRes}
		q := &model.QueryParamOrder{}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType), &time.Time{}).Return([]*model.Order{order}, nil)
		orderItemMockRepository.On("FindOrderItemByOrderID", mock.AnythingOfType(testutils.GormDBPointerType), orderID).Return([]*model.OrderItem{orderItem}, nil)

		menuRes, err := s.FindAll(q)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return response when find all orders with query param success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		expectedRes := []*dto.OrderRes{orderRes}
		q := &model.QueryParamOrder{
			Date: "lastYear",
		}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*time.Time")).Return([]*model.Order{order}, nil)
		orderItemMockRepository.On("FindOrderItemByOrderID", mock.AnythingOfType(testutils.GormDBPointerType), orderID).Return([]*model.OrderItem{orderItem}, nil)

		menuRes, err := s.FindAll(q)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when failed to find order", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		q := &model.QueryParamOrder{
			Date: "lastWeek",
		}
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*time.Time")).Return([]*model.Order{order}, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.FindAll(q)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to find order items", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderRepository)
		deliveryMockRepository := new(mocks.DeliveryRepository)
		paymentOptMockRepository := new(mocks.PaymentOptionRepository)
		couponMockRepository := new(mocks.CouponRepository)
		orderItemMockRepository := new(mocks.OrderItemRepository)
		promoMockRepository := new(mocks.PromotionRepository)
		cfg := &service.OrderConfig{
			DB:             gormDB,
			DeliveryRepo:   deliveryMockRepository,
			PaymentOptRepo: paymentOptMockRepository,
			CouponRepo:     couponMockRepository,
			OrderRepo:      mockRepository,
			OrderItemRepo:  orderItemMockRepository,
			PromotionRepo:  promoMockRepository,
		}
		s := service.NewOrder(cfg)
		q := &model.QueryParamOrder{
			Date: "lastMonth",
		}
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*time.Time")).Return([]*model.Order{order}, nil)
		orderItemMockRepository.On("FindOrderItemByOrderID", mock.AnythingOfType(testutils.GormDBPointerType), orderID).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.FindAll(q)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
