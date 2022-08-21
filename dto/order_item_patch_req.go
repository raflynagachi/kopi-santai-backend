package dto

type OrderItemPatchReq struct {
	Quantity    int    `json:"quantity" binding:"required,numeric,gte=0"`
	Description string `json:"description"`
}
