package middleware

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
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
