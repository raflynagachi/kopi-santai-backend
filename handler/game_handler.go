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
	AddCouponPrizeToUser(c *gin.Context)
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
	glRes, err := h.gameService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(glRes))
}

func (h *gameHandler) AddCouponPrizeToUser(c *gin.Context) {
	userPayload, _ := c.Get("user")
	userID := userPayload.(*dto.UserJWT).ID

	payload, _ := c.Get("payload")
	var req *dto.GameResultPostReq
	req = payload.(*dto.GameResultPostReq)

	userCouponRes, err := h.gameService.AddCouponPrizeToUser(req, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(userCouponRes))
}
