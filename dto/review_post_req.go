package dto

type ReviewPostReq struct {
	MenuID      uint    `json:"menuID" binding:"required"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating" binding:"required,gte=1,lte=5"`
}
