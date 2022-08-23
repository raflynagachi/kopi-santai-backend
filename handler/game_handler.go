package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GameHandler interface {
	FindByUserID(c *gin.Context)
	FindAll(c *gin.Context)
}

type gameHandler struct {
	gameService service.GameService
}

type GameConfig struct {
	GameService service.GameService
}

func NewGame(c *GameConfig) GameHandler {
	return &gameHandler{gameService: c.GameService}
}

func (h *gameHandler) FindByUserID(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	userID := userPayload.(*dto.UserJWT).ID
	if !ok || idParam != userID {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}

	glRes, err := h.gameService.FindByUserID(userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(glRes))
}

func (h *gameHandler) FindAll(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}

	glRes, err := h.gameService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(glRes))
}
