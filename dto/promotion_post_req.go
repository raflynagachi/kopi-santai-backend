package dto

type PromotionPostReq struct {
	CouponID    uint   `json:"couponID" binding:"required,numeric"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       []byte `json:"image" binding:"required"`
	MinSpent    uint   `json:"minSpent" binding:"required,gte=0"`
}
