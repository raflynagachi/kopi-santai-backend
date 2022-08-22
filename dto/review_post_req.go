package dto

type ReviewPostReq struct {
	MenuID      uint    `json:"menuID" binding:"required,numeric"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating" binding:"required,numeric,gte=0,lte=5"`
}
