package controllers

import (
	"fmt"
	"github.com/Edgarmontenegro123/X-Twitter/internal/tweet"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TweetController struct {
	TweetService *tweet.TweetService
}

// NewTweetController crea nueva instancia de TweetController
func NewTweetController(tweetService *tweet.TweetService) *TweetController {
	return &TweetController{TweetService: tweetService}
}

// PublishTweetController maneja la publicación de tweets
func (tc *TweetController) PublishTweetController(c *gin.Context) {
	//Decodificar el cuerpo JSON de la solicitud
	var request struct {
		UserID  int    `json:"user_id"`
		Content string `json:"content"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error al decodificar el cuerpo JSON: %v", err)})
		return
	}

	// Validar la presencia de UserID y Content en la solicitud
	if request.UserID == 0 || request.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID y Content son campos requeridos"})
		return
	}

	// Llamar al servicio para publicar el tweet
	err = tc.TweetService.PublishTweet(request.UserID, request.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al publicar el tweet: %v", err)})
		return
	}
	// Respuesta exitosa
	c.Status(http.StatusCreated)
}

// GetTweetsController maneja la obtención de tweets
func (tc *TweetController) GetTweetsController(c *gin.Context) {
	// Obtener el ID del usuario desde los parámetros de la solicitud
	userID := c.Param("user_id")

	// Convertir el userID a un entero
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Imprimir un mensaje de registro para verificar el userID
	fmt.Printf("UserID: %d\n", userIDInt)

	// Verificar si el usuario existe
	userExists := tc.TweetService.UserExists(userIDInt)
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "El usuario no existe"})
		return
	}

	// Obtener tweets del servicio de tweets
	tweets, err := tc.TweetService.GetTweets(userIDInt)
	if err != nil {
		fmt.Printf("Error al obtener tweets: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener tweets"})
		return
	}
	// Enviar la lista de tweets como respuesta
	c.JSON(http.StatusOK, gin.H{"tweets": tweets})
}
