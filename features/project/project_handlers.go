package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	projectService *ProjectService
}

func NewProjectHandler(projectService *ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) HandleGetAllProjects(c *gin.Context) {
	projects, err := h.projectService.GetAllProjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
		"message" : "Succcess!",
	})
}

func (h *ProjectHandler) HandleCreateProject(c *gin.Context) {
	payload := c.MustGet("payload").(*ProjectCreateReqDto)
	savedProject, err := h.projectService.CreateProject(
		c.Request.Context(),
		&ProjectCreateInput{
			Id: uuid.NewString(),
			ProjectCreateReqDto: payload,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Failed to create project!",
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": savedProject,
		"message": "Success!",
	})
}