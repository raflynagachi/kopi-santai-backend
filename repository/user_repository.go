package repository

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/customerror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(tx *gorm.DB, user *model.User) (*model.User, error)
	FindByEmail(tx *gorm.DB, email string) (*model.User, error)
}

type userRepository struct{}

func NewUser() UserRepository {
	return &userRepository{}
}

func (_ *userRepository) Create(tx *gorm.DB, user *model.User) (*model.User, error) {
	result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(user)
	if int(result.RowsAffected) == 0 {
		return nil, new(customerror.EmailAlreadyExistError)
	}
	return user, result.Error
}

func (_ *userRepository) FindByEmail(tx *gorm.DB, email string) (*model.User, error) {
	var user *model.User
	err := tx.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, new(customerror.EmailNotFoundError)
	}
	return user, err
}
