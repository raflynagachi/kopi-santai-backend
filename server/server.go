package server

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/config"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/repository"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"log"
)

func Init() {
	userRepo := repository.NewUser()
	menuRepo := repository.NewMenu()
	orderItemRepo := repository.NewOrderItem()
	orderRepo := repository.NewOrder()
	paymentOptRepo := repository.NewPaymentOption()
	couponRepo := repository.NewCoupon()
	deliveryRepo := repository.NewDelivery()
	reviewRepo := repository.NewReview()

	authService := service.NewAuth(&service.AuthConfig{
		DB:             db.Get(),
		UserRepository: userRepo,
		AppConfig:      config.Config,
	})
	userService := service.NewUser(&service.UserConfig{
		DB:             db.Get(),
		UserRepository: userRepo,
	})
	menuService := service.NewMenu(&service.MenuConfig{
		DB:             db.Get(),
		MenuRepository: menuRepo,
	})
	orderItemService := service.NewOrderItem(&service.OrderItemConfig{
		DB:                  db.Get(),
		OrderItemRepository: orderItemRepo,
		MenuRepository:      menuRepo,
	})
	orderService := service.NewOrder(&service.OrderConfig{
		DB:             db.Get(),
		DeliveryRepo:   deliveryRepo,
		PaymentOptRepo: paymentOptRepo,
		CouponRepo:     couponRepo,
		OrderRepo:      orderRepo,
		OrderItemRepo:  orderItemRepo,
	})
	reviewService := service.NewReview(&service.ReviewConfig{
		DB:         db.Get(),
		ReviewRepo: reviewRepo,
	})

	router := NewRouter(&RouterConfig{
		AuthService:      authService,
		UserService:      userService,
		MenuService:      menuService,
		OrderItemService: orderItemService,
		OrderService:     orderService,
		ReviewService:    reviewService,
	})
	log.Fatalln(router.Run())
}
