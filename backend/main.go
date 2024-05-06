package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var registeredUsers []User

func main() {
	router := gin.Default()

	// Интеграция CORS для всех маршрутов
	router.Use(corsMiddleware())

	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)

	router.Run(":8081")
}

func registerHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registeredUsers = append(registeredUsers, user)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s registered successfully!", user.Email)})
}

func loginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, registeredUser := range registeredUsers {
		if registeredUser.Email == user.Email && registeredUser.Password == user.Password {
			c.String(http.StatusOK, "OK")
			return
		}
	}

	c.String(http.StatusUnauthorized, "Unauthorized")
}

// Функция для установки заголовков CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
