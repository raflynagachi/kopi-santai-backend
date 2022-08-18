package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type UserRes struct {
	FullName       string `json:"fullName"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profilePicture"`
}

func (r *UserRes) FromUser(u *model.User) *UserRes {
	return &UserRes{
		FullName:       u.FullName,
		Phone:          u.Phone,
		Email:          u.Email,
		Username:       u.Username,
		Address:        u.Address,
		ProfilePicture: u.ProfilePicture,
	}
}
