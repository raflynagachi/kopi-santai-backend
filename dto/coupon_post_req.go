package dto

type CouponPostReq struct {
	Name   string  `json:"name" binding:"required"`
	Amount float64 `json:"amount" binding:"required,numeric,gte=0,lte=100"`
}
