package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CouponHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type couponHandler struct {
	couponService service.CouponService
}

type CouponConfig struct {
	CouponService service.CouponService
}

func NewCoupon(c *CouponConfig) CouponHandler {
	return &couponHandler{couponService: c.CouponService}
}

func (h *couponHandler) Create(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	role := userPayload.(*dto.UserJWT).Role
	if role != model.AdminRole {
		_ = c.Error(apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error()))
		return
	}

	payload, _ := c.Get("payload")
	var req *dto.CouponPostReq
	req = payload.(*dto.CouponPostReq)

	coupon, err := h.couponService.Create(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(coupon))
}

func (h *couponHandler) DeleteByID(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	role := userPayload.(*dto.UserJWT).Role
	if role != model.AdminRole {
		_ = c.Error(apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error()))
		return
	}

	orderItemRes, err := h.couponService.DeleteByID(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
