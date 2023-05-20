package dto

type MenuUpdateReq struct {
	CategoryID uint    `json:"categoryID" binding:"omitempty,numeric"`
	Name       string  `json:"name"`
	Price      float64 `json:"price" binding:"omitempty,numeric,gte=0"`
	Image      []byte  `json:"image"`
}
