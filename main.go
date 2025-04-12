package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/middlewares"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {
	router := gin.Default()

	v1Router := router.Group("/api/v1")

	v1Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})
	v1Router.POST("/testing", middlewares.CheckValidationMiddleware(User{}), func(c *gin.Context) {
		userData := c.MustGet("body").(*User)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s ", userData.Username),
		})
	})

	router.Run(":8080")
}
