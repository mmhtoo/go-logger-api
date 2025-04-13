package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/config"
	"github.com/mmhtoo/go-logger-api/features/project"
	"github.com/mmhtoo/go-logger-api/helpers"
	"github.com/mmhtoo/go-logger-api/middlewares"
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

	if env.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	v1Router := router.Group("/api/v1")

	// handlers
	projectRepo := project.NewProjectRepository(database)
	projectService := project.NewProjectService(projectRepo)
	projectHandler := project.NewProjectHandler(projectService)
	v1Router.GET("/projects", projectHandler.HandleGetAllProjects)
	v1Router.POST("/projects", middlewares.CheckValidationMiddleware(project.ProjectCreateReqDto{}), projectHandler.HandleCreateProject)
	v1Router.PUT("/projects/:id", middlewares.CheckValidationMiddleware(project.ProjectUpdateReqDto{}), projectHandler.HandleUpdateProject)
	v1Router.GET("/projects/:id", projectHandler.HandleFindById)

	v1Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.NewAPIBaseResponse("Pong!"))
	})

	router.Run(env.PORT)
}
