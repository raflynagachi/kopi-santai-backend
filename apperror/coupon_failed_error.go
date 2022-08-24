package apperror

type CouponNotFoundError struct {
}

func (e *CouponNotFoundError) Error() string {
	return "coupon not found"
}
