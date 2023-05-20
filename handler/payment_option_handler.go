package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/service"
)

type PaymentOptionHandler interface {
	FindAll(c *gin.Context)
}

type paymentOptionHandler struct {
	paymentOptService service.PaymentOptionService
}

type PaymentOptConfig struct {
	PaymentOptService service.PaymentOptionService
}

func NewPaymentOpt(c *PaymentOptConfig) PaymentOptionHandler {
	return &paymentOptionHandler{paymentOptService: c.PaymentOptService}
}

func (h *paymentOptionHandler) FindAll(c *gin.Context) {
	userCouponRes, err := h.paymentOptService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userCouponRes))
}
