package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler interface {
	CreateOrderItem(c *gin.Context)
	FindOrderItemByUserID(c *gin.Context)
	UpdateOrderItemByID(c *gin.Context)
}

type orderHandler struct {
	orderService service.OrderService
}

type OrderConfig struct {
	OrderService service.OrderService
}

func NewOrder(c *OrderConfig) OrderHandler {
	return &orderHandler{orderService: c.OrderService}
}

func (h *orderHandler) CreateOrderItem(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	payload, _ := c.Get("payload")
	var req *dto.OrderItemPostReq
	req = payload.(*dto.OrderItemPostReq)

	orderItemRes, err := h.orderService.CreateOrderItem(req, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderHandler) FindOrderItemByUserID(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	orderItemRes, err := h.orderService.FindOrderItemByUserID(userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderHandler) UpdateOrderItemByID(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	payload, _ := c.Get("payload")
	var req *dto.OrderItemPatchReq
	req = payload.(*dto.OrderItemPatchReq)

	orderItemRes, err := h.orderService.UpdateOrderItemByID(idParam.(uint), userID, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
