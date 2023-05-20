package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
