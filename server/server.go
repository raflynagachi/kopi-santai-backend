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

	router := NewRouter(&RouterConfig{
		AuthService:      authService,
		UserService:      userService,
		MenuService:      menuService,
		OrderItemService: orderItemService,
	})
	log.Fatalln(router.Run())
}
