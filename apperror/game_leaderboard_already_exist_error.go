package apperror

type GameLeaderboardAlreadyExistError struct{}

func (e *GameLeaderboardAlreadyExistError) Error() string {
	return "user already in game leaderboard"
}
