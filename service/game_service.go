package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
)

type GameService interface {
	FindByUserID(userID uint) (*dto.GameLeaderboardRes, error)
	FindAll() ([]*dto.GameLeaderboardRes, error)
}

type gameService struct {
	db       *gorm.DB
	gameRepo repository.GameRepository
}

type GameConfig struct {
	DB       *gorm.DB
	GameRepo repository.GameRepository
}

func NewGame(c *GameConfig) GameService {
	return &gameService{
		db:       c.DB,
		gameRepo: c.GameRepo,
	}
}

func (s *gameService) FindByUserID(userID uint) (*dto.GameLeaderboardRes, error) {
	tx := s.db.Begin()
	gl, err := s.gameRepo.FindByUserID(tx, userID)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	glRes := new(dto.GameLeaderboardRes).FromGameLeaderboard(gl)
	return glRes, nil
}

func (s *gameService) FindAll() ([]*dto.GameLeaderboardRes, error) {
	tx := s.db.Begin()
	gls, err := s.gameRepo.FindAll(tx)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var glsRes []*dto.GameLeaderboardRes
	for _, gl := range gls {
		glsRes = append(glsRes, new(dto.GameLeaderboardRes).FromGameLeaderboard(gl))
	}

	return glsRes, nil
}
