package apperror

type UserUnauthorizedError struct{}

func (e *UserUnauthorizedError) Error() string {
	return "unauthorized user"
}
