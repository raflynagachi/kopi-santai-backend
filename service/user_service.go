package service

import (
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"gorm.io/gorm"
)

type UserService interface {
	GetProfileDetail(id uint) (*dto.ProfileRes, error)
	UpdateProfile(id uint, req *dto.UserUpdateReq) (*dto.UserRes, error)
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

func userUpdateReqToUser(req *dto.UserUpdateReq) *model.User {
	return &model.User{
		FullName:       req.FullName,
		Phone:          req.Phone,
		Email:          req.Email,
		Address:        req.Address,
		ProfilePicture: req.ProfilePicture,
	}
}

func (s *userService) GetProfileDetail(id uint) (*dto.ProfileRes, error) {
	tx := s.db.Begin()
	user, err := s.userRepository.FindByIDWithCoupons(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.NotFoundError(err.Error())
	}

	profileRes := new(dto.ProfileRes).FromUser(user)
	return profileRes, nil
}

func (s *userService) UpdateProfile(id uint, req *dto.UserUpdateReq) (*dto.UserRes, error) {
	user := userUpdateReqToUser(req)

	tx := s.db.Begin()
	user, err := s.userRepository.Update(tx, id, user)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	userRes := new(dto.UserRes).FromUser(user)
	return userRes, nil
}
