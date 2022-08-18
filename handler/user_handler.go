package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/customerror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	GetProfileDetail(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

type UserConfig struct {
	UserService service.UserService
}

func NewUser(c *UserConfig) UserHandler {
	return &userHandler{userService: c.UserService}
}

func (h *userHandler) GetProfileDetail(c *gin.Context) {
	idParam, _ := c.Get("id")
	payload, ok := c.Get("user")
	user := payload.(*dto.UserJWT)
	if !ok || idParam != user.ID {
		_ = c.Error(httperror.UnauthorizedError(new(customerror.UserUnauthorizedError).Error()))
		return
	}

	userRes, err := h.userService.GetProfileDetail(user.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userRes))
}
