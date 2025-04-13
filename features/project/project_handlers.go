package project

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mmhtoo/go-logger-api/helpers"
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
	c.JSON(
		http.StatusOK, 
		helpers.NewAPIDataResponse(projects, "Successfully retrieved!"),
	)
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
		c.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(err.Error(),"Failed to create project!"),
	  )
	}
	c.JSON(
		http.StatusCreated,
		helpers.NewAPIDataResponse(savedProject,"Successfully created!"),
	)
}

func (h *ProjectHandler) HandleUpdateProject(c *gin.Context){
	payload := c.MustGet("payload").(*ProjectUpdateReqDto)
	projectId, isOk := c.Params.Get("id")
	if(!isOk){
		c.JSON(
			http.StatusBadRequest,
			helpers.NewAPIErrorResponse(
				errors.New("Invalid project id!"),
				"Validation failed!",
			),
		)
		return
	}
	err := h.projectService.UpdateProject(
		c.Request.Context(),
		&ProjectUpdateInput{
			Id: projectId,
			ProjectUpdateReqDto: payload,
		},
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(
				err,
				"Failed to update!",
			),
		)
		return
	}
	c.JSON(
		http.StatusOK,
		helpers.NewAPIBaseResponse(
			"Successfully updated!",
		),
	)
}

func (h *ProjectHandler) HandleFindById(c *gin.Context){
	id, isOk := c.Params.Get("id")
	if !isOk {
		c.JSON(
			http.StatusBadRequest,
			helpers.NewAPIErrorResponse(
				errors.New("Invalid project id!"),
				"Validation failed!",
			),
		)
	}
	project, err := h.projectService.FindById(
		c.Request.Context(),
		id,
	)
	if err != nil {
		if project.Id == "" {
			c.JSON(
				http.StatusNotFound,
				helpers.NewAPIBaseResponse("Not project found with id: "+id),
			)
			return
		}
		c.JSON(
			http.StatusInternalServerError,
			helpers.NewAPIErrorResponse(
				err.Error(),
				"Failed to find project!",
			),
		)
		return
	}
	c.JSON(
		http.StatusOK,
		helpers.NewAPIDataResponse(
			project, 
			"Successfully retrieved!",
		),
	)
}