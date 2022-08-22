package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReviewHandler interface {
	Create(c *gin.Context)
	FindByMenuID(c *gin.Context)
}

type reviewHandler struct {
	reviewService service.ReviewService
}

type ReviewConfig struct {
	ReviewService service.ReviewService
}

func NewReview(c *ReviewConfig) ReviewHandler {
	return &reviewHandler{reviewService: c.ReviewService}
}

func (h *reviewHandler) Create(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	userID := userPayload.(*dto.UserJWT).ID

	payload, _ := c.Get("payload")
	var req *dto.ReviewPostReq
	req = payload.(*dto.ReviewPostReq)

	reviewRes, err := h.reviewService.Create(req, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(reviewRes))
}

func (h *reviewHandler) FindByMenuID(c *gin.Context) {
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

	reviewRes, err := h.reviewService.FindByMenuID(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(reviewRes))
}
