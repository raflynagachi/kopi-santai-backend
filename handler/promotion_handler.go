package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PromotionHandler interface {
	FindAll(c *gin.Context)
}

type promotionHandler struct {
	promoService service.PromotionService
}

type PromoConfig struct {
	PromoService service.PromotionService
}

func NewPromo(c *PromoConfig) PromotionHandler {
	return &promotionHandler{promoService: c.PromoService}
}

func (h *promotionHandler) FindAll(c *gin.Context) {
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

	promoSRes, err := h.promoService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(promoSRes))
}
