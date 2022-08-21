package apperror

type CouponFailedError struct {
}

func (e *CouponFailedError) Error() string {
	return "coupon failed - requirement doesn't match"
}
