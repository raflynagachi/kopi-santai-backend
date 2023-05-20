package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/service"
)

type DeliveryHandler interface {
	UpdateStatus(c *gin.Context)
}

type deliveryHandler struct {
	deliveryService service.DeliveryService
}

type DeliveryConfig struct {
	DeliveryService service.DeliveryService
}

func NewDelivery(c *DeliveryConfig) DeliveryHandler {
	return &deliveryHandler{deliveryService: c.DeliveryService}
}

func (h *deliveryHandler) UpdateStatus(c *gin.Context) {
	idParam, _ := c.Get("id")

	payload, _ := c.Get("payload")
	var req *dto.DeliveryUpdateStatusReq
	req = payload.(*dto.DeliveryUpdateStatusReq)

	deliveryRes, err := h.deliveryService.UpdateStatus(idParam.(uint), req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(deliveryRes))
}
