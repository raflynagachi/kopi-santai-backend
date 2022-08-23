package model

import "gorm.io/gorm"

type GameLeaderboard struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Score  uint
	User   User
}

func (_ *GameLeaderboard) TableName() string {
	return "game_leaderboards_tab"
}
