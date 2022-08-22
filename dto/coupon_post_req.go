package dto

import "time"

type CouponPostReq struct {
	Name        string    `json:"name" binding:"required"`
	Amount      float64   `json:"amount" binding:"required,numeric,gte=0"`
	MinSpent    float64   `json:"minSpent" binding:"required,numeric,gte=0"`
	ExpiredDate time.Time `json:"expiredDate" binding:"required"`
}
