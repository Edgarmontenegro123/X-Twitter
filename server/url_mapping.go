package server

import (
	"github.com/Edgarmontenegro123/X-Twitter/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, tweetController *controllers.TweetController, userController *controllers.UserController) {
	// Configurar rutas
	router.POST("/publish_tweet", tweetController.PublishTweetController)
	router.GET("/get_tweets/:user_id", tweetController.GetTweetsController)
	router.POST("/follow_user", userController.FollowUserController)
	router.GET("/get_followers/:user_id", userController.GetFollowersController)
}
