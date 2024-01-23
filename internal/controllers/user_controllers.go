package controllers

import (
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserController maneja las solicitudes relacionadas con usuarios
type UserController struct {
	UserService *user.UserService
}

type FollowRequest struct {
	UserID     int `json:"user_id" binding:"required"`
	FollowerID int `json:"follower_id" binding:"required"`
}

type CreateUserRequest struct {
	ID       int    `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
}

// NewUserController crea una nueva instancia de UserController
func NewUserController(userService *user.UserService) *UserController {
	return &UserController{UserService: userService}
}

// CreateUserController maneja la creación de usuarios
func (uc *UserController) CreateUserController(c *gin.Context) {
	var createUserRequest CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de solicitud no válidos"})
		return
	}
	// Verificar si se proporciona un ID en la solicitud
	if createUserRequest.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un ID diferente de 0 para crear el usuario"})
		return
	}

	// Crear una nueva instancia de usuario
	newUser := user.User{
		ID: createUserRequest.ID,
	}

	// Guardar el nuevo usuario en la base de datos
	uc.UserService.SaveUser(newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}

// FollowUserController maneja la solicitud de seguir a un usuario
func (uc *UserController) FollowUserController(c *gin.Context) {
	// Obtener datos de la solicitud
	var followRequest FollowRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de solicitud no válidos"})
		return
	}
	// Intentar seguir al usuario
	err := uc.UserService.Follow(followRequest.UserID, followRequest.FollowerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Usuario seguido con éxito"})
}

// GetFollowersController maneja la obtención de seguidores de un usuario
func (uc *UserController) GetFollowersController(c *gin.Context) {
	// Obtener el ID del usuario desde los parámetros de la ruta
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}
	// Obtener la lista de seguidores del usuario
	followers, err := uc.UserService.GetFollowers(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	// Respuesta exitosa con la lista de seguidores
	c.JSON(http.StatusOK, gin.H{"followers": followers})
}
