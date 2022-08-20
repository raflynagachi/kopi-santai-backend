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
	orderRepo := repository.NewOrder()

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
	orderService := service.NewOrder(&service.OrderConfig{
		DB:              db.Get(),
		OrderRepository: orderRepo,
		MenuRepository:  menuRepo,
	})

	router := NewRouter(&RouterConfig{
		AuthService:  authService,
		UserService:  userService,
		MenuService:  menuService,
		OrderService: orderService,
	})
	log.Fatalln(router.Run())
}
