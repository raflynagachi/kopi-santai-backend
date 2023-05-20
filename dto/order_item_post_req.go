package dto

type OrderItemPostReq struct {
	MenuID      uint   `json:"menuID" binding:"required,numeric"`
	Quantity    int    `json:"quantity" binding:"required,numeric,gte=0"`
	Description string `json:"description"`
}
