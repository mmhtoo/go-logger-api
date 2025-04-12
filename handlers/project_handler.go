package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmhtoo/go-logger-api/services"
)

type ProjectHandler struct {
	Service *services.ProjectService
}

func (h *ProjectHandler) CreateNewProjectHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
	})
}

func (h *ProjectHandler) GetAllProjectHandler(c *gin.Context){
	projects,err := h.Service.GetAllProjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})
}