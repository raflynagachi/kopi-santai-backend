package middleware

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/httperror"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParamIDValidator(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appErr := httperror.BadRequestError("ID has wrong format")
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		return
	}
	c.Set("id", uint(id))
}
