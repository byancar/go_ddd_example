package app

import (
	"go_ddd_example/internal/app/user"
	"go_ddd_example/internal/infrastructure/config"
	users "go_ddd_example/internal/infrastructure/persistence/sqlc"
)

type App struct {
	UserService user.UserService
}

func NewApp(cfg *config.AppConfig, querier users.Querier) *App {
	userRepository := users.NewUserRepository(querier)
	userService := user.NewUserService(userRepository)

	return &App{
		UserService: userService,
	}
}
