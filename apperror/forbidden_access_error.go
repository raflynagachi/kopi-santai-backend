package apperror

type ForbiddenAccessError struct {
}

func (_ *ForbiddenAccessError) Error() string {
	return "forbidden access"
}
