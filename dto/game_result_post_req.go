package dto

type GameResultPostReq struct {
	Score uint `json:"score" binding:"required,numeric"`
}
