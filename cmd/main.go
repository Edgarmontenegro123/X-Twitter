package main

import (
	"fmt"
	"github.com/Edgarmontenegro123/X-Twitter/internal/db"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
)

func main() {
	// Crear una instancia de la db en memoria
	database := db.NewDatabase()

	// Crear una instancia del servicio de usuarios
	userService := user.NewUserService(database)

	// Crear y guardar un usuario en la db
	user1 := user.User{ID: 1, Username: "usuario1"}
	database.SaveUser(user1)

	// Crear y guardar otro usuario
	user2 := user.User{ID: 2, Username: "usuario2"}
	database.SaveUser(user2)

	// Imprimir estado actual de la base de datos
	fmt.Println("Usuarios en la base de datos:")
	for _, u := range database.GetUsers() {
		fmt.Printf("%+v\n", u)
	}

	// usuario1 sigue a usuario2
	err := userService.Follow(user1.ID, user2.ID)
	if err != nil {
		fmt.Println("Error al seguir a un usuario", err)
	}

	// Obtener lista de seguidores de user1
	followers, err := userService.GetFollowers(user1.ID)
	if err != nil {
		fmt.Println("Error al obtener la lista de seguidores", err)
	} else {
		fmt.Printf("Seguidores de %s: %v\n\n", user1.Username, followers)
	}
}
