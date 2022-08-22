package dto

import "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"

type ReviewRes struct {
	MenuID      uint    `json:"menuID"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

func (_ *ReviewRes) FromReview(r *model.Review) *ReviewRes {
	return &ReviewRes{
		MenuID:      r.MenuID,
		Description: r.Description,
		Rating:      r.Rating,
	}
}
