package apperror

type MenuNotFoundError struct{}

func (e *MenuNotFoundError) Error() string {
	return "menu doesn't exist"
}
