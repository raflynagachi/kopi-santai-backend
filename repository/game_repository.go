package repository

import (
	"fmt"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GameRepository interface {
	CreateLeaderboard(tx *gorm.DB, gl *model.GameLeaderboard) (*model.GameLeaderboard, error)
	IsTargetScoreReached(tx *gorm.DB, prevScore, score uint) (*model.Game, error)
	IsUserTriedLessThanX(tx *gorm.DB, userID uint, x uint) (*model.GameLeaderboard, error)
	UpdateScore(tx *gorm.DB, id uint, gl *model.GameLeaderboard) (*model.GameLeaderboard, error)
	FindByUserID(tx *gorm.DB, userID uint) (*model.GameLeaderboard, error)
	FindAll(tx *gorm.DB) ([]*model.GameLeaderboard, error)
	ResetTriedChance(tx *gorm.DB) error
}

type gameRepository struct {
}

func NewGame() GameRepository {
	return &gameRepository{}
}

func (r *gameRepository) CreateLeaderboard(tx *gorm.DB, gl *model.GameLeaderboard) (*model.GameLeaderboard, error) {
	result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&gl)
	if int(result.RowsAffected) == 0 {
		return nil, new(apperror.GameLeaderboardAlreadyExistError)
	}
	return gl, result.Error
}

func (r *gameRepository) IsTargetScoreReached(tx *gorm.DB, prevScore, score uint) (*model.Game, error) {
	var game *model.Game
	cumulatedScore := score + prevScore

	result := tx.Where("target_score > ? AND target_score <= ?", prevScore, cumulatedScore).First(&game)
	if int(result.RowsAffected) == 0 {
		return nil, new(apperror.NotWinGamePrizeError)
	}
	return game, result.Error
}

func (r *gameRepository) UpdateScore(tx *gorm.DB, userID uint, gl *model.GameLeaderboard) (*model.GameLeaderboard, error) {
	var updatedGl *model.GameLeaderboard
	result := tx.Where("user_id = ?", userID).First(&updatedGl).Updates(&gl)
	return updatedGl, result.Error
}

func (r *gameRepository) ResetTriedChance(tx *gorm.DB) error {
	result := tx.Table("game_leaderboards_tab").Where("tried > 0").Updates(map[string]interface{}{"tried": 0})
	return result.Error
}

func (r *gameRepository) FindByUserID(tx *gorm.DB, userID uint) (*model.GameLeaderboard, error) {
	var gameLeaderboard *model.GameLeaderboard

	result := tx.Preload("User").Where("user_id = ?", userID).First(&gameLeaderboard)
	return gameLeaderboard, result.Error
}

func (r *gameRepository) IsUserTriedLessThanX(tx *gorm.DB, userID uint, x uint) (*model.GameLeaderboard, error) {
	var gameLeaderboard *model.GameLeaderboard

	result := tx.Preload("User").Where("user_id = ? AND tried < ?", userID, x).First(&gameLeaderboard)
	return gameLeaderboard, result.Error
}

func (r *gameRepository) FindAll(tx *gorm.DB) ([]*model.GameLeaderboard, error) {
	var gameLeaderboards []*model.GameLeaderboard
	sortBy := "score"
	sort := "desc"
	orderStatement := fmt.Sprintf("%s %s", sortBy, sort)

	var ids []uint
	tx.Table("users_tab").Distinct().Select("id").Where("role = ?", model.AdminRole).Find(&ids)
	result := tx.Preload("User").Where("user_id NOT IN (?)", ids).Order(orderStatement).Limit(5).Find(&gameLeaderboards)
	return gameLeaderboards, result.Error
}
