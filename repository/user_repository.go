package repository

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dberror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(tx *gorm.DB, email string) (*model.User, error)
}

type userRepository struct{}

func NewUser() UserRepository {
	return &userRepository{}
}

func (_ *userRepository) FindByEmail(tx *gorm.DB, email string) (*model.User, error) {
	var user *model.User
	err := tx.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, new(dberror.EmailNotFoundError)
	}
	return user, err
}
