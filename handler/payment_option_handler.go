package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
