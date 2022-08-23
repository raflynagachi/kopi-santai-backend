package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type GameLeaderboardRes struct {
	Name  string `json:"name"`
	Score uint   `json:"score"`
}

func (_ *GameLeaderboardRes) FromGameLeaderboard(leaderboard *model.GameLeaderboard) *GameLeaderboardRes {
	return &GameLeaderboardRes{
		Name:  leaderboard.User.FullName,
		Score: leaderboard.Score,
	}
}
