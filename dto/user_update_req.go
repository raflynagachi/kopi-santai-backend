package dto

type UserUpdateReq struct {
	FullName       string `json:"fullName"`
	Phone          string `json:"phone" binding:"omitempty,e164,min=5,max=15"`
	Address        string `json:"address"`
	Email          string `json:"email" binding:"omitempty,email"`
	ProfilePicture []byte `json:"profilePicture"`
}
