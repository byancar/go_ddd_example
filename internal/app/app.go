package app

import (
	"go_ddd_example/internal/app/user"
	"go_ddd_example/internal/infrastructure/config"
)

type App struct {
	UserService user.UserService
}

func NewApp(cfg *config.AppConfig) *App {
	//userRepository := "" //TODO implement sqlc repository
	//userService := user.NewUserService(userRepository)
	//
	//return &App{
	//	UserService: userService,
	//}
	return nil
}
