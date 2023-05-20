package apperror

type ReviewCreatedError struct{}

func (e *ReviewCreatedError) Error() string {
	return "review has been created before"
}
