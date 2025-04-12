package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/config"
	"github.com/mmhtoo/go-logger-api/handlers"
	"github.com/mmhtoo/go-logger-api/repositories"
	"github.com/mmhtoo/go-logger-api/services"
)

func main() {
	// load env variables
	env := config.LoadEnv()
	// db connection
	database, err := config.NewDatabase(
		context.Background(),
		config.DatabaseParameters{
			Host: env.DB_HOST,
			Port: env.DB_PORT,
			DBName: env.DB_NAME,
			Timeout: 10 * time.Second,
		},
		config.DatabaseCredentials{
			Username: env.DB_USERNAME,
			Password: env.DB_PASSWORD,
		},
	)

	if err != nil {
		panic("Failed to connect database " + err.Error())
	}

	router := gin.Default()
	v1Router := router.Group("/api/v1")

	// handlers
	projectHandler := handlers.ProjectHandler{
		Service: &services.ProjectService{
			ProjectRepository: &repositories.ProjectRepository{
				Database: database,
			},
		},
	}
	v1Router.POST("/projects", projectHandler.CreateNewProjectHandler)
	v1Router.GET("/projects", projectHandler.GetAllProjectHandler)

	v1Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong! "+env.DB_NAME,
		})
	})

	router.Run(":8080")
}
