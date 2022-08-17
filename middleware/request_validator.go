package middleware

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/httperror"
	"github.com/gin-gonic/gin"
)

func RequestValidator(model any) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			appErr := httperror.BadRequestError(err.Error())
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		c.Set("payload", model)
	}
}
