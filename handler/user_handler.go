package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/service"
)

type UserHandler interface {
	GetProfileDetail(c *gin.Context)
	UpdateProfile(c *gin.Context)
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
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}

	userRes, err := h.userService.GetProfileDetail(user.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userRes))
}

func (h *userHandler) UpdateProfile(c *gin.Context) {
	idParam, _ := c.Get("id")
	payloadJwt, ok := c.Get("user")
	user := payloadJwt.(*dto.UserJWT)
	if !ok || idParam != user.ID {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}

	payload, _ := c.Get("payload")
	var req *dto.UserUpdateReq
	req = payload.(*dto.UserUpdateReq)

	userRes, err := h.userService.UpdateProfile(user.ID, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userRes))
}
