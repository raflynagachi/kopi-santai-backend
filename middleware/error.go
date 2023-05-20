package middleware

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // only call error handler at last chain
	if len(c.Errors) == 0 {
		return
	}

	firstError := c.Errors[0].Err
	fmt.Println("ErrorHandler: ", firstError)

	appErr, isAppErr := firstError.(apperror.AppError)
	if isAppErr {
		c.JSON(appErr.StatusCode, appErr)
		return
	}

	serverErr := apperror.InternalServerError(firstError.Error())
	c.JSON(serverErr.StatusCode, serverErr)
}
