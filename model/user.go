package model

import "gorm.io/gorm"

var UserRole = "USER"
var ImagePlaceholder = "https://via.placeholder.com/150"

type User struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	FullName       string
	Phone          string
	Email          string
	Username       string
	Role           string
	Address        string
	Password       string
	ProfilePicture string
}

func (u *User) TableName() string {
	return "users_tab"
}
