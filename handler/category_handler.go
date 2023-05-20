package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/service"
)

type CategoryHandler interface {
	FindAll(c *gin.Context)
}

type categoryHandler struct {
	categoryService service.CategoryService
}

type CategoryConfig struct {
	CategoryService service.CategoryService
}

func NewCategory(c *CategoryConfig) CategoryHandler {
	return &categoryHandler{categoryService: c.CategoryService}
}

func (h *categoryHandler) FindAll(c *gin.Context) {
	userCouponRes, err := h.categoryService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userCouponRes))
}
