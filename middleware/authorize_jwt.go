package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/config"
	"github.com/raflynagachi/kopi-santai-backend/dto"
)

func validateToken(encoded string) (*jwt.Token, error) {
	return jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperror.UnauthorizedError(new(apperror.InvalidTokenError).Error())
		}
		return config.Config.JWTSecret, nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	if config.Config.ENV == "testing" {
		userJwt := &dto.UserJWT{
			ID:    1,
			Email: "john.doe@mail,com",
			Role:  "ADMIN",
		}
		c.Set("user", userJwt)
		fmt.Println("disable JWT authentication on dev env")
		return
	}
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	unauthorizedErr := apperror.UnauthorizedError(new(apperror.InvalidTokenError).Error())
	if len(s) < 2 {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	encodedToken := s[1]
	token, err := validateToken(encodedToken)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	fmt.Println("claims: ", claims)

	userJson, err := json.Marshal(claims["user"])
	var user *dto.UserJWT
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	c.Set("user", user)
}
