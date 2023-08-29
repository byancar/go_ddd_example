package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ddd_example/internal/infrastructure/config"
	"log"
)

func main() {

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	fmt.Println("Database Driver:", appConfig.Database.Driver)

	//app := app.NewApp(appConfig)
	router := gin.Default()

	//userHandler := user.NewUserHandler(app.UserService)
	//
	//api := router.Group("/api"){
	//	api.POST("/users", userHandler )
	//}

	serverPort := ":8080" // Porta do servidor
	log.Println("Server is running on port", serverPort)
	err = router.Run(serverPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
