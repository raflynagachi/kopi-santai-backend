package service_test

import (
	"errors"
	"testing"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_GetProfileDetail(t *testing.T) {
	t.Run("should return response when get profile detail success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		cfg := &service.UserConfig{
			DB:             gormDB,
			UserRepository: mockRepository,
		}
		s := service.NewUser(cfg)
		expectedRes := &dto.ProfileRes{
			User:    &dto.UserRes{},
			Coupons: []*dto.CouponRes(nil),
		}
		mockRepository.On("FindByIDWithCoupons", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&model.User{}, nil)

		actualRes, err := s.GetProfileDetail(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when get profile detail failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		cfg := &service.UserConfig{
			DB:             gormDB,
			UserRepository: mockRepository,
		}
		s := service.NewUser(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindByIDWithCoupons", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.GetProfileDetail(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestUserService_UpdateProfile(t *testing.T) {
	t.Run("should return response when update profile success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		cfg := &service.UserConfig{
			DB:             gormDB,
			UserRepository: mockRepository,
		}
		s := service.NewUser(cfg)
		expectedRes := &dto.UserRes{}
		req := &dto.UserUpdateReq{}
		mockRepository.On("Update", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), &model.User{}).Return(&model.User{}, nil)

		actualRes, err := s.UpdateProfile(uint(1), req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when update profile failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		cfg := &service.UserConfig{
			DB:             gormDB,
			UserRepository: mockRepository,
		}
		s := service.NewUser(cfg)
		req := &dto.UserUpdateReq{}
		dbErr := errors.New("db error")
		mockRepository.On("Update", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), &model.User{}).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.UpdateProfile(uint(1), req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
