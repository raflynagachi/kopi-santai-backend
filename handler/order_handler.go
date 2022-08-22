package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler interface {
	CreateOrder(c *gin.Context)
	FindOrderByIDAndUserID(c *gin.Context)
	FindAll(c *gin.Context)
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

func (h *orderHandler) CreateOrder(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	payload, _ := c.Get("payload")
	var req *dto.OrderPostReq
	req = payload.(*dto.OrderPostReq)

	orderItemRes, err := h.orderService.CreateOrder(req, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderHandler) FindOrderByIDAndUserID(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	orderItemRes, err := h.orderService.FindOrderByIDAndUserID(idParam.(uint), userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *orderHandler) FindAll(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	role := userPayload.(*dto.UserJWT).Role
	userID := userPayload.(*dto.UserJWT).ID

	if role != model.AdminRole {
		orderItemRes, err := h.orderService.FindOrderByUserID(userID)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
		return
	}

	queryParam := &model.QueryParamOrder{
		Date: helper.GetQuery(c, "date", ""),
	}

	orderItemRes, err := h.orderService.FindAll(queryParam)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
