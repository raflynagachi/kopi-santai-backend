package service_test

import (
	"errors"
	"testing"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/config"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var user = model.User{
	FullName:       "John Doe",
	Phone:          "+6282211223344",
	Email:          "john.doe@mail.com",
	Role:           "USER",
	Address:        "Jl. Nusa Indah, Desa Beligan",
	Password:       "$2a$10$X9iRj9AhH9KpDQHX9sK8N.x37Bfif8Y6ZCglNOuh48MUaJwwk4Pwy",
	ProfilePicture: model.ImagePlaceholder,
	Coupons:        nil,
}

func TestAuthService_Register(t *testing.T) {
	t.Run("should return token when register success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		gameMockRepository := new(mocks.GameRepository)
		authConfig := service.AuthConfig{
			DB:             gormDB,
			UserRepository: mockRepository,
			GameRepository: gameMockRepository,
			AppConfig:      config.Config,
		}
		s := service.NewAuth(&authConfig)
		userCreated := user
		userCreated.ID = 1
		registerReq := &dto.RegisterPostReq{
			FullName: "John Doe",
			Phone:    "+6282211223344",
			Address:  "Jl. Nusa Indah, Desa Beligan",
			Email:    "john.doe@mail.com",
			Password: user.Password,
		}
		gl := &model.GameLeaderboard{UserID: userCreated.ID}
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &user).Return(&userCreated, nil)
		gameMockRepository.On("CreateLeaderboard", mock.AnythingOfType(testutils.GormDBPointerType), gl).Return(nil, nil)

		tokenRes, err := s.Register(registerReq)

		assert.Nil(t, err)
		assert.NotNil(t, tokenRes)
	})

	t.Run("should return error when register failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		authConfig := service.AuthConfig{DB: gormDB, UserRepository: mockRepository, AppConfig: config.Config}
		s := service.NewAuth(&authConfig)
		registerReq := dto.RegisterPostReq{
			FullName: "John Doe",
			Phone:    "+6282211223344",
			Address:  "Jl. Nusa Indah, Desa Beligan",
			Email:    "john.doe@mail.com",
			Password: user.Password,
		}
		dbErr := new(apperror.EmailAlreadyExistError)
		mockRepository.On(
			"Create",
			mock.AnythingOfType(testutils.GormDBPointerType),
			&user,
		).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		tokenRes, err := s.Register(&registerReq)

		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
		assert.Nil(t, tokenRes)
	})

	t.Run("should return error when create game leaderboard failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		gameMockRepository := new(mocks.GameRepository)
		authConfig := service.AuthConfig{
			DB:             gormDB,
			UserRepository: mockRepository,
			GameRepository: gameMockRepository,
			AppConfig:      config.Config,
		}
		s := service.NewAuth(&authConfig)
		userCreated := user
		userCreated.ID = 1
		registerReq := dto.RegisterPostReq{
			FullName: "John Doe",
			Phone:    "+6282211223344",
			Address:  "Jl. Nusa Indah, Desa Beligan",
			Email:    "john.doe@mail.com",
			Password: user.Password,
		}
		dbErr := errors.New("db error")
		gl := &model.GameLeaderboard{UserID: userCreated.ID}
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &user).Return(&userCreated, nil)
		gameMockRepository.On("CreateLeaderboard", mock.AnythingOfType(testutils.GormDBPointerType), gl).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		tokenRes, err := s.Register(&registerReq)

		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
		assert.Nil(t, tokenRes)
	})
}

func TestAuthService_Login(t *testing.T) {
	t.Run("should get token when login success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		authConfig := service.AuthConfig{DB: gormDB, UserRepository: mockRepository, AppConfig: config.Config}
		s := service.NewAuth(&authConfig)
		loginReq := dto.LoginPostReq{
			Email:    "john.doe@mail.com",
			Password: "password",
		}
		createdUser := user
		createdUser.ID = 1
		mockRepository.On(
			"FindByEmail",
			mock.AnythingOfType(testutils.GormDBPointerType),
			loginReq.Email,
		).Return(&createdUser, nil)

		tokenRes, err := s.Login(&loginReq)

		assert.Nil(t, err)
		assert.NotNil(t, tokenRes)
	})

	t.Run("should return NotFoundError when email doesn't exist", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		authConfig := service.AuthConfig{DB: gormDB, UserRepository: mockRepository, AppConfig: config.Config}
		s := service.NewAuth(&authConfig)
		loginReq := dto.LoginPostReq{
			Email:    "johnfake@mail.com",
			Password: "password",
		}
		dbErr := new(apperror.EmailNotFoundError)
		mockRepository.On(
			"FindByEmail",
			mock.AnythingOfType(testutils.GormDBPointerType),
			loginReq.Email,
		).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.Login(&loginReq)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return UnauthorizedError when email and password doesn't match", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.UserRepository)
		authConfig := service.AuthConfig{DB: gormDB, UserRepository: mockRepository, AppConfig: config.Config}
		s := service.NewAuth(&authConfig)
		loginReq := dto.LoginPostReq{
			Email:    "johnfake@mail.com",
			Password: "passwordWrong",
		}
		dbErr := new(apperror.PasswordError)
		mockRepository.On(
			"FindByEmail",
			mock.AnythingOfType(testutils.GormDBPointerType),
			loginReq.Email,
		).Return(&user, nil)
		expectedErr := apperror.UnauthorizedError(dbErr.Error())

		_, err := s.Login(&loginReq)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
