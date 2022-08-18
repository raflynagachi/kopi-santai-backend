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
	FindByID(tx *gorm.DB, id uint) (*model.User, error)
	FindByIDWithCoupons(tx *gorm.DB, id uint) (*model.User, error)
	Update(tx *gorm.DB, id uint, user *model.User) (*model.User, error)
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

func (_ *userRepository) FindByIDWithCoupons(tx *gorm.DB, id uint) (*model.User, error) {
	var user *model.User
	err := tx.Preload("Coupons").First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, new(customerror.EmailNotFoundError)
	}
	return user, err
}

func (_ *userRepository) FindByID(tx *gorm.DB, id uint) (*model.User, error) {
	var user *model.User
	err := tx.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, new(customerror.EmailNotFoundError)
	}
	return user, err
}

func (_ *userRepository) Update(tx *gorm.DB, id uint, user *model.User) (*model.User, error) {
	var updatedUser *model.User
	err := tx.First(&updatedUser, id).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
