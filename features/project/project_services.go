package project

import (
	"context"
	"log"
)

type ProjectService struct {
	projectRepository *ProjectRepository
}

func NewProjectService(projectRepository *ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
	}
}

func (service *ProjectService) GetAllProjects(ctx context.Context) ([]ProjectEntity, error) {
	projects, err := service.projectRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (service *ProjectService) CreateProject(ctx context.Context, input *ProjectCreateInput) (ProjectEntity, error){
	savedProject, err := service.projectRepository.Save(input, ctx)
	if err != nil {
		log.Fatalf("Error while saving project: %s", err)
		return savedProject, err
	}
	return savedProject, nil
}

func (service *ProjectService) UpdateProject(ctx context.Context, input *ProjectUpdateInput) (error){
	err := service.projectRepository.UpdateById(input, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProjectService) FindById(ctx context.Context, id string) (ProjectEntity, error){
	project, err := service.projectRepository.findById(id, ctx)
	return project, err
}