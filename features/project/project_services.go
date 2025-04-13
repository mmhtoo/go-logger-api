package project

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mmhtoo/go-logger-api/config"
	"github.com/mmhtoo/go-logger-api/features/jwt_secret"
	"github.com/mmhtoo/go-logger-api/helpers"
)

type ProjectService struct {
	projectRepository *ProjectRepository
	jwtSecretService *jwt_secret.JwtSecretService
}

func NewProjectService(
	projectRepository *ProjectRepository, 
	jwtSecretService *jwt_secret.JwtSecretService,
) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
		jwtSecretService: jwtSecretService,
	}
}

func (service *ProjectService) GetAllProjects(ctx context.Context) ([]ProjectEntity, error) {
	projects, err := service.projectRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (service *ProjectService) CreateProject(
	ctx context.Context, 
	input *ProjectCreateInput,
) (ProjectEntity, error){
	result, err := helpers.WithTx(
		ctx, 
		service.projectRepository.database,
		func(trxn *sql.Tx) (any,error) {
			savedProject, err := service.projectRepository.SaveWithTx(input, trxn)
			if err != nil {
				log.Printf("Error while saving project: %s", err)
				return savedProject, err
			}
			// save jwt secret for current created project
			err = service.jwtSecretService.CreateAndSaveJwtSecret(
				&jwt_secret.CreateAndSaveJwtSecretInput{
					KeyName: config.DEFAULT_JWT_SECRET_KEY_NAME,
					ProjectId: savedProject.Id,
					UserId: input.CreatedUserId,
				}, 
				trxn,
			)
			if err != nil {
				log.Printf("Error while creating JWT secret: %s", err)
				return savedProject, err
			}
			return savedProject, nil
		},
	)
	if err != nil {
		return ProjectEntity{}, err
	}
	fmt.Printf("result ", result)
	return result.(ProjectEntity), nil
}

func (service *ProjectService) UpdateProject(
	ctx context.Context, 
	input *ProjectUpdateInput,
) (error){
	err := service.projectRepository.UpdateById(input, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProjectService) FindById(
	ctx context.Context, 
	id string,
) (ProjectEntity, error){
	project, err := service.projectRepository.findById(id, ctx)
	return project, err
}