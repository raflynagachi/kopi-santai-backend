package middleware

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"github.com/gin-gonic/gin"
)

func AuthorizeUser(c *gin.Context) {
	_, ok := c.Get("user")
	err := apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error())
	if !ok {
		c.AbortWithStatusJSON(err.StatusCode, err)
	}
}
