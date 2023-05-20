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

func TestReviewService_Create(t *testing.T) {
	t.Run("should return response when create review success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.ReviewRepository)
		cfg := &service.ReviewConfig{
			DB:         gormDB,
			ReviewRepo: mockRepository,
		}
		s := service.NewReview(cfg)
		expectedRes := &dto.ReviewRes{UserID: 1}
		req := &dto.ReviewPostReq{}
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &model.Review{UserID: 1}).Return(&model.Review{UserID: 1, User: &model.User{}}, nil)

		actualRes, err := s.Create(req, uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when create review failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.ReviewRepository)
		cfg := &service.ReviewConfig{
			DB:         gormDB,
			ReviewRepo: mockRepository,
		}
		s := service.NewReview(cfg)
		req := &dto.ReviewPostReq{}
		dbErr := errors.New("db error")
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &model.Review{UserID: 1}).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.Create(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestReviewService_FindByMenuID(t *testing.T) {
	t.Run("should return response when find menu reviews success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.ReviewRepository)
		cfg := &service.ReviewConfig{
			DB:         gormDB,
			ReviewRepo: mockRepository,
		}
		s := service.NewReview(cfg)
		expectedRes := []*dto.ReviewRes{{}}
		mockRepository.On("FindByMenuID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return([]*model.Review{{User: &model.User{}}}, nil)

		actualRes, err := s.FindByMenuID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when find menu reviews failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.ReviewRepository)
		cfg := &service.ReviewConfig{
			DB:         gormDB,
			ReviewRepo: mockRepository,
		}
		s := service.NewReview(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindByMenuID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.FindByMenuID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
