package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
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
