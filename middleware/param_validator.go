package middleware

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"github.com/gin-gonic/gin"
	"strconv"
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
