package service_test

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var orderItem = &model.OrderItem{
	OrderID: &orderID,
	Menu: &model.Menu{
		Category: &model.Category{},
		Reviews:  []*model.Review{{}},
	},
}

func TestOrderItemService_CreateOrderItem(t *testing.T) {
	t.Run("should return response when create order item success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		req := &dto.OrderItemPostReq{}
		expectedRes := &dto.OrderItemRes{
			Menu: &dto.MenuRes{},
		}
		mockRepository.On("CreateOrderItem", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.OrderItem")).Return(orderItem, nil)

		menuRes, err := s.CreateOrderItem(req, uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when create order item failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		req := &dto.OrderItemPostReq{}
		dbErr := errors.New("db error")
		mockRepository.On("CreateOrderItem", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.OrderItem")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.CreateOrderItem(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestOrderItemService_FindOrderItemByUserID(t *testing.T) {
	t.Run("should return response when find order item by userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		expectedRes := []*dto.OrderItemRes{{
			Menu: &dto.MenuRes{},
		}}
		mockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.OrderItem{orderItem}, nil)

		menuRes, err := s.FindOrderItemByUserID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when find order item by userID failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.FindOrderItemByUserID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestOrderItemService_UpdateOrderItemByID(t *testing.T) {
	t.Run("should return response when update order item by id success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		req := &dto.OrderItemPatchReq{}
		expectedRes := &dto.OrderItemRes{
			Menu: &dto.MenuRes{},
		}
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(true, nil)
		mockRepository.On("UpdateOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.OrderItem")).Return(orderItem, nil)

		menuRes, err := s.UpdateOrderItemByID(uint(1), uint(1), req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when order item accessed by unauthorized user", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		req := &dto.OrderItemPatchReq{}
		dbErr := new(apperror.UserUnauthorizedError)
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(false, dbErr)
		expectedErr := apperror.UnauthorizedError(dbErr.Error())

		_, err := s.UpdateOrderItemByID(uint(1), uint(1), req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when order item not found", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		req := &dto.OrderItemPatchReq{}
		dbErr := new(apperror.OrderItemNotFoundError)
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(false, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.UpdateOrderItemByID(uint(1), uint(1), req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when update order item by id failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		req := &dto.OrderItemPatchReq{}
		dbErr := errors.New("db error")
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(true, nil)
		mockRepository.On("UpdateOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.OrderItem")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.UpdateOrderItemByID(uint(1), uint(1), req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestOrderItemService_DeleteOrderItemByID(t *testing.T) {
	t.Run("should return response when delete order item by id success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		expectedRes := gin.H{"isDeleted": true}
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(true, nil)
		mockRepository.On("DeleteOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)

		menuRes, err := s.DeleteOrderItemByID(uint(1), uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when order item accessed by unauthorized user", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		dbErr := new(apperror.UserUnauthorizedError)
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(false, dbErr)
		expectedErr := apperror.UnauthorizedError(dbErr.Error())

		_, err := s.DeleteOrderItemByID(uint(1), uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when order item not found", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		dbErr := new(apperror.OrderItemNotFoundError)
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(false, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.DeleteOrderItemByID(uint(1), uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when delete order item by id failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		dbErr := errors.New("db error")
		expectedRes := gin.H{"isDeleted": false}
		mockRepository.On("IsOrderItemOfUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), uint(1)).Return(true, nil)
		mockRepository.On("DeleteOrderItemByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		res, err := s.DeleteOrderItemByID(uint(1), uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
		assert.Equal(t, expectedRes, res)
	})
}

func TestOrderItemService_DeleteOrderItemByUserID(t *testing.T) {
	t.Run("should return response when delete order item by userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		expectedRes := gin.H{"isDeleted": true}
		mockRepository.On("DeleteOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)

		menuRes, err := s.DeleteOrderItemByUserID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return response when delete order item by userID failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.OrderItemRepository)
		menuRepository := new(mocks.MenuRepository)
		cfg := &service.OrderItemConfig{
			DB:                  gormDB,
			OrderItemRepository: mockRepository,
			MenuRepository:      menuRepository,
		}
		s := service.NewOrderItem(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("DeleteOrderItemByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.DeleteOrderItemByUserID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
