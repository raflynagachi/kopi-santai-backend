package server

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/handler"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/httperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/middleware"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterConfig struct {
	AuthService service.AuthService
	UserService service.UserService
}

const apiNotFoundMessage = "API not found"

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, httperror.NotFoundError(apiNotFoundMessage))
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	r.NoRoute(NoRouteHandler)

	authHandler := handler.NewAuth(&handler.AuthConfig{AuthService: c.AuthService})
	userHandler := handler.NewUser(&handler.UserConfig{UserService: c.UserService})

	r.POST("/login", middleware.RequestValidator(&dto.LoginPostReq{}), authHandler.Login)
	r.POST("/register", middleware.RequestValidator(&dto.RegisterPostReq{}), authHandler.Register)

	r.Use(middleware.AuthorizeJWT)
	r.GET("/users/:id", middleware.ParamIDValidator, userHandler.GetProfileDetail)
	return r
}
