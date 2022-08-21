package server

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/handler"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/middleware"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterConfig struct {
	AuthService      service.AuthService
	UserService      service.UserService
	MenuService      service.MenuService
	OrderItemService service.OrderItemService
}

const apiNotFoundMessage = "API not found"

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, apperror.NotFoundError(apiNotFoundMessage))
}

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	r.NoRoute(NoRouteHandler)

	authHandler := handler.NewAuth(&handler.AuthConfig{AuthService: c.AuthService})
	userHandler := handler.NewUser(&handler.UserConfig{UserService: c.UserService})
	menuHandler := handler.NewMenu(&handler.MenuConfig{MenuService: c.MenuService})
	orderHandler := handler.NewOrderItem(&handler.OrderItemConfig{OrderService: c.OrderItemService})

	r.POST("/login", middleware.RequestValidator(&dto.LoginPostReq{}), authHandler.Login)
	r.POST("/register", middleware.RequestValidator(&dto.RegisterPostReq{}), authHandler.Register)

	r.GET("/menus", menuHandler.FindAll)
	r.GET("/menus/:id", middleware.ParamIDValidator, menuHandler.GetMenuDetail)

	r.Use(middleware.AuthorizeJWT)
	r.GET("/users/:id", middleware.ParamIDValidator, userHandler.GetProfileDetail)
	r.PATCH("/users/:id", middleware.ParamIDValidator, middleware.RequestValidator(&dto.UserUpdateReq{}), userHandler.UpdateProfile)

	r.POST("/order-items", middleware.RequestValidator(&dto.OrderItemPostReq{}), orderHandler.CreateOrderItem)
	r.GET("/order-items", orderHandler.FindOrderItemByUserID)
	r.PATCH("/order-items/:id", middleware.ParamIDValidator, middleware.RequestValidator(&dto.OrderItemPatchReq{}), orderHandler.UpdateOrderItemByID)
	r.DELETE("/order-items/:id", middleware.ParamIDValidator, orderHandler.DeleteOrderItemByID)

	return r
}
