package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"gorm.io/gorm"
)

type UserService interface {
	GetProfileDetail(id uint) (*dto.ProfileRes, error)
}

type userService struct {
	db             *gorm.DB
	userRepository repository.UserRepository
}

type UserConfig struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

func NewUser(c *UserConfig) UserService {
	return &userService{
		db:             c.DB,
		userRepository: c.UserRepository,
	}
}

func (s *userService) GetProfileDetail(id uint) (*dto.ProfileRes, error) {
	tx := s.db.Begin()
	user, err := s.userRepository.FindByID(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, httperror.NotFoundError(err.Error())
	}

	profileRes := new(dto.ProfileRes).FromUser(user)
	return profileRes, nil
}
