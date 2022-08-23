package repository

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GameRepository interface {
	CreateLeaderboard(tx *gorm.DB, gl *model.GameLeaderboard) (*model.GameLeaderboard, error)
	FindByUserID(tx *gorm.DB, userID uint) (*model.GameLeaderboard, error)
	FindAll(tx *gorm.DB) ([]*model.GameLeaderboard, error)
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

func (r *gameRepository) FindByUserID(tx *gorm.DB, userID uint) (*model.GameLeaderboard, error) {
	var gameLeaderboard *model.GameLeaderboard

	result := tx.Preload("User").Where("user_id = ?", userID).First(&gameLeaderboard)
	return gameLeaderboard, result.Error
}

func (r *gameRepository) FindAll(tx *gorm.DB) ([]*model.GameLeaderboard, error) {
	var gameLeaderboards []*model.GameLeaderboard
	sortBy := "score"
	sort := "desc"
	orderStatement := fmt.Sprintf("%s %s", sortBy, sort)

	result := tx.Preload("User").Order(orderStatement).Limit(10).Find(&gameLeaderboards)
	return gameLeaderboards, result.Error
}
