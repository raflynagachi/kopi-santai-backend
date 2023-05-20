package apperror

type NotWinGamePrizeError struct {
}

func (e *NotWinGamePrizeError) Error() string {
	return "score updated - with no prize"
}
