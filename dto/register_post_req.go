package dto

type RegisterPostReq struct {
	FullName string `json:"fullName" binding:"required"`
	Phone    string `json:"phone" binding:"required,e164,min=5,max=15"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
