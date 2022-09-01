package apperror

type EmailNotFoundError struct{}

func (e *EmailNotFoundError) Error() string {
	return "email doesn't match any user"
}
