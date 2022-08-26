package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
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
	gl, err := s.gameRepo.FindByUserID(tx, userID)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	updatedGl := &model.GameLeaderboard{
		Score: req.Score + gl.Score,
	}
	updatedGl, err = s.gameRepo.UpdateScore(tx, userID, updatedGl)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	tx = s.db.Begin()
	game, err := s.gameRepo.IsTargetScoreReached(tx, gl.Score, req.Score)
	if err != nil {
		return nil, apperror.UnprocessableEntityError(err.Error())
	}

	uc := &model.UserCoupon{
		UserID:   userID,
		CouponID: game.CouponID,
	}
	userCoupon, err := s.couponRepo.AddCouponToUser(tx, uc)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	userCouponRes := new(dto.UserCouponRes).From(userCoupon)
	return userCouponRes, nil
}
