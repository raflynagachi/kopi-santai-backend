package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

type UserRes struct {
	FullName       string `json:"fullName"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	ProfilePicture []byte `json:"profilePicture"`
}

func (r *UserRes) FromUser(u *model.User) *UserRes {
	return &UserRes{
		FullName:       u.FullName,
		Phone:          u.Phone,
		Email:          u.Email,
		Address:        u.Address,
		ProfilePicture: u.ProfilePicture,
	}
}
