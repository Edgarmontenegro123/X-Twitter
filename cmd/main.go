package main

import (
	"fmt"
	"github.com/Edgarmontenegro123/X-Twitter/internal/db"
	"github.com/Edgarmontenegro123/X-Twitter/internal/tweet"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
)

func main() {
	// Crear una instancia de la db en memoria
	database := db.NewDatabase()

	// Crear una instancia del servicio de usuarios y tweets
	userService := user.NewUserService(database)
	tweetService := tweet.NewTweetService(database)

	// Crear y guardar un usuario1 en la db
	user1 := user.User{ID: 1, Username: "usuario1"}
	database.SaveUser(user1)
	// Crear y guardar usuario2
	user2 := user.User{ID: 2, Username: "usuario2"}
	database.SaveUser(user2)
	// Crear y guardar usuario3
	user3 := user.User{ID: 3, Username: "usuario3"}
	database.SaveUser(user3)
	// Crear y guardar usuario4
	user4 := user.User{ID: 4, Username: "usuario4"}
	database.SaveUser(user4)

	// Publicar tweets
	// Usuario1
	err := tweetService.PublishTweet(user1.ID, "Mi primer tweet!")
	if err != nil {
		fmt.Println("Error al publicar tweet", err)
	}
	err = tweetService.PublishTweet(user1.ID, "Mi segundo tweet!")
	if err != nil {
		fmt.Println("Error al publicar tweet", err)
	}

	// Usuario2
	err = tweetService.PublishTweet(user2.ID, "Tweet de usuario2")
	if err != nil {
		fmt.Println("Error al publicar tweet", err)
	}

	// Usuario3
	err = tweetService.PublishTweet(user3.ID, "Tweet de usuario3")
	if err != nil {
		fmt.Println("Error al publicar tweet", err)
	}

	// Obtener tweets del usuario1
	tweets, err := tweetService.GetTweets(user1.ID)
	if err != nil {
		fmt.Println("Error al obtener tweets", err)
	} else {
		fmt.Printf("Tweets de %s: %v\n\n", user1.Username, tweets)
	}

	// usuario2 sigue a usuario1
	err = userService.Follow(user1.ID, user2.ID)
	if err != nil {
		fmt.Println("Error al seguir a un usuario", err)
	}
	// usuario3 sigue a usuario1
	err = userService.Follow(user1.ID, user3.ID)
	if err != nil {
		fmt.Println("Error al seguir a un usuario", err)
	}
	// usuario4 sigue a usuario2
	err = userService.Follow(user4.ID, user2.ID)
	if err != nil {
		fmt.Println("Error al seguir a un usuario", err)
	}
	// usuario3 sigue a usuario4
	err = userService.Follow(user4.ID, user3.ID)
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

	// Obtener tweets timeline del usuario2
	timelineTweets, err := tweetService.GetTweetsTimeline(user2.ID)
	if err != nil {
		fmt.Println("Error al obtener timeline", err)
	} else {
		fmt.Printf("Timeline de %s: %v\n\n", user2.Username, timelineTweets)
	}

	// Imprimir estado final de la base de datos
	fmt.Println("Estado final de la base de datos:")
	for _, u := range database.GetUsers() {
		fmt.Printf("%+v\n", u)
	}
	for userID, tweets := range database.Tweets {
		fmt.Printf("Tweets del usuario %d: %+v\n", userID, tweets)
	}

	tweetService.PrintUserTimeline(4)
}
