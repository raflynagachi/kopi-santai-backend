package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PromotionHandler interface {
	FindAll(c *gin.Context)
	FindAllUnscoped(c *gin.Context)
	CreatePromotion(c *gin.Context)
	DeletePromotionByID(c *gin.Context)
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
	promoSRes, err := h.promoService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(promoSRes))
}

func (h *promotionHandler) FindAllUnscoped(c *gin.Context) {
	promoSRes, err := h.promoService.FindAllUnscoped()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(promoSRes))
}

func (h *promotionHandler) CreatePromotion(c *gin.Context) {
	payload, _ := c.Get("payload")
	var req *dto.PromotionPostReq
	req = payload.(*dto.PromotionPostReq)

	promoSRes, err := h.promoService.CreatePromotion(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(promoSRes))
}

func (h *promotionHandler) DeletePromotionByID(c *gin.Context) {
	idParam, _ := c.Get("id")

	promoSRes, err := h.promoService.DeletePromotionByID(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(promoSRes))
}
