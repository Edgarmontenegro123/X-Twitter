package main

import (
	"fmt"
	"github.com/Edgarmontenegro123/X-Twitter/internal/controllers"
	"github.com/Edgarmontenegro123/X-Twitter/internal/db"
	"github.com/Edgarmontenegro123/X-Twitter/internal/tweet"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
	"github.com/Edgarmontenegro123/X-Twitter/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear instancia de DB en memoria
	database := db.NewDatabase()

	// Crear una instancia del servicio de usuarios y tweets
	userService := user.NewUserService(database)
	tweetService := tweet.NewTweetService(database)

	// Crear instancias de los controllers
	tweetController := controllers.NewTweetController(tweetService)
	userController := controllers.NewUserController(userService)

	// Configurar enrutador gin
	router := gin.Default()

	// Configurar rutas usando la función setupRoutes
	server.SetupRoutes(router, tweetController, userController)

	// Ejecutar el servidor en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
	}

}
