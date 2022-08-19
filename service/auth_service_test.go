package service_test

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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
		authConfig := service.AuthConfig{DB: gormDB, UserRepository: mockRepository, AppConfig: config.Config}
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
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), &user).Return(&userCreated, nil)

		tokenRes, err := s.Register(&registerReq)

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
		errMessage := new(apperror.EmailAlreadyExistError).Error()
		mockRepository.On(
			"Create",
			mock.AnythingOfType(testutils.GormDBPointerType),
			&user,
		).Return(nil, apperror.UnprocessableEntityError(errMessage))
		expectedErr := apperror.UnprocessableEntityError(errMessage)

		tokenRes, err := s.Register(&registerReq)

		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, errMessage)
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
		errMessage := new(apperror.EmailNotFoundError).Error()
		mockRepository.On(
			"FindByEmail",
			mock.AnythingOfType(testutils.GormDBPointerType),
			loginReq.Email,
		).Return(nil, apperror.NotFoundError(errMessage))
		expectedErr := apperror.NotFoundError(errMessage)

		_, err := s.Login(&loginReq)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, errMessage)
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
		errMessage := "crypto/bcrypt: hashedPassword is not the hash of the given password"
		mockRepository.On(
			"FindByEmail",
			mock.AnythingOfType(testutils.GormDBPointerType),
			loginReq.Email,
		).Return(&user, nil)
		expectedErr := apperror.UnauthorizedError(errMessage)

		_, err := s.Login(&loginReq)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, errMessage)
	})
}
