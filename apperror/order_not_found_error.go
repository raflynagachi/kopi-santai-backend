package apperror

type OrderNotFoundError struct {
}

func (e *OrderNotFoundError) Error() string {
	return "order not found"
}
