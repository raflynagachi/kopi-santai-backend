package middleware

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"github.com/gin-gonic/gin"
)

func AuthorizeAdmin(c *gin.Context) {
	userPayload, _ := c.Get("user")
	role := userPayload.(*dto.UserJWT).Role

	err := apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error())
	if role != model.AdminRole {
		c.AbortWithStatusJSON(err.StatusCode, err)
	}
}
