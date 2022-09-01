package middleware

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"github.com/gin-gonic/gin"
	"reflect"
)

func RequestValidator(model any) gin.HandlerFunc {
	return func(c *gin.Context) {
		modelPtr := reflect.ValueOf(model).Elem()
		modelPtr.Set(reflect.Zero(modelPtr.Type()))
		c.Set("payload", new(any))
		if err := c.ShouldBindJSON(model); err != nil {
			appErr := apperror.BadRequestError(err.Error())
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			return
		}
		c.Set("payload", model)
	}
}
