package dto

type MenuPostReq struct {
	CategoryID uint    `json:"categoryID" binding:"required,numeric"`
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required,numeric,gte=0"`
	Image      []byte  `json:"image" binding:"required"`
}
