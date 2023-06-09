package server

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/raflynagachi/kopi-santai-backend/config"
	"github.com/raflynagachi/kopi-santai-backend/db"
	"github.com/raflynagachi/kopi-santai-backend/repository"
	"github.com/raflynagachi/kopi-santai-backend/service"
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
	gameRepo := repository.NewGame()
	promoRepo := repository.NewPromo()
	categoryRepo := repository.NewCategory()

	authService := service.NewAuth(&service.AuthConfig{
		DB:             db.Get(),
		UserRepository: userRepo,
		GameRepository: gameRepo,
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
		PromotionRepo:  promoRepo,
	})
	reviewService := service.NewReview(&service.ReviewConfig{
		DB:         db.Get(),
		ReviewRepo: reviewRepo,
	})
	deliveryService := service.NewDelivery(&service.DeliveryConfig{
		DB:           db.Get(),
		DeliveryRepo: deliveryRepo,
	})
	couponService := service.NewCoupon(&service.CouponConfig{
		DB:         db.Get(),
		CouponRepo: couponRepo,
	})
	gameService := service.NewGame(&service.GameConfig{
		DB:         db.Get(),
		GameRepo:   gameRepo,
		CouponRepo: couponRepo,
	})
	promoService := service.NewPromo(&service.PromoConfig{
		DB:        db.Get(),
		PromoRepo: promoRepo,
	})
	paymentOptService := service.NewPaymentOpt(&service.PaymentOptConfig{
		DB:             db.Get(),
		PaymentOptRepo: paymentOptRepo,
	})
	categoryService := service.NewCategory(&service.CategoryConfig{
		DB:           db.Get(),
		CategoryRepo: categoryRepo,
	})

	router := NewRouter(&RouterConfig{
		AuthService:       authService,
		UserService:       userService,
		MenuService:       menuService,
		OrderItemService:  orderItemService,
		OrderService:      orderService,
		ReviewService:     reviewService,
		DeliveryService:   deliveryService,
		CouponService:     couponService,
		GameService:       gameService,
		PromoService:      promoService,
		PaymentOptService: paymentOptService,
		CategoryService:   categoryService,
	})

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(func() {
		err := gameRepo.ResetTriedChance(db.Get())
		if err != nil {
			return
		}
	})
	s.StartAsync()
	log.Fatalln(router.Run(":" + config.Config.Port))
}
