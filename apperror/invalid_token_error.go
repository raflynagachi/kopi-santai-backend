package apperror

type InvalidTokenError struct{}

func (e *InvalidTokenError) Error() string {
	return "invalid authentication token"
}
