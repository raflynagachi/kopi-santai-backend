package model

import "gorm.io/gorm"

var UserRole = "USER"
var AdminRole = "ADMIN"
var ImagePlaceholder = "https://via.placeholder.com/150"

type User struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	FullName       string
	Phone          string
	Email          string
	Role           string
	Address        string
	Password       string
	ProfilePicture string
	Coupons        []*Coupon `gorm:"many2many:users_coupons_tab;"`
}

func (u *User) TableName() string {
	return "users_tab"
}
