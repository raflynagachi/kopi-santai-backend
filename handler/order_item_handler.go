package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderItemHandler interface {
	CreateOrderItem(c *gin.Context)
	FindOrderItemByUserID(c *gin.Context)
	UpdateOrderItemByID(c *gin.Context)
	DeleteOrderItemByID(c *gin.Context)
}

type orderItemHandler struct {
	orderItemService service.OrderItemService
}

type OrderItemConfig struct {
	OrderService service.OrderItemService
}

func NewOrderItem(c *OrderItemConfig) OrderItemHandler {
	return &orderItemHandler{orderItemService: c.OrderService}
}

func (h *orderItemHandler) CreateOrderItem(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	payload, _ := c.Get("payload")
	var req *dto.OrderItemPostReq
	req = payload.(*dto.OrderItemPostReq)

	orderItemRes, err := h.orderItemService.CreateOrderItem(req, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderItemHandler) FindOrderItemByUserID(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	orderItemRes, err := h.orderItemService.FindOrderItemByUserID(userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderItemHandler) UpdateOrderItemByID(c *gin.Context) {
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

	orderItemRes, err := h.orderItemService.UpdateOrderItemByID(idParam.(uint), userID, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderItemHandler) DeleteOrderItemByID(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	orderItemRes, err := h.orderItemService.DeleteOrderItemByID(idParam.(uint), userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
