package service

import (
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
	FindAll() ([]*dto.CouponRes, error)
	FindCouponByUserID(userID uint) ([]*dto.CouponRes, error)
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
		Name:   req.Name,
		Amount: req.Amount,
	}

	tx := s.db.Begin()
	coupon, err := s.couponRepo.Create(tx, c)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	couponRes := new(dto.CouponRes).FromCoupon(coupon)
	return couponRes, nil
}

func (s *couponService) FindCouponByUserID(userID uint) ([]*dto.CouponRes, error) {
	tx := s.db.Begin()
	promos, err := s.couponRepo.FindCouponByUserID(tx, userID)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var couponRes []*dto.CouponRes
	for _, promo := range promos {
		if promo.Coupon != nil {
			couponRes = append(couponRes, new(dto.CouponRes).FromCoupon(promo.Coupon))
		}
	}

	return couponRes, nil
}

func (s *couponService) FindAll() ([]*dto.CouponRes, error) {
	tx := s.db.Begin()
	coupons, err := s.couponRepo.FindAll(tx)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return nil, apperror.InternalServerError(err.Error())
	}

	var couponRes []*dto.CouponRes
	for _, coupon := range coupons {
		if coupon != nil {
			couponRes = append(couponRes, new(dto.CouponRes).FromCoupon(coupon))
		}
	}

	return couponRes, nil
}

func (s *couponService) DeleteByID(id uint) (gin.H, error) {
	tx := s.db.Begin()
	isDeleted, err := s.couponRepo.DeleteByID(tx, id)
	if err != nil {
		tx.Rollback()
		return gin.H{"isDeleted": false}, apperror.BadRequestError(err.Error())
	}
	_, err = s.couponRepo.DeleteUserCoupon(tx, id)
	helper.CommitOrRollback(tx, err)
	if err != nil {
		return gin.H{"isDeleted": false}, apperror.InternalServerError(err.Error())
	}

	return gin.H{"isDeleted": isDeleted}, nil
}
