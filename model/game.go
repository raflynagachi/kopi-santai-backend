package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	CouponID    uint
	TargetScore uint
}

func (_ *Game) TableName() string {
	return "games_tab"
}
