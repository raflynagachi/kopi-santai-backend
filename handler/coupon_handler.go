package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CouponHandler interface {
	Create(c *gin.Context)
	FindCouponByUserID(c *gin.Context)
	FindAll(c *gin.Context)
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

func (h *couponHandler) FindCouponByUserID(c *gin.Context) {
	userPayload, _ := c.Get("user")
	userID := userPayload.(*dto.UserJWT).ID

	userCouponRes, err := h.couponService.FindCouponByUserID(userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userCouponRes))
}

func (h *couponHandler) FindAll(c *gin.Context) {
	userCouponRes, err := h.couponService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userCouponRes))
}

func (h *couponHandler) DeleteByID(c *gin.Context) {
	idParam, _ := c.Get("id")

	orderItemRes, err := h.couponService.DeleteByID(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
