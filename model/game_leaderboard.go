package model

import "gorm.io/gorm"

const MaxTried = uint(5)

type GameLeaderboard struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Score  uint
	Tried  uint
	User   *User
}

func (_ *GameLeaderboard) TableName() string {
	return "game_leaderboards_tab"
}
