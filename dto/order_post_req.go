package dto

import "time"

type OrderPostReq struct {
	PaymentOptID uint      `json:"paymentOptID" binding:"required"`
	CouponID     uint      `json:"couponID"`
	OrderedDate  time.Time `json:"orderedDate" binding:"required"`
}
