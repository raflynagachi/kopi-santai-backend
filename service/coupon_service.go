package service

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CouponService interface {
	Create(req *dto.CouponPostReq) (*dto.CouponRes, error)
	DeleteByID(id uint) (gin.H, error)
}

type couponService struct {
	db         *gorm.DB
	couponRepo repository.CouponRepository
}

type CouponConfig struct {
	DB         *gorm.DB
	CouponRepo repository.CouponRepository
}

func NewCoupon(c *CouponConfig) CouponService {
	return &couponService{
		db:         c.DB,
		couponRepo: c.CouponRepo,
	}
}

func (s *couponService) Create(req *dto.CouponPostReq) (*dto.CouponRes, error) {
	c := &model.Coupon{
		Name:        req.Name,
		Amount:      req.Amount,
		IsAvailable: true,
		MinSpent:    req.MinSpent,
		ExpiredDate: req.ExpiredDate,
	}

	tx := s.db.Begin()
	coupon, err := s.couponRepo.Create(tx, c)
	fmt.Printf("MANTAP: %+v", coupon)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	couponRes := new(dto.CouponRes).FromCoupon(coupon)
	return couponRes, nil
}

func (s *couponService) DeleteByID(id uint) (gin.H, error) {
	tx := s.db.Begin()
	isDeleted, err := s.couponRepo.DeleteByID(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return gin.H{"isDeleted": false}, apperror.BadRequestError(err.Error())
	}

	return gin.H{"isDeleted": isDeleted}, nil
}
