package service_test

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPromotionService_FindAll(t *testing.T) {
	t.Run("should return response when find all promotion success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		expectedRes := []*dto.PromotionRes{{
			Coupon: &dto.CouponRes{},
		}}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return([]*model.Promotion{{
			Coupon: &model.Coupon{},
		}}, nil)

		menuRes, err := s.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when find all promotion failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAll()

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestPromotionService_FindAllUnscoped(t *testing.T) {
	t.Run("should return response when find all promotion unscoped success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		expectedRes := []*dto.PromotionRes{{
			Coupon: &dto.CouponRes{},
		}}
		mockRepository.On("FindAllUnscoped", mock.AnythingOfType(testutils.GormDBPointerType)).Return([]*model.Promotion{{
			Coupon: &model.Coupon{},
		}}, nil)

		menuRes, err := s.FindAllUnscoped()

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when find all promotion unscoped failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindAllUnscoped", mock.AnythingOfType(testutils.GormDBPointerType)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAllUnscoped()

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestPromotionService_CreatePromotion(t *testing.T) {
	t.Run("should return response when create promotion success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		req := &dto.PromotionPostReq{}
		expectedRes := &dto.PromotionRes{Coupon: &dto.CouponRes{}}
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Promotion")).Return(&model.Promotion{Coupon: &model.Coupon{}}, nil)

		menuRes, err := s.CreatePromotion(req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when create promotion failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		req := &dto.PromotionPostReq{}
		dbErr := errors.New("db error")
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Promotion")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.CreatePromotion(req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestPromotionService_DeletePromotionByID(t *testing.T) {
	t.Run("should return response when delete promotion by id success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		expectedRes := gin.H{"isDeleted": true}
		mockRepository.On("Delete", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)

		menuRes, err := s.DeletePromotionByID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when delete promotion by id failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PromotionRepository)
		cfg := &service.PromoConfig{
			DB:        gormDB,
			PromoRepo: mockRepository,
		}
		s := service.NewPromo(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("Delete", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.DeletePromotionByID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
