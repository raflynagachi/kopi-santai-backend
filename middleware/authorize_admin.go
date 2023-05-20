package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/model"
)

func AuthorizeAdmin(c *gin.Context) {
	userPayload, _ := c.Get("user")
	role := userPayload.(*dto.UserJWT).Role

	err := apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error())
	if role != model.AdminRole {
		c.AbortWithStatusJSON(err.StatusCode, err)
	}
}
