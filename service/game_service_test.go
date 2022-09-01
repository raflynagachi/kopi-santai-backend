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
)

var gameLeaderboard = model.GameLeaderboard{
	ID:     1,
	UserID: 1,
	Score:  100,
	User:   &user,
}

var game = model.Game{
	ID:          1,
	CouponID:    1,
	TargetScore: 100,
}

func TestGameService_FindByUserID(t *testing.T) {
	t.Run("should return response when find by userID success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		expectedRes := &dto.GameLeaderboardRes{
			UserID: 1,
			Name:   "John Doe",
			Score:  100,
		}
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&gameLeaderboard, nil)

		actualRes, err := s.FindByUserID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when find by userID failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.FindByUserID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestGameService_FindAll(t *testing.T) {
	t.Run("should return response when find all score success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		expectedRes := []*dto.GameLeaderboardRes{{
			UserID: 1,
			Name:   "John Doe",
			Score:  100,
		}}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return([]*model.GameLeaderboard{&gameLeaderboard}, nil)

		actualRes, err := s.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when find all score failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAll()

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestGameService_AddCouponPrizeToUser(t *testing.T) {
	t.Run("should return response when add coupon when user score reached success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		req := &dto.GameResultPostReq{
			Score: 100,
		}
		userCoupon := model.UserCoupon{
			User:   &model.User{},
			Coupon: &model.Coupon{},
		}
		expectedRes := &dto.UserCouponRes{
			UserRes:   &dto.UserRes{},
			CouponRes: &dto.CouponRes{},
		}
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&gameLeaderboard, nil)
		mockRepository.On("UpdateScore", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.GameLeaderboard")).Return(nil, nil)
		mockRepository.On("IsTargetScoreReached", mock.AnythingOfType(testutils.GormDBPointerType), gameLeaderboard.Score, req.Score).Return(&game, nil)
		couponMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&model.Coupon{ID: 1}, nil)
		couponMockRepository.On("AddCouponToUser", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.UserCoupon")).Return(&userCoupon, nil)

		actualRes, err := s.AddCouponPrizeToUser(req, uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when user game account not found failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		req := &dto.GameResultPostReq{
			Score: 100,
		}
		dbErr := errors.New("db error")
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.AddCouponPrizeToUser(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when update score failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		req := &dto.GameResultPostReq{
			Score: 100,
		}
		dbErr := errors.New("db error")
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&gameLeaderboard, nil)
		mockRepository.On("UpdateScore", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.GameLeaderboard")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.AddCouponPrizeToUser(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when user score reached in failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		req := &dto.GameResultPostReq{
			Score: 100,
		}
		dbErr := errors.New("db error")
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&gameLeaderboard, nil)
		mockRepository.On("UpdateScore", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.GameLeaderboard")).Return(nil, nil)
		mockRepository.On("IsTargetScoreReached", mock.AnythingOfType(testutils.GormDBPointerType), gameLeaderboard.Score, req.Score).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.AddCouponPrizeToUser(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when add coupon when user score reached failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		req := &dto.GameResultPostReq{
			Score: 100,
		}
		dbErr := new(apperror.CouponNotFoundError)
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&gameLeaderboard, nil)
		mockRepository.On("UpdateScore", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.GameLeaderboard")).Return(nil, nil)
		mockRepository.On("IsTargetScoreReached", mock.AnythingOfType(testutils.GormDBPointerType), gameLeaderboard.Score, req.Score).Return(&game, nil)
		couponMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.AddCouponPrizeToUser(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when add coupon when user score reached failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.GameRepository)
		couponMockRepository := new(mocks.CouponRepository)
		cfg := &service.GameConfig{
			DB:         gormDB,
			GameRepo:   mockRepository,
			CouponRepo: couponMockRepository,
		}
		s := service.NewGame(cfg)
		req := &dto.GameResultPostReq{
			Score: 100,
		}
		dbErr := errors.New("db error")
		mockRepository.On("FindByUserID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&gameLeaderboard, nil)
		mockRepository.On("UpdateScore", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.GameLeaderboard")).Return(nil, nil)
		mockRepository.On("IsTargetScoreReached", mock.AnythingOfType(testutils.GormDBPointerType), gameLeaderboard.Score, req.Score).Return(&game, nil)
		couponMockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(&model.Coupon{ID: 1}, nil)
		couponMockRepository.On("AddCouponToUser", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.UserCoupon")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.AddCouponPrizeToUser(req, uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
