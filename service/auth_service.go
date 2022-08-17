package service

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthService interface {
	Login(req *dto.LoginPostReq) (*dto.TokenRes, error)
}

type authService struct {
	db             *gorm.DB
	userRepository repository.UserRepository
	appConfig      config.AppConfig
}

type AuthConfig struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
	AppConfig      config.AppConfig
}

func NewAuth(c *AuthConfig) AuthService {
	return &authService{
		db:             c.DB,
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *dto.UserJWT `json:"user"`
}

func (s *authService) generateJWTToken(user *dto.UserJWT) (*dto.TokenRes, error) {
	idExp := s.appConfig.JWTExpiredInMinuteTime * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	claims := &idTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.appConfig.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Unix(tokenExp, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := s.appConfig.JWTSecret
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return &dto.TokenRes{}, err
	}
	return &dto.TokenRes{Token: tokenString}, nil
}

func (s *authService) Login(req *dto.LoginPostReq) (*dto.TokenRes, error) {
	tx := s.db.Begin()
	user, err := s.userRepository.FindByEmail(tx, req.Email)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, httperror.NotFoundError(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, httperror.UnauthorizedError(err.Error())
	}

	userRes := new(dto.UserJWT).FromUser(user)
	token, err := s.generateJWTToken(userRes)
	return token, err
}
