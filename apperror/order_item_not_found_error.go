package apperror

type OrderItemNotFoundError struct {
}

func (e *OrderItemNotFoundError) Error() string {
	return "order item not found"
}
