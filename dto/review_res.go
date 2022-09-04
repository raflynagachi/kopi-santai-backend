package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type ReviewRes struct {
	UserID      uint    `json:"userID"`
	MenuID      uint    `json:"menuID"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	UserEmail   string  `json:"userEmail"`
}

func (_ *ReviewRes) FromReview(r *model.Review) *ReviewRes {
	return &ReviewRes{
		UserID:      r.UserID,
		MenuID:      r.MenuID,
		Description: r.Description,
		Rating:      r.Rating,
		UserEmail:   r.User.Email,
	}
}
