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

var coupon = model.Coupon{
	Name:   "Santai Coupon",
	Amount: 20,
}

func TestCouponService_Create(t *testing.T) {
	t.Run("should return response when create coupon success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		req := &dto.CouponPostReq{
			Name:   "Santai Coupon",
			Amount: 20,
		}
		couponCreated := coupon
		couponCreated.ID = 1
		expectedRes := &dto.CouponRes{
			ID:     1,
			Name:   "Santai Coupon",
			Amount: 20,
		}
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &coupon).Return(&couponCreated, nil)

		actualRes, err := s.Create(req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when create coupon failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		req := &dto.CouponPostReq{
			Name:   "Santai Coupon",
			Amount: 20,
		}
		dbErr := errors.New("db error")
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &coupon).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.Create(req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestCouponService_FindCouponByUserID(t *testing.T) {
	t.Run("should return response when find coupon by userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		userCoupon := []*model.UserCoupon{{
			ID:       1,
			UserID:   1,
			CouponID: 1,
			Coupon:   &model.Coupon{},
		}}
		expectedRes := []*dto.CouponRes{{}}
		mockRepository.On("FindCouponByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(userCoupon, nil)

		actualRes, err := s.FindCouponByUserID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when find coupon by userID failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		dbErr := errors.New("db error")
		mockRepository.On("FindCouponByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindCouponByUserID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestCouponService_FindAll(t *testing.T) {
	t.Run("should return response when find all coupon success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		expectedRes := []*dto.CouponRes{{}}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return([]*model.Coupon{{}}, nil)

		actualRes, err := s.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when find all coupon failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAll()

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestCouponService_DeleteByID(t *testing.T) {
	t.Run("should return response when delete coupon success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		expectedRes := gin.H{"isDeleted": true}
		mockRepository.On("DeleteByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)

		actualRes, err := s.DeleteByID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when delete user coupon failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		dbErr := errors.New("db error")
		mockRepository.On("DeleteByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)
		mockRepository.On("DeleteUserCoupon", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.DeleteByID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when delete coupon failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CouponRepository)
		couponConfig := &service.CouponConfig{DB: gormDB, CouponRepo: mockRepository}
		s := service.NewCoupon(couponConfig)
		dbErr := errors.New("db error")
		mockRepository.On("DeleteByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.DeleteByID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
