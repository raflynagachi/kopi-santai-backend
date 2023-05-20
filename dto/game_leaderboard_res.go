package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

type GameLeaderboardRes struct {
	UserID uint   `json:"userID"`
	Name   string `json:"name"`
	Score  uint   `json:"score"`
	Tried  uint   `json:"tried"`
}

func (_ *GameLeaderboardRes) FromGameLeaderboard(leaderboard *model.GameLeaderboard) *GameLeaderboardRes {
	return &GameLeaderboardRes{
		UserID: leaderboard.UserID,
		Name:   leaderboard.User.FullName,
		Score:  leaderboard.Score,
		Tried:  leaderboard.Tried,
	}
}
