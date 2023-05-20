package dto

type DeliveryUpdateStatusReq struct {
	Status string `json:"status" binding:"required"`
}
