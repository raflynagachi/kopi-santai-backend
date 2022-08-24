package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var req *dto.DeliveryUpdateStatusReq
	req = payload.(*dto.DeliveryUpdateStatusReq)

	orderItemRes, err := h.deliveryService.UpdateStatus(idParam.(uint), req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}