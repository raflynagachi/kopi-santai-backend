package apperror

type OrderItemsEmptyError struct {
}

func (e *OrderItemsEmptyError) Error() string {
	return "order items is empty"
}
