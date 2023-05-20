package apperror

type ReviewNotFoundError struct {
}

func (e *ReviewNotFoundError) Error() string {
	return "review not found"
}
