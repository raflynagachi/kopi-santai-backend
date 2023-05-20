package apperror

type PasswordError struct {
}

func (e *PasswordError) Error() string {
	return "email and password doesn't match"
}
