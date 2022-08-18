package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type ProfileRes struct {
	FullName       string `json:"fullName"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profilePicture"`
}

func (_ *ProfileRes) FromUser(u *model.User) *ProfileRes {
	return &ProfileRes{
		FullName:       u.FullName,
		Phone:          u.Phone,
		Email:          u.Email,
		Username:       u.Username,
		Address:        u.Address,
		ProfilePicture: u.ProfilePicture,
	}
}
