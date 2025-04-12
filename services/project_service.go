package services

import (
	"context"

	"github.com/mmhtoo/go-logger-api/entities"
	"github.com/mmhtoo/go-logger-api/repositories"
)

type ProjectService struct {
	ProjectRepository *repositories.ProjectRepository
}

func (service *ProjectService) GetAllProjects(ctx context.Context) ([]entities.ProjectEntity, error){
	projects, err := service.ProjectRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return projects, nil
}