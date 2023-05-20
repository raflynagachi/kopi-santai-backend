package apperror

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "ID doesn't match any user"
}
