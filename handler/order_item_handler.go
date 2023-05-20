package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/service"
)

type OrderItemHandler interface {
	CreateOrderItem(c *gin.Context)
	FindOrderItemByUserID(c *gin.Context)
	UpdateOrderItemByID(c *gin.Context)
	DeleteOrderItemByID(c *gin.Context)
	DeleteOrderItemByUserID(c *gin.Context)
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
	userPayload, _ := c.Get("user")
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
	userPayload, _ := c.Get("user")
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
	userPayload, _ := c.Get("user")
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
	userPayload, _ := c.Get("user")
	userID := userPayload.(*dto.UserJWT).ID

	orderItemRes, err := h.orderItemService.DeleteOrderItemByID(idParam.(uint), userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderItemHandler) DeleteOrderItemByUserID(c *gin.Context) {
	userPayload, _ := c.Get("user")
	userID := userPayload.(*dto.UserJWT).ID

	orderItemRes, err := h.orderItemService.DeleteOrderItemByUserID(userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
