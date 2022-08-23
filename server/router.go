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
	OrderService     service.OrderService
	ReviewService    service.ReviewService
	DeliveryService  service.DeliveryService
	CouponService    service.CouponService
	GameService      service.GameService
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
	orderItemHandler := handler.NewOrderItem(&handler.OrderItemConfig{OrderService: c.OrderItemService})
	orderHandler := handler.NewOrder(&handler.OrderConfig{OrderService: c.OrderService})
	reviewHandler := handler.NewReview(&handler.ReviewConfig{ReviewService: c.ReviewService})
	deliveryHandler := handler.NewDelivery(&handler.DeliveryConfig{DeliveryService: c.DeliveryService})
	couponHandler := handler.NewCoupon(&handler.CouponConfig{CouponService: c.CouponService})
	gameHandler := handler.NewGame(&handler.GameConfig{GameService: c.GameService})

	r.POST("/login", middleware.RequestValidator(&dto.LoginPostReq{}), authHandler.Login)
	r.POST("/register", middleware.RequestValidator(&dto.RegisterPostReq{}), authHandler.Register)

	r.GET("/menus", menuHandler.FindAll)
	r.GET("/menus/:id", middleware.ParamIDValidator, menuHandler.GetMenuDetail)

	r.Use(middleware.AuthorizeJWT)
	r.GET("/users/:id", middleware.ParamIDValidator, userHandler.GetProfileDetail)
	r.PATCH("/users/:id", middleware.ParamIDValidator, middleware.RequestValidator(&dto.UserUpdateReq{}), userHandler.UpdateProfile)

	r.POST("/order-items", middleware.RequestValidator(&dto.OrderItemPostReq{}), orderItemHandler.CreateOrderItem)
	r.GET("/order-items", orderItemHandler.FindOrderItemByUserID)
	r.PATCH("/order-items/:id", middleware.ParamIDValidator, middleware.RequestValidator(&dto.OrderItemPatchReq{}), orderItemHandler.UpdateOrderItemByID)
	r.DELETE("/order-items/:id", middleware.ParamIDValidator, orderItemHandler.DeleteOrderItemByID)

	r.POST("/orders", middleware.RequestValidator(&dto.OrderPostReq{}), orderHandler.CreateOrder)
	r.GET("/orders/:id", middleware.ParamIDValidator, orderHandler.FindOrderByIDAndUserID)

	r.POST("/reviews", middleware.RequestValidator(&dto.ReviewPostReq{}), reviewHandler.Create)

	r.GET("/games", gameHandler.FindAll)
	r.GET("/games/:id", middleware.ParamIDValidator, gameHandler.FindByUserID)

	// ADMIN
	r.GET("/orders", orderHandler.FindAll)
	r.PATCH("/deliveries/:id", middleware.ParamIDValidator, middleware.RequestValidator(&dto.DeliveryUpdateStatusReq{}), deliveryHandler.UpdateStatus)
	r.POST("/menus", middleware.RequestValidator(&dto.MenuPostReq{}), menuHandler.CreateMenu)
	r.PATCH("/menus/:id", middleware.ParamIDValidator, middleware.RequestValidator(&dto.MenuUpdateReq{}), menuHandler.UpdateMenu)
	r.DELETE("/menus/:id", middleware.ParamIDValidator, menuHandler.DeleteByID)
	r.GET("/menus/:id/reviews", middleware.ParamIDValidator, reviewHandler.FindByMenuID)
	r.POST("/coupons", middleware.RequestValidator(&dto.CouponPostReq{}), couponHandler.Create)
	r.DELETE("/coupons/:id", middleware.ParamIDValidator, couponHandler.DeleteByID)
	return r
}
