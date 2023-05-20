package dto

import "github.com/raflynagachi/kopi-santai-backend/model"

type UserJWT struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (_ *UserJWT) FromUser(u *model.User) *UserJWT {
	return &UserJWT{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role,
	}
}
