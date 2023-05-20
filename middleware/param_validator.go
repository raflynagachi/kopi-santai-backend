package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
)

func ParamIDValidator(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appErr := apperror.BadRequestError("ID has wrong format")
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		return
	}
	c.Set("id", uint(id))
}
