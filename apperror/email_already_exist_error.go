package apperror

type EmailAlreadyExistError struct{}

func (e *EmailAlreadyExistError) Error() string {
	return "email already registered"
}
