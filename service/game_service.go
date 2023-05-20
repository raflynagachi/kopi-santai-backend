package service

import (
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"gorm.io/gorm"
)

type GameService interface {
	FindByUserID(userID uint) (*dto.GameLeaderboardRes, error)
	FindAll() ([]*dto.GameLeaderboardRes, error)
	AddCouponPrizeToUser(req *dto.GameResultPostReq, userID uint) (*dto.UserCouponRes, error)
}

type gameService struct {
	db         *gorm.DB
	gameRepo   repository.GameRepository
	couponRepo repository.CouponRepository
}

type GameConfig struct {
	DB         *gorm.DB
	GameRepo   repository.GameRepository
	CouponRepo repository.CouponRepository
}

func NewGame(c *GameConfig) GameService {
	return &gameService{
		db:         c.DB,
		gameRepo:   c.GameRepo,
		couponRepo: c.CouponRepo,
	}
}

func (s *gameService) FindByUserID(userID uint) (*dto.GameLeaderboardRes, error) {
	tx := s.db.Begin()
	gl, err := s.gameRepo.FindByUserID(tx, userID)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	glRes := new(dto.GameLeaderboardRes).FromGameLeaderboard(gl)
	return glRes, nil
}

func (s *gameService) FindAll() ([]*dto.GameLeaderboardRes, error) {
	tx := s.db.Begin()
	gls, err := s.gameRepo.FindAll(tx)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var glsRes []*dto.GameLeaderboardRes
	for _, gl := range gls {
		glsRes = append(glsRes, new(dto.GameLeaderboardRes).FromGameLeaderboard(gl))
	}

	return glsRes, nil
}

func (s *gameService) AddCouponPrizeToUser(req *dto.GameResultPostReq, userID uint) (*dto.UserCouponRes, error) {
	tx := s.db.Begin()
	gl, err := s.gameRepo.IsUserTriedLessThanX(tx, userID, model.MaxTried)
	if err != nil {
		tx.Rollback()
		return nil, apperror.BadRequestError(err.Error())
	}

	updatedGl := &model.GameLeaderboard{
		Score: req.Score + gl.Score,
		Tried: gl.Tried + 1,
	}
	updatedGl, err = s.gameRepo.UpdateScore(tx, userID, updatedGl)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	tx = s.db.Begin()
	game, err := s.gameRepo.IsTargetScoreReached(tx, gl.Score, req.Score)
	if err != nil {
		tx.Rollback()
		return nil, apperror.BadRequestError(err.Error())
	}

	coupon, err := s.couponRepo.FindByID(tx, game.CouponID)
	if err != nil {
		tx.Rollback()
		return nil, apperror.InternalServerError(new(apperror.CouponNotFoundError).Error())
	}

	uc := &model.UserCoupon{
		UserID:   userID,
		CouponID: coupon.ID,
	}
	userCoupon, err := s.couponRepo.AddCouponToUser(tx, uc)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	userCouponRes := new(dto.UserCouponRes).From(userCoupon)
	return userCouponRes, nil
}
