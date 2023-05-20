package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
)

func AuthorizeUser(c *gin.Context) {
	_, ok := c.Get("user")
	err := apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error())
	if !ok {
		c.AbortWithStatusJSON(err.StatusCode, err)
	}
}
