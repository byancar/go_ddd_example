package main

import (
	"github.com/gin-gonic/gin"
	"go_ddd_example/internal/app"
	"go_ddd_example/internal/app/user"
	"go_ddd_example/internal/infrastructure/config"
	users "go_ddd_example/internal/infrastructure/persistence/sqlc"
	"log"
)

func main() {

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	conn, err := config.NewConnection(appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.Database)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(appConfig, users.New(conn))
	router := gin.Default()
	userHandler := user.NewUserHandler(app.UserService)

	api := router.Group("/api")
	{
		api.POST("/users", userHandler.Create)
	}

	serverPort := ":8080"
	log.Println("Server is running on port", serverPort)
	err = router.Run(serverPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
