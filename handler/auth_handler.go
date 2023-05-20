package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/service"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authHandler struct {
	authService service.AuthService
}

type AuthConfig struct {
	AuthService service.AuthService
}

func NewAuth(c *AuthConfig) AuthHandler {
	return &authHandler{authService: c.AuthService}
}

func (h *authHandler) Login(c *gin.Context) {
	payload, _ := c.Get("payload")
	var req *dto.LoginPostReq
	req = payload.(*dto.LoginPostReq)

	tokenRes, err := h.authService.Login(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(tokenRes))
}

func (h *authHandler) Register(c *gin.Context) {
	payload, _ := c.Get("payload")
	var req *dto.RegisterPostReq
	req = payload.(*dto.RegisterPostReq)

	tokenRes, err := h.authService.Register(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(tokenRes))
}
